/**
 * @Author: Gong Yanhui
 * @Description: 测试用例
 * @Date: 2023-11-03 14:07
 */

package main

import "testing"

const pubKey = "MzQ1NDA3MTcxMzQzMDEwMjcyOTgzNjAzNTU0NzA3Mjk5MzEwMDIzMjg2MDYyNDM2MDUxODQwNTYyNzQ3NjIwNjM0NDE1NjA4MjkwNDgrNTUwOTQ2MzQ1Mzg3MDEyODM4NTU5NTgwOTM0MDc3Mjg5NjU2NTI3MjE5MDkzNDk4Mjk4NzUxNzM1MDc1MjkzMDkxODMzMjQzNTAyNzQ="
const priKey = "MHcCAQEEIJtQM5xn5CU0m/4qXGzHowyzg1xyIQsim5lEf9pFFN4XoAoGCCqGSM49AwEHoUQDQgAETF1aZG/htJJEAvVLXUDz2P9jTii0Oi7tnwjrfjaTfHh5znVFAh2ee0vypXZWz8VVVWVyIlARoiYmrM7FoJJnQg=="
const content = "This is a test message"
const signature = "3130343137323338333733393831353733303434373937393136333339353433313736383838383634343136333535313438353839343934353431333631383138333531353838353031353034392b3432303035363638363731383131303432313233373330323236363136363038323634373235383035383131373730393632363038323238343235353837313734353134303032313631303436"

func TestGenKeyPair(t *testing.T) {
	privateKey, publicKey, e := GenKeyPair()
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("pubKey:", publicKey)
	t.Log("priKey:", privateKey)
}

func TestBuildPublicKey(t *testing.T) {
	publicKey, e := BuildPublicKey(pubKey)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("publicKey:", publicKey)
}

func TestBuildPrivateKey(t *testing.T) {
	privateKey, e := BuildPrivateKey(priKey)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("privateKey:", privateKey)
}

func TestSign(t *testing.T) {
	sign, e := Sign(content, priKey)
	if e != nil {
		t.Error(e)
	} else {
		t.Log("signature:", sign)
	}
}

func TestVerifySign(t *testing.T) {
	verify := VerifySign(content, signature, pubKey)
	t.Log("verify:", verify)
}
