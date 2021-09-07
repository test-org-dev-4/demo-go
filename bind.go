package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func connect() {
	l, err := net.Listen("tcp", "0.0.0.0:2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	var i int
	i = 5
}

func sshConfigure() {
	_ = ssh.InsecureIgnoreHostKey()
}
