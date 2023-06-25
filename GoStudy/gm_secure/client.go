/**
 * @Author: Gong Yanhui
 * @Description:
 * @Date: 2023-06-20 18:00
 */

package main

func main() {
	//caCert, err := ioutil.ReadFile("certificate.crt")
	//if err != nil {
	//	log.Fatalf("Failed to load root certificate: %v", err)
	//}
	//
	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)
	//
	//// 创建国密配置
	//config := &gmtls.Config{
	//	RootCAs: caCertPool,
	//	CipherSuites: []uint16{
	//		gmtls.GMTLS_ECDHE_SM2_WITH_SM4_SM3,
	//	},
	//}
	//
	//conn, err := gmtls.Dial("tcp", "127.0.0.1:8080", config)
	//if err != nil {
	//	fmt.Println("Dial err:", err)
	//	return
	//}
	//defer conn.Close()
	//
	//buf := make([]byte, 1024)
	//_, err = conn.Read(buf)
	//if err != nil {
	//	fmt.Println("Read err:", err)
	//	return
	//}
	//
	//fmt.Println(string(buf))
}
