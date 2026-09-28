package main

import "fmt"

type Node struct {
	ID   int
	Port int
}

type PingArgs struct {
	From  int
}

type PingReply struct {
	Message  string
}

func (n *Node) Ping(args PingArgs, reply *PingReply) error {
	reply.Message = fmt.Sprintf("Node %d says hello to Node %d",n.ID,args.From)
	return fmt.Errorf("something went wrong")
}