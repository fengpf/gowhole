package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"gowhole/project/lab/labrpc"
)

// import "bytes"
// import "labgob"

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make(). set
// CommandValid to true to indicate that the ApplyMsg contains a newly
// committed log entry.
//
// in Lab 3 you'll want to send other kinds of messages (e.g.,
// snapshots) on the applyCh; at that point you can add fields to
// ApplyMsg, but set CommandValid to false for these other uses.
//
type ApplyMsg struct {
	CommandValid bool
	Command      interface{}
	CommandIndex int
}

const (
	Follower  StateType = iota //跟随者
	Candidate                  //候选人
	Leader                     //领导人

)

const (
	HeartbeatInterval    = time.Duration(120) * time.Millisecond //100 //心跳超时，要求1秒10次，所以是100ms一次
	ElectionTimeoutLower = time.Duration(150) * time.Millisecond
	ElectionTimeoutUpper = time.Duration(300) * time.Millisecond
)

type StateType uint64

var stmap = [...]string{
	"Follower",
	"Candidate",
	"Leader",
}

func (st StateType) String() string {
	return stmap[uint64(st)]
}

//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.

	//2A
	state          StateType   //服务器角色
	electionTimer  *time.Timer // 选举定时器
	heartbeatTimer *time.Timer // 心跳定时器
	voteCount      int         //投票数
	//*** persistent state on all serves ***
	//**** updated on stable storage before responding to rpcs ****

	//State: the log, the current index, &c) which must be updated in response to events arising in concurrent goroutines.
	//latest term server has seen (initialized to 0 on first boot，increases monotonically-单调地)
	//服务器最后一次知道的任期号（初始化为 0，持续递增）
	currentTerm int
	//candidateID that recieved vote in current term(or null if none)
	//在当前获得选票的候选人的 Id
	votedFor int //
	//log entries; each entry containes command for state machine,
	// and term when entry was received by leader(first index is 1)
	//日志条目集；每一个条目包含一个用户状态机执行的指令，和收到时的任期号
	//[]log
	entries []LogEntry

	//*** volatile state on all servers ***
	//所有服务器上经常变的
	//index of highest log entry known to be committed(initialized to 0, increases monotonically)
	//已知的最大的已经被提交的日志条目的索引值
	commitIndex int
	// index of highest log entry applied to state machine(initialized to 0, increases monotonically)
	//最后被应用到状态机的日志条目索引值（初始化为 0，持续递增）
	lastApplied int

	//*** volatile state on leaders ***
	//**** reinitialized after election ****
	//在领导人里经常改变的 （选举后重新初始化）

	// for each server, index of the next log entry to send to
	// that server (initialized to leader last log index+1)
	//对于每一个服务器，需要发送给他的下一个日志条目的索引值（初始化为领导人最后索引值加一）
	nextIndex []int
	//for each server, index of highest log entry known
	// to be replicated on server (initialized to 0, increases monotonically)
	//对于每一个服务器，已经复制给他的日志的最高索引值
	matchIndex []int
}

func (rf Raft) String() string {
	return fmt.Sprintf("[node(%d), state(%v), term(%d)]",
		rf.me, rf.state, rf.currentTerm)
}

type LogEntry struct {
	Term    int
	Command interface{}
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {

	var term int
	var isleader bool

	// Your code here (2A).
	rf.mu.Lock()
	term = rf.currentTerm
	isleader = rf.state == Leader
	rf.mu.Unlock()

	return term, isleader
}

//
// save Raft's persistent state to stable storage,
// where it can later be retrieved after a crash and restart.
// see paper's Figure 2 for a description of what should be persistent.
//
func (rf *Raft) persist() {
	// Your code here (2C).
	// Example:
	// w := new(bytes.Buffer)
	// e := labgob.NewEncoder(w)
	// e.Encode(rf.xxx)
	// e.Encode(rf.yyy)
	// data := w.Bytes()
	// rf.persister.SaveRaftState(data)
}

//
// restore previously persisted state.
//
func (rf *Raft) readPersist(data []byte) {
	if data == nil || len(data) < 1 { // bootstrap without any state?
		return
	}
	// Your code here (2C).
	// Example:
	// r := bytes.NewBuffer(data)
	// d := labgob.NewDecoder(r)
	// var xxx
	// var yyy
	// if d.Decode(&xxx) != nil ||
	//    d.Decode(&yyy) != nil {
	//   error...
	// } else {
	//   rf.xxx = xxx
	//   rf.yyy = yyy
	// }
}

//
// example RequestVote RPC arguments structure.
// field names must start with capital letters!
//
type RequestVoteArgs struct {
	// Your data here (2A, 2B).

	//2A
	// *** invoked by candidates to gather votes ***
	//
	Term         int //candidates's term  候选人的任期号
	CandidateId  int //candidate requesting vote  请求选票的候选人的 Id
	LastLogIndex int //index of candidate's last log entry  候选人的最后日志条目的索引值
	LastLogTerm  int //term of canidate's last log entry  候选人最后日志条目的任期号
}

//
// example RequestVote RPC reply structure.
// field names must start with capital letters!
//
type RequestVoteReply struct {
	// Your data here (2A).

	//2A
	//results:
	Term        int  //currentTerm，for candidate to update itself 当前任期号，以便于候选人去更新自己的任期号
	VoteGranted bool //true means candidate received vote 候选人赢得了此张选票时为真
}

//
// example RequestVote RPC handler.
//
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here (2A, 2B).

	//2A
	// **** receiver implementtation: ****
	//1、reply false if term < currentTerm
	//2、if votedFor is null or candidateId,
	// and candidate's log is at least
	// as up-to-date as receiver's log, grant vote.

	rf.mu.Lock()
	defer rf.mu.Unlock()

	//DPrintf("Candidate node[%v] term[%v] request vote: node[%v] state[%v] term[%v]\n",
	//	args.CandidateId, args.Term, rf.me, rf.state, rf.currentTerm)

	//Reply false if term < currentTerm
	if args.Term < rf.currentTerm ||
		(args.Term == rf.currentTerm && rf.votedFor != -1 && rf.votedFor != args.CandidateId) {
		reply.Term = rf.currentTerm
		reply.VoteGranted = false
		return
	}

	//If RPC request or response contains term T > currentTerm: set currentTerm = T, convert to follower (§5.1)
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.switchStateTo(Follower)
		// do not return here.
	}

	//If votedFor is null or candidateId, and candidate’s log is at least as up-to-date as receiver’s log, grant vote
	//如果 votedFor 为空或者为 candidateId，并且候选人的日志至少和自己一样新，那么就投票给他

	if rf.votedFor == -1 || rf.votedFor == args.CandidateId /*&& (rf.lastApplied == args.lastLogIndex && rf.log[rf.lastApplied].term == args.lastLogTerm) */ {
		DPrintf("node[%v] vote[%v] for node[%v]\n", rf.me, rf.voteCount, args.CandidateId)
		reply.VoteGranted = true
		rf.votedFor = args.CandidateId
	}

	rf.votedFor = args.CandidateId
	reply.Term = rf.currentTerm // not used, for better logging
	reply.VoteGranted = true
	// reset timer after grant vote
	rf.electionTimer.Reset(randTimeDuration(ElectionTimeoutLower, ElectionTimeoutUpper))
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
//
func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) bool {
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	return ok
}

//
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election. even if the Raft instance has been killed,
// this function should return gracefully.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	index := -1
	term := -1
	isLeader := true

	// Your code here (2B).

	return index, term, isLeader
}

//
// the tester calls Kill() when a Raft instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
//
func (rf *Raft) Kill() {
	// Your code here, if desired.
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
//
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here (2A, 2B, 2C).

	//2A
	//DPrintf("Make raft%v...", me)
	rf.currentTerm = 0
	rf.votedFor = -1    // voted for no one
	rf.state = Follower //当服务器程序启动时，他们都是跟随者身份

	rf.heartbeatTimer = time.NewTimer(HeartbeatInterval)

	//Raft 算法使用随机选举超时时间的方法来确保很少会发生选票瓜分的情况，就算发生也能很快的解决。
	//为了阻止选票起初就被瓜分，选举超时时间是从一个固定的区间（例如 150-300 毫秒）随机选择
	//这样可以把服务器都分散开以至于在大多数情况下只有一个服务器会选举超时
	rf.electionTimer = time.NewTimer(randTimeDuration(ElectionTimeoutLower, ElectionTimeoutUpper))

	go func(node *Raft) {
		for {
			select {
			case <-rf.electionTimer.C:

				rf.mu.Lock()
				switch rf.state {
				// 如果一个跟随者在一段时间里没有接收到任何消息，也就是选举超时，
				// 那么他就会认为系统中没有可用的领导者,并且发起选举以选出新的领导者。
				case Follower:
					//第三种可能的结果是候选人既没有赢得选举也没有输：如果有多个跟随者同时成为候选人，
					//那么选票可能会被瓜分以至于没有候选人可以赢得大多数人的支持。
					rf.switchStateTo(Candidate)

				case Candidate:
					//当这种情况发生的时候，
					//每一个候选人都会超时，然后通过增加当前任期号来开始一轮新的选举。
					//然而，没有其他机制的话，选票可能会被无限的重复瓜分

					rf.startElection()
				}
				rf.mu.Unlock()

			case <-rf.heartbeatTimer.C:

				rf.mu.Lock()
				//领导者周期性的向所有跟随者发送心跳包（即不包含日志项内容的附加日志项 RPCs）来维持自己的权威
				if rf.state == Leader {
					rf.heartbeats()
					rf.heartbeatTimer.Reset(HeartbeatInterval)
				}
				rf.mu.Unlock()
			}
		}
	}(rf)

	// initialize from state persisted before a crash
	//rf.mu.Lock()
	rf.readPersist(persister.ReadRaftState())
	//rf.mu.Unlock()

	return rf
}

//切换状态，调用者需要加锁
func (rf *Raft) switchStateTo(state StateType) {
	if state == rf.state {
		return
	}

	DPrintf("Term %d: server %d convert from %v to %v\n",
		rf.currentTerm, rf.me, rf.state, state)

	rf.state = state
	switch state {
	case Follower:
		rf.heartbeatTimer.Stop()
		rf.electionTimer.Reset(randTimeDuration(ElectionTimeoutLower, ElectionTimeoutUpper))
		rf.votedFor = -1

	case Candidate: //成为候选人后立马进行选举
		rf.startElection()

	case Leader:

		rf.electionTimer.Stop()
		rf.heartbeats()
		rf.heartbeatTimer.Reset(HeartbeatInterval)
	}
}

func randTimeDuration(lower, upper time.Duration) time.Duration {
	num := rand.Int63n(upper.Nanoseconds()-lower.Nanoseconds()) + lower.Nanoseconds()
	return time.Duration(num) * time.Nanosecond
}

func (rf *Raft) startElection() {
	//DPrintf("raft%v is starting election", rf.me)

	//要开始一次选举过程，跟随者先要增加自己的当前任期号，并且转换到候选人状态？
	rf.currentTerm += 1

	//每一个候选人在开始一次选举的时候会重置一个随机的选举超时时间，然后在超时时间内等待投票的结果
	//这样减少了在新的选举中另外的选票瓜分的可能性
	rf.electionTimer.Reset(randTimeDuration(ElectionTimeoutLower, ElectionTimeoutUpper))

	args := RequestVoteArgs{
		Term:        rf.currentTerm,
		CandidateId: rf.me,
	}

	var voteCount int32
	for peer := range rf.peers {
		if peer == rf.me {
			rf.votedFor = rf.me //vote for me
			atomic.AddInt32(&voteCount, 1)
			continue
		}

		go func(server int) {
			// DPrintf("raft%v[%v] is sending RequestVote RPC to raft%v\n", rf.me, rf.state, peer)

			reply := RequestVoteReply{}
			//并行的向集群中的其他服务器节点发送请求投票的 RPCs 来给自己投票
			if rf.sendRequestVote(server, &args, &reply) {

				rf.mu.Lock()
				//DPrintf("%v got RequestVote response from node %d, VoteGranted=%v, Term=%d\n",
				//	rf, server, reply.VoteGranted, reply.Term)

				//候选人会继续保持着当前状态直到以下三件事情之一发生：
				//(a) 他自己赢得了这次的选举，
				//(b) 其他的服务器成为领导者，
				//(c) 一段时间之后没有任何一个获胜的人。这些结果会分别的在下面的段落里进行讨论。
				if reply.VoteGranted && rf.state == Candidate {
					atomic.AddInt32(&voteCount, 1)

					//当一个候选人从整个集群的大多数服务器节点获得了针对同一个任期号的选票，那么他就赢得了这次选举并成为领导人
					//要求大多数选票的规则确保了最多只会有一个候选人赢得此次选举
					if atomic.LoadInt32(&voteCount) > int32(len(rf.peers)/2) {
						//一旦候选人赢得选举，他就立即成为领导人
						rf.switchStateTo(Leader)
					}
				} else if rf.currentTerm < reply.Term {
					//If RPC request or response contains term T > currentTerm:set currentTerm = T, convert to follower (§5.1)
					//在等待投票的时候，候选人可能会从其他的服务器接收到声明它是领导人的附加日志项 RPC
					//如果这个领导人的任期号（包含在此次的 RPC中）不小于候选人当前的任期号，
					//那么候选人会承认领导人合法并回到跟随者状态
					rf.currentTerm = reply.Term
					rf.switchStateTo(Follower)
				}
				//如果此次 RPC 中的任期号比自己小，那么候选人就会拒绝这次的 RPC 并且继续保持候选人状态。

				rf.mu.Unlock()

			} else {
				// DPrintf("raft%v[%v] vote:raft%v no reply, currentTerm:%v\n", rf.me, rf.state, peer, rf.currentTerm)
			}
		}(peer)
	}
}

//发送心跳包
func (rf *Raft) heartbeats() {
	for i := range rf.peers {
		if i == rf.me {
			continue
		}

		go func(server int) {
			rf.mu.Lock()
			if rf.state != Leader {
				rf.mu.Unlock()
				return
			}

			args := AppendEntriesArgs{
				Term:         rf.currentTerm,
				LeaderID:     rf.me,
				LeaderCommit: rf.commitIndex,
			}
			rf.mu.Unlock()

			reply := AppendEntriesReply{}
			if rf.sendAppendEntries(server, &args, &reply) {
				//If RPC request or response contains term T > currentTerm:set currentTerm = T, convert to follower (§5.1)
				rf.mu.Lock()

				if rf.state != Leader {
					rf.mu.Unlock()
					return
				}

				if reply.Success { //复制日志

				} else {
					if reply.Term > rf.currentTerm {
						rf.currentTerm = reply.Term
						rf.switchStateTo(Follower)
					}
				}

				rf.mu.Unlock()
			}
		}(i)
	}
}
