package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	message = "Hello world\n\n"
	//message = "AMQP\x00\x01\x00\x00\n\n"
	//message = "GET http://my.tls.server:8000/index.html HTTP/1.1\n\n"
	//message = "CONNECT my.tls.server:8000 HTTP/1.1\n\n"
)

func main() {
	// direct no sni
	//tlsClient("192.168.124.1:8000", "")
	tlsClient("192.168.124.1:8000", "192.168.124.1")
	// direct using sni
	//tlsClient("10.0.135.20:8000", "192.168.124.1")

	// haproxy no sni
	//tlsClient("127.0.0.1:8443", "")
	// haproxy using sni
	//tlsClient("127.0.0.1:8443", "192.168.124.1")
	//tlsClient("127.0.0.1:8443", "my.tls.server")

	// squid no sni
	//tlsClient("127.0.0.1:443", "")
	// squid using sni
	//tlsClient("127.0.0.1:443", "my.tls.server")
	//tlsClient("my.tls.server:8000", "my.tls.server")

	// squid no sni
	//tlsClient("127.0.0.1:3128", "")
	// squid using sni
	//tlsClient("127.0.0.1:3128", "192.168.124.1")
}

func tlsClient(addr string, sni string) {
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("client: error loading keys: %s", err)
	}
	caPem, err := os.ReadFile("certs/ca.pem")
	if err != nil {
		log.Fatalf("client: error reading ca.pem: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caPem)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ServerName:   sni, // SNI to force TLS validation against the indicated name (not 127.0.0.1)
	}

	// connecting
	conn, err := tls.Dial("tcp", addr, config)
	if err != nil {
		log.Printf("client: dial error %s: %s", addr, err)
		fmt.Println()
		return
	}

	defer conn.Close()
	log.Printf("client: connected to: %s", conn.RemoteAddr())

	state := conn.ConnectionState()
	for _, v := range state.PeerCertificates {
		//key, _ := x509.MarshalPKIXPublicKey(v.PublicKey)
		//fmt.Println(string(key))
		log.Println(v.Subject)
	}
	log.Println("client: handshake:", state.HandshakeComplete)

	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client: error writing: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", message, n)

	reply := make([]byte, 256)
	n, err = conn.Read(reply)

	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Printf("client: exiting")
	fmt.Println()
}
