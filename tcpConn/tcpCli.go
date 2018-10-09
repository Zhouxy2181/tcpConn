package main

import (
	"net"
	"log"
	"fmt"
	"strconv"
)

func main()  {
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("conn err", err)
	}
	defer con.Close()

	str := "hello world"
	slen := fmt.Sprintf("%04d", len(str))

	n, err := con.Write([]byte(slen+str))
	if err != nil {
		log.Fatal("write err", err,n)
	}

	rslen := make([]byte, 4)
	n, err = con.Read(rslen)
	if err != nil {
		log.Fatal("read rsp msg len err", err, n)
	}

	n, err = strconv.Atoi(string(rslen))
	if err != nil {
		log.Fatal("string to int err", err)
	}

	rspmsg := make([]byte, n)
	n,err = con.Read(rspmsg)
	if err != nil {
		log.Fatal("read rsp msg err", err, n)
	}

	fmt.Println(string(rspmsg))
}
