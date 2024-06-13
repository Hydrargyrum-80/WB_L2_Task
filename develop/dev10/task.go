package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	host := flag.String("host", "", "Target host (IP or domain name)")
	port := flag.Int("port", 0, "Target port")
	timeout := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()
	if *host == "" || *port == 0 {
		fmt.Println("Usage: go-telnet --host <host> --port <port> [--timeout <timeout>]")
		os.Exit(1)
	}
	serverAddr := fmt.Sprintf("%s:%d", *host, *port)
	conn, err := net.DialTimeout("tcp", serverAddr, *timeout)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", serverAddr, err)
		os.Exit(1)
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		_, err := io.Copy(conn, os.Stdin)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		done <- struct{}{}
	}()
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		done <- struct{}{}
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	select {
	case <-done:
	case <-signals:
	}
}
