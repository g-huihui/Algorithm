/**
 * @Author: Gong Yanhui
 * @Description: 私钥签名 公钥校验
 * @Date: 2023-03-29 14:46
 */

package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//RsaMethod()

	//generatePemAndSave2Local()

	readLocalPEM2SignAndVerify()
}

// RsaMethod 使用rsa库生成私钥数据 私钥签名 公钥校验样例
func RsaMethod() {
	// 示例数据
	plainText := []byte("Hello, world!")

	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("私钥生成失败:", err)
		return
	}

	// 获取私钥的字节表示形式
	//privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	//fmt.Println("私钥:", privateKeyBytes)

	hashed := sha256.Sum256(plainText)
	// 使用私钥进行加密
	ciphertext, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}

	fmt.Printf("加密后的数据: %x\n", ciphertext)

	// 获取公钥的字节表示形式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("公钥获取失败:", err)
		return
	}

	// 将公钥字节表示形式转换为PEM格式
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	// 解码公钥PEM
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil || block.Type != "PUBLIC KEY" {
		fmt.Println("公钥解码失败")
		return
	}

	// 解析公钥
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("公钥解析失败:", err)
		return
	}

	// 将接口转换为公钥类型
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		fmt.Println("公钥类型转换失败")
		return
	}

	hash := sha256.Sum256(plainText)
	// 使用公钥进行解密
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], ciphertext)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}

	fmt.Println("解密成功")
}

// 创建私钥和公钥文件并保存到本地
func generatePemAndSave2Local() {
	// 生成RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Failed to generate key pair:", err)
		return
	}

	// 保存私钥到文件
	file, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// 使用PKCS1格式编码私钥
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	// 将私钥进行PEM编码
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	// 写入文件
	err = pem.Encode(file, privateBlock)
	if err != nil {
		fmt.Println("Failed to save private key:", err)
		return
	}

	// 保存公钥到文件
	file, err = os.Create("public_key.pem")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// 使用PKIX格式编码公钥
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		fmt.Println("Failed to get public key:", err)
		return
	}
	// 将公钥进行PEM编码
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	// 写入文件
	err = pem.Encode(file, publicBlock)
	if err != nil {
		fmt.Println("Failed to save public key:", err)
		return
	}
}

// 从本地读取私钥和公钥 验证用私钥签名数据 公钥验签数据
func readLocalPEM2SignAndVerify() {
	plainText := []byte("Hello, world!")

	// 使用私钥签名数据
	privateKey, err := loadPrivateKeyFromFile("private_key.pem")
	if err != nil {
		fmt.Println("Failed to load private key:", err)
		return
	}
	hashed := sha256.Sum256(plainText)
	// 使用私钥进行加密
	ciphertext, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println("加密失败:", err)
		return
	}

	// 使用公钥验签数据
	publicKey, err := loadPublicKeyFromFile("public_key.pem")
	if err != nil {
		fmt.Println("Failed to load public key:", err)
		return
	}
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], ciphertext)
	if err != nil {
		return
	}

	fmt.Println("校验通过")
}

func loadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解码PEM块
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA private key")
	}

	// 解析PKCS1格式的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func loadPublicKeyFromFile(filename string) (*rsa.PublicKey, error) {
	// 读取文件内容
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 解码PEM块
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// 解析PKIX格式的公钥
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 转换为RSA公钥类型
	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return publicKey, nil
}
