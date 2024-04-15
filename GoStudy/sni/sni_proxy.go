/**
 * @Author: Gong Yanhui
 * @Description: https://www.agwa.name/blog/post/writing_an_sni_proxy_in_go
 * @Date: 2024-04-15 14:20
 */

package sni

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

func BeginSNIProxy() {
	l, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, errIn := l.Accept()
		if errIn != nil {
			log.Println(errIn)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(clientConn net.Conn) {
	defer clientConn.Close()

	// read Client Hello from clientConn
	if err := clientConn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Println(err)
		return
	}

	clientHello, clientReader, err := peekClientHello(clientConn)
	if err != nil {
		log.Println(err)
		return
	}

	// 增加DNS解析缓存

	//if !strings.HasSuffix(clientHello.ServerName, ".internal.example.com") { // TODO 应用匹配
	//	log.Println("ClientHello.ServerName is not *.internal.example.com")
	//	return
	//}

	backendConn, err := net.Dial("tcp", net.JoinHostPort(clientHello.ServerName, "443"))
	if err != nil {
		log.Println(err)
		return
	}
	defer backendConn.Close()

	// proxy clientConn <==> backendConn
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		io.Copy(clientConn, backendConn)
		clientConn.(*net.TCPConn).CloseWrite()
		wg.Done()
	}()
	go func() {
		io.Copy(backendConn, clientReader)
		backendConn.(*net.TCPConn).CloseWrite()
		wg.Done()
	}()

	wg.Wait()
}

func readClientHello(reader io.Reader) (*tls.ClientHelloInfo, error) {
	var hello *tls.ClientHelloInfo

	err := tls.Server(readOnlyConn{reader}, &tls.Config{
		GetConfigForClient: func(argHello *tls.ClientHelloInfo) (*tls.Config, error) {
			hello = new(tls.ClientHelloInfo)
			*hello = *argHello
			return nil, nil
		},
	}).Handshake()
	if hello == nil { // always fails
		return nil, err
	}

	return hello, nil
}

func peekClientHello(reader io.Reader) (*tls.ClientHelloInfo, io.Reader, error) {
	peekedBytes := new(bytes.Buffer)
	hello, err := readClientHello(io.TeeReader(reader, peekedBytes))
	if err != nil {
		return nil, nil, err
	}
	return hello, io.MultiReader(peekedBytes, reader), nil
}
