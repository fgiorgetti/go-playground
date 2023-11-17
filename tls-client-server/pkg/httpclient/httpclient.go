package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("HTTP_PROXY", "http://192.168.124.1:3128")
	os.Setenv("HTTPS_PROXY", "http://192.168.124.1:3128")

	httpRequest("http://192.168.124.1:8080", "")
	httpRequest("https://www.google.com", "www.google.com")
	httpRequest("https://192.168.124.1:8000", "192.168.124.1")
}

func httpRequest(url, sni string) {
	cert, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
	if err != nil {
		log.Fatalf("client: error loading keys: %s", err)
	}
	caPem, err := os.ReadFile("certs/ca.pem")
	if err != nil {
		log.Fatalf("client: error reading ca.pem: %s", err)
	}
	caCertPool, _ := x509.SystemCertPool()
	caCertPool.AppendCertsFromPEM(caPem)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ServerName:   sni, // SNI to force TLS validation against the indicated name
	}

	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}
	if sni != "" {
		t.TLSClientConfig = config
	}
	c := &http.Client{Transport: t}

	res, err := c.Get(url)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	fmt.Println("Status:", res.Status, "Code:", res.StatusCode)
	body, _ := io.ReadAll(res.Body)
	fmt.Println("Body:", string(body))
	fmt.Println()
}
