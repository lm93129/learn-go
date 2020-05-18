package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func TestTcpClient(t *testing.T) {
	c, err := net.DialTimeout("tcp", "192.168.5.58:6667", 10*time.Second)
	if err != nil {
		fmt.Println("failed to connect", err)
	}
	time.Sleep(time.Millisecond)
	defer c.Close()
	sendBody, _ := hex.DecodeString("7e0200003f000004021895000b00000000000400030158ccaa06cb79f500950000000016010516541501040000697402020000030200002504000000002b040000000030010031010b3201467c7e")
	go func() {
		for {
			_, _ = c.Write(sendBody)
		}
	}()
	for {
		var buf = make([]byte, 45)
		n, _ := c.Read(buf)
		t.Log(n, string(buf[:n]))
		if string(buf[:n]) == "7E80010005000004021895000B000B0200000D7E" {
			return
		}
		t.Error("返回数据错误")
	}

}

func TestStarRock(t *testing.T) {
	addr := "192.168.5.58:6667"
	connections := 100
	t.Logf("连接到 %s", addr)
	var conns []net.Conn
	for i := 0; i < connections; i++ {
		c, err := net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			fmt.Println("failed to connect", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(time.Millisecond)
	}
	defer func() {
		for _, c := range conns {
			_ = c.Close()
		}
	}()
	log.Printf("完成初始化 %d 连接", len(conns))
	tts := time.Second

	tts = time.Millisecond * 5
	sendBody, _ := hex.DecodeString("7e0200003f000004021895000b00000000000400030158ccaa06cb79f500950000000016010516541501040000697402020000030200002504000000002b040000000030010031010b3201467c7e")

	go func() {
		for i := 0; i < connections; i++ {
			conn := conns[i]
			send(conn, sendBody, tts)
		}

	}()
	for {
		go func() {
			for i := 0; i < connections; i++ {
				conn := conns[i]
				body := read(conn)
				t.Log(body)
			}
		}()
	}

}

func send(conn net.Conn, sendBody []byte, tts time.Duration) {
	time.Sleep(tts)
	_, _ = conn.Write(sendBody)
}

func read(c net.Conn) string {

	var buf = make([]byte, 45)
	n, _ := c.Read(buf)
	return string(buf[:n])

}
