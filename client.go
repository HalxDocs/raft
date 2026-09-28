package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func pingNode(address string, from int) {
	client, err := rpc.DialHTTP("tcp", address)
	if err != nil {
		log.Printf("failed to connect to %s: %v", address, err)
		return
	}
	defer client.Close()

	args := PingArgs{
		From: from,
	}

	var reply PingReply

	err = client.Call("Node.ping", args, &reply)
}