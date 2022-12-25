/**
 * @Author: Gong Yanhui
 * @Description: Dial 函数的底层实现及超时处理	https://laravelacademy.org/post/20970
 * @Date: 2022-12-25 17:27
 */

package main

// 在使用 Dial 函数建立网络连接时，可以使用 net 包提供的 DialTimeout 函数主动传入额外的超时参数来建立连接
/*
	和 Dial 函数调用一样，只是设置了超时字段而已，如果使用 Dial 函数，
	默认会通过操作系统提供的机制来处理连接超时，对于 TCP 连接，通常是 3 分钟左右
*/

/**
SetDeadline(t time.Time) error
SetReadDeadline(t time.Time) error
SetWriteDeadline(t time.Time) error
	我们可以通过 SetDeadline 设置统一的读写超时时间，
	也可以通过 SetReadDeadline 和 SetWriteDeadline 分别设置读超时和写超时。
	注意这三个方法传入的都是绝对时间值，而不是相对时间长度
err = conn.SetDeadline(time.Now().Add(5 * time.Second))
*/

/*
	// 可以通过 ParseIP 函数验证 IP 地址的有效性
	func net.ParseIP()
	// 根据域名查找对应 IP 地址
	func ResolveIPAddr(net, addr string) (*IPAddr, error)
	func LookupHost(name string) (cname string, addrs []string, err error)
*/
