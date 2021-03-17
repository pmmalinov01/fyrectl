package sshsession

import (
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func RunCommand(cmd string) {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("p@veL123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, _ := ssh.Dial("tcp", "9.30.249.80:22", config)
	defer conn.Close()

	sess, err := conn.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()
	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	go io.Copy(os.Stdout, sessStdOut)
	sessStdErr, err := sess.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	go io.Copy(os.Stderr, sessStdErr)
	err = sess.Run(cmd)
	if err != nil {
		log.Fatal(err)
	}
}
