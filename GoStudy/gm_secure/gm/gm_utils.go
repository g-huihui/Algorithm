/**
 * @Author: Gong Yanhui
 * @Description: 国密工具类
 * @Date: 2024-01-08 14:34
 */

package gm

import (
	"encoding/pem"
	"fmt"
	"github.com/tjfoc/gmsm/x509"
	"io/ioutil"
	"time"
)

// CheckGmCert 检查国密证书有效期
func CheckGmCert() {
	// 从文件中读取国密证书内容
	certBytes, err := ioutil.ReadFile("/Users/gongyanhui/Downloads/国密cert/sm2.1.sig.crt.pem")
	if err != nil {
		fmt.Println("无法读取证书文件:", err)
		return
	}

	// 解析 PEM 格式的证书数据
	block, _ := pem.Decode(certBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		fmt.Println("无法解析证书")
		return
	}

	// 解析证书
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Println("无法解析证书:", err)
		return
	}

	// 检查证书的有效期
	now := time.Now()
	if now.Before(cert.NotBefore) {
		fmt.Println("证书尚未生效")
	} else if now.After(cert.NotAfter) {
		fmt.Println("证书已过期")
	} else {
		fmt.Println("证书有效")
	}

	// 输出证书的有效期信息
	fmt.Println("证书有效期从:", cert.NotBefore)
	fmt.Println("证书有效期至:", cert.NotAfter)
}
