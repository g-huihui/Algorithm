/**
 * @Author: Gong Yanhui
 * @Description: 生成RSA证书和私钥
 * @Date: 2023-06-20 14:59
 */

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {

	generateCertAndPrivateKeyThenSave2Local()
}

// 生成证书和私钥并保存到本地
func generateCertAndPrivateKeyThenSave2Local() {
	// 创建 RSA 密钥对
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate RSA key pair:", err)
		return
	}

	// 创建自签名证书
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "example.com"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privKey.PublicKey, privKey)
	if err != nil {
		fmt.Println("Failed to create self-signed certificate:", err)
		return
	}

	// 将私钥保存到文件
	privKeyFile, err := os.Create("private.key")
	if err != nil {
		fmt.Println("Failed to create private key file:", err)
		return
	}
	defer func() {
		_ = privKeyFile.Close()
	}()

	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}
	err = pem.Encode(privKeyFile, privKeyPEM)
	if err != nil {
		fmt.Println("Failed to write private key to file:", err)
		return
	}

	fmt.Println("Private key saved to private.key")

	// 将证书保存到文件
	certFile, err := os.Create("certificate.crt")
	if err != nil {
		fmt.Println("Failed to create certificate file:", err)
		return
	}
	defer func() {
		_ = certFile.Close()
	}()

	certPEM := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	err = pem.Encode(certFile, certPEM)
	if err != nil {
		fmt.Println("Failed to write certificate to file:", err)
		return
	}

	fmt.Println("Certificate saved to certificate.crt")
}
