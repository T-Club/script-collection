package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
	"strconv"
)

func main() {
	flag.IntVar(&cpuCores, "n", 2, "number of cpu cores to use")
	flag.StringVar(&ip, "i", "127.0.0.1", "target ip address")
	flag.Parse()

	runtime.GOMAXPROCS(cpuCores)
	s := &scanner{ip}
	ch := make(chan int)
	for i := 0; i < 65535; i++ {
		go s.scan(ch, i)
	}
	count := 0
	for {
		select {
		case <-ch:
			count++
			if count >= 65535 {
				return
			}
		default:
		}
	}
}

func (s *scanner) scan(ch chan int, port int) {
	if _, err := net.Dial("tcp", s.ip+":"+strconv.Itoa(port)); err != nil {
		ch <- 0
		return
	}
	ch <- 0
	fmt.Println("open:", port)
}

type scanner struct {
	ip string
}

var (
	cpuCores int
	ip       string
)
