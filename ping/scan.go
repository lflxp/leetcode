package main

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(host, port string) bool {
	remote := fmt.Sprintf("%s:%s", host, port)
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", remote) //转换IP格式 // 根据域名查找ip
	//fmt.Printf("%s", tcpAddr)
	// conn, err := net.DialTCP("tcp", nil, tcpAddr) //查看是否连接成功
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), 500*time.Microsecond) //建立tcp连接
	if err != nil {
		// fmt.Printf("no==%s:%s\r\n", host, port)
		return false
	}
	defer conn.Close()
	// fmt.Printf("ok==%s:%s\r\n", host, port)
	return true
}

func main() {
	ip := "192.168.1.1"
	port := 10000
	for i := 0; i < port; i++ {
		if ScanPort(ip, string(i)) {
			fmt.Printf("%s:%d ok \n", ip, i)
		}
	}
}
