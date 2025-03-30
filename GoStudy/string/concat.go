/**
 * @Author: Gong Yanhui
 * @Description: 字符串拼接
 * @Date: 2025-03-30 15:03
 */

package string

import (
	"bytes"
	"strings"
)

// 使用 + 拼接字符串性能非常差 这是由于string是不可修改的
// 在使用 + 进行拼接字符串 每次都会产生申请空间 拼接 复制等操作 数据量大的情况下非常消耗资源和性能

// bytes.Buffer 是一个缓冲byte类型的缓冲器存放着都是byte 底层就是一个 []byte
func appendWithBytesBuffer() string {
	var buffer bytes.Buffer

	buffer.Write([]byte("Hello"))
	buffer.WriteString("World")
	// func NewBuffer(buf []byte) *Buffer
	// func NewBufferString(s string) *Buffer

	return buffer.String()
}

// strings.Builder 的方法和bytes.Buffer的方法的命名几乎一致
func appendWithStringBuilder() string {
	var builder strings.Builder

	builder.Write([]byte("Hello"))
	builder.WriteString("World")

	return builder.String()
}

/*
相同点:
1. Buffer和Builder底层都是采用[]byte数组进行装载数据
异同点:
1. Builder直接append的方式向字节数组后添加字符串
2. String() Buffer的string是一种强转 我们知道在强转的时候是需要进行申请空间并拷贝的 而Builder只是指针的转换 *(*string)(unsafe.Pointer(&b.buf))
*/
