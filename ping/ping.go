package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/tatsushid/go-fastping"
)

func Ping(ip []string) error {
	// saddr := net.IPAddr{IP: net.ParseIP("0.0.0.0")}
	// taddr, _ := net.ResolveIPAddr("ip", ip)
	// conn, err := net.DialIP("ip4:icmp", &saddr, taddr)
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close()
	// return nil

	p := fastping.NewPinger()

	for _, x := range ip {
		ra, err := net.ResolveIPAddr("ip4:icmp", x)
		if err != nil {
			fmt.Println(err)
			return err
		}
		p.AddIPAddr(ra)
	}

	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	// p.OnIdle = func() {
	// 	fmt.Println("finish")
	// }
	err := p.Run()
	if err != nil {
		return err
	}
	defer p.Stop()
	return nil
}

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

func ParseIps(in string) ([]string, error) {
	rs := []string{}
	if strings.Contains(in, "-") {
		tmp_a := strings.Split(in, ".")
		if len(tmp_a) != 4 {
			fmt.Println(tmp_a)
			return nil, errors.New("ip地址不正确")
		}
		A := []string{}
		B := []string{}
		C := []string{}
		D := []string{}
		for m, n := range tmp_a {
			if strings.Contains(n, "-") {
				tmp := strings.Split(n, "-")
				a, err := strconv.Atoi(tmp[0])
				if err != nil {
					return rs, err
				}
				b, err := strconv.Atoi(tmp[1])
				if err != nil {
					return rs, err
				}
				for i := a; i <= b; i++ {
					if m == 0 {
						A = append(A, fmt.Sprintf("%d", i))
					} else if m == 1 {
						B = append(B, fmt.Sprintf("%d", i))
					} else if m == 2 {
						C = append(C, fmt.Sprintf("%d", i))
					} else if m == 3 {
						D = append(D, fmt.Sprintf("%d", i))
					}
				}
			} else {
				if m == 0 {
					A = append(A, n)
				} else if m == 1 {
					B = append(B, n)
				} else if m == 2 {
					C = append(C, n)
				} else if m == 3 {
					D = append(D, n)
				}
			}
		}
		for _, a1 := range A {
			for _, b1 := range B {
				for _, c1 := range C {
					for _, d1 := range D {
						rs = append(rs, fmt.Sprintf("%s.%s.%s.%s", a1, b1, c1, d1))
					}
				}
			}
		}
	} else {
		rs = append(rs, in)
	}
	return rs, nil
}

func main() {
	// for i := 0; i < 256; i++ {
	// 	go func(n int) {
	// 		ips := fmt.Sprintf("192.168.50.%d", n)
	// 		err := Ping(ips)
	// 		if err != nil {
	// 			fmt.Println(fmt.Sprintf("%s %s", ips, err.Error()))
	// 		} else {
	// 			fmt.Println(fmt.Sprintf("%s ok", ips))
	// 		}
	// 	}(i)
	// }

	// ips := []string{}
	// for i := 0; i < 256; i++ {
	// 	ips = append(ips, fmt.Sprintf("192.168.50.%d", i))
	// }
	// err := Ping(ips)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < 256; i++ {
	// 	for _, x := range []string{"22", "80", "8080", "9090"} {
	// 		if ScanPort(fmt.Sprintf("192.168.50.%d", i), x) {
	// 			fmt.Printf("192.168.50.%d %s ok\n", i, x)
	// 		}
	// 	}
	// }

	a := "192.168.1.2"
	rs, err := ParseIps(a)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(rs)
	for _, ip := range rs {
		for i := 0; i < 100; i++ {
			if ScanPort(ip, fmt.Sprintf("%d", i)) {
				fmt.Printf(fmt.Sprintf("%s -> %d \n", ip, i))
			}
		}
	}
}
