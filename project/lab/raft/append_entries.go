package raft

type AppendEntriesArgs struct {
	Term         int //leader's term
	LeaderID     int //so follower can redirect clients
	PrevLogIndex int //index of log entry immediately preeding new ones
	PrevLogTerm  int //term of prevLogIndex entry
	//log entries to store (empty for heartbeat; may send more than one for efficiency)
	Entries     []LogEntry
	LeaderCommit int //leader's commitIndex
}

type AppendEntriesReply struct {
	Term    int  //currentTerm, for leader to update itself
	Success bool //true if follower contained entry matching prevLogIndex and prevLogTerm
}

func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	//DPrintf("leader[raft%v][term:%v] beat term:%v [raft%v][%v]", args.LeaderID, args.Term, rf.currentTerm, rf.me, rf.state)
	reply.Success = true
	if args.Term < rf.currentTerm {
		// Reply false if term < currentTerm (§5.1)
		reply.Success = false
		reply.Term = rf.currentTerm
		return
	}

	if args.Term > rf.currentTerm {
		//If RPC request or response contains term T > currentTerm:set currentTerm = T, convert to follower (§5.1)
		rf.currentTerm = args.Term
		rf.switchStateTo(Follower)

		// do not return here.
	}

	rf.electionTimer.Reset(randTimeDuration(ElectionTimeoutLower, ElectionTimeoutUpper))
}

func (rf *Raft) sendAppendEntries(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}
