package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	cert, err := tls.LoadX509KeyPair("certs/server.pem", "certs/server.key")
	if err != nil {
		log.Fatalf("server: error loading keys: %s", err)
	}

	// tls config
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	config.Rand = rand.Reader

	// server listener
	ip := "0.0.0.0"
	if len(os.Args) > 1 && os.Args[1] != "" {
		ip = os.Args[1]
	}
	service := ip + ":8000"
	listener, err := tls.Listen("tcp", service, config)
	if err != nil {
		log.Fatalf("server: error listening on %s: %s", service, err)
	}
	log.Printf("server: listening on %s", service)

	// main loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept error: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted connection from %s", conn.RemoteAddr())
		log.Printf("server: accepted connection to %s", conn.LocalAddr())
		tlsConn, ok := conn.(*tls.Conn)
		if !ok {
			log.Printf("Not a tls connection")
			continue
		}
		log.Print("tlsconn.localaddr ", tlsConn.LocalAddr())
		log.Print("tlsconn.localaddr ", tlsConn.NetConn().RemoteAddr())
		state := tlsConn.ConnectionState()

		for _, v := range state.PeerCertificates {
			log.Println(x509.MarshalPKIXPublicKey(v.PublicKey))
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 512)
	for {
		log.Printf("server: conn: waiting...")
		log.Printf("PID: %d", os.Getpid())
		time.Sleep(time.Minute * 2)
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("server: conn: read error: %s", err)
			break
		}
		log.Printf("server: conn: echo %q", string(buf[:n]))
		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Printf("server: write error: %s", err)
			break
		}
		log.Printf("server: conn: wrote %d bytes", n)
	}
	log.Println("server: conn: closed")
}
