package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

    nodeID := flag.Int("id", 1,"node ID")
	port   := flag.Int("port",8001, "HTTP port")
	
	flag.Parse()

	node := &Node{
		ID: *nodeID,
		Port: *port,
	}

	err := rpc.Register(node)
	if err != nil {
		log.Fatal(err)
	}

	rpc.HandleHTTP()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Raft node", *nodeID)
	})

	address :=fmt.Sprintf(":%d", node.Port)

    listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal(err) 
	}

	log.Printf("Node %d listening on %s", node.ID, address)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}

}