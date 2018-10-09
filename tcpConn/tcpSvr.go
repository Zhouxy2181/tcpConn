package main

import (
	"fmt"
	"net"
	"log"
	"strconv"
)

func main()  {
	l, err := net.Listen("tcp","localhost:8080")
	if err != nil {
		log.Fatal("listen err", err)
	}
	defer l.Close()

	con, err := l.Accept()
	if err != nil {
		log.Fatal("conn err", err)
	}

	slen := make([]byte, 4)
	err = readn(con, slen)
	if err != nil {
		log.Fatal("read slen err", err)
	}

	mlen, err := strconv.Atoi(string(slen))
	if err != nil {
		log.Fatal("string to int err", err)
	}

	msglen := make([]byte, mlen)
	err = readn(con, msglen)
	if err != nil {
		log.Fatal("read msg err", err)
	}

	fmt.Println(string(msglen))
}

func readn(conn net.Conn, slen []byte) error {
	index := 0
	for index < len(slen) {
		n, err := conn.Read(slen[index:])
		if err != nil {
			return err
		}
		index += n
	}
	return nil
}
