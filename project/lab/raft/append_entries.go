package raft

type AppendEntriesArgs struct {
	Term         int //leader's term
	LeaderID     int //so follower can redirect clients
	PrevLogIndex int //index of log entry immediately preeding new ones
	PrevLogTerm  int //term of prevLogIndex entry
	//log entries to store (empty for heartbeat; may send more than one for efficiency)
	Entries      []LogEntry
	LeaderCommit int //leader's commitIndex
}

type AppendEntriesReply struct {
	Term    int  //currentTerm, for leader to update itself
	Success bool //true if follower contained entry matching prevLogIndex and prevLogTerm

	ConflictTerm  int // 2C
	ConflictIndex int // 2C
}

func (rf *Raft) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()

	//DPrintf("leader[raft%v][term:%v] beat term:%v [raft%v][%v]", args.LeaderID, args.Term, rf.currentTerm, rf.me, rf.state)
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

	if args.PrevLogIndex <= rf.snapshottedIndex {
		reply.Success = true

		// sync log if needed
		if args.PrevLogIndex+len(args.Entries) > rf.snapshottedIndex {
			// if snapshottedIndex == prevLogIndex, all log entries should be added.
			startIdx := rf.snapshottedIndex - args.PrevLogIndex
			// only keep the last snapshotted one
			rf.entries = rf.entries[:1]
			rf.entries = append(rf.entries, args.Entries[startIdx:]...)
		}

		return
	}

	// entries before args.PrevLogIndex might be unmatch
	// return false and ask Leader to decrement PrevLogIndex
	absoluteLastLogIndex := rf.getAbsoluteLogIndex(len(rf.entries) - 1)
	if absoluteLastLogIndex < args.PrevLogIndex {
		reply.Success = false
		reply.Term = rf.currentTerm
		// optimistically thinks receiver's log matches with Leader's as a subset
		reply.ConflictIndex = absoluteLastLogIndex + 1
		// no conflict term
		reply.ConflictTerm = -1
		return
	}

	if rf.entries[rf.getRelativeLogIndex(args.PrevLogIndex)].Term != args.PrevLogTerm {
		reply.Success = false
		reply.Term = rf.currentTerm
		// receiver's log in certain term unmatches Leader's log
		reply.ConflictTerm = rf.entries[rf.getRelativeLogIndex(args.PrevLogIndex)].Term

		// expecting Leader to check the former term
		// so set ConflictIndex to the first one of entries in ConflictTerm
		conflictIndex := args.PrevLogIndex
		// apparently, since rf.entries[0] are ensured to match among all servers
		// ConflictIndex must be > 0, safe to minus 1
		for rf.entries[rf.getRelativeLogIndex(conflictIndex-1)].Term == reply.ConflictTerm {
			conflictIndex--
			if conflictIndex == rf.snapshottedIndex+1 {
				// this may happen after snapshot,
				// because the term of the first log may be the current term
				// before lab 3b this is not going to happen, since rf.entries[0].Term = 0
				break
			}
		}
		reply.ConflictIndex = conflictIndex
		return
	}

	// compare from rf.entries[args.PrevLogIndex + 1]
	unmatch_idx := -1
	for idx := range args.Entries {
		if len(rf.entries) < rf.getRelativeLogIndex(args.PrevLogIndex+2+idx) ||
			rf.entries[rf.getRelativeLogIndex(args.PrevLogIndex+1+idx)].Term != args.Entries[idx].Term {
			// unmatch log found
			unmatch_idx = idx
			break
		}
	}

	if unmatch_idx != -1 {
		// there are unmatch entries
		// truncate unmatch Follower entries, and apply Leader entries
		rf.entries = rf.entries[:rf.getRelativeLogIndex(args.PrevLogIndex+1+unmatch_idx)]
		rf.entries = append(rf.entries, args.Entries[unmatch_idx:]...)
	}

	// Leader guarantee to have all committed entries
	// TODO: Is that possible for lastLogIndex < args.LeaderCommit?
	if args.LeaderCommit > rf.commitIndex {
		absoluteLastLogIndex := rf.getAbsoluteLogIndex(len(rf.entries) - 1)
		if args.LeaderCommit <= absoluteLastLogIndex {
			rf.setCommitIndex(args.LeaderCommit)
		} else {
			rf.setCommitIndex(absoluteLastLogIndex)
		}
	}
	
	reply.Success = true
}

func (rf *Raft) sendAppendEntries(server int, args *AppendEntriesArgs, reply *AppendEntriesReply) bool {
	ok := rf.peers[server].Call("Raft.AppendEntries", args, reply)
	return ok
}
