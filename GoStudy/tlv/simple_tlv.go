/**
 * @Author: Gong Yanhui
 * @Description: 简洁版TLV结构的编码解码功能
 * @Date: 2024-03-26 11:36
 */

package tlv

import (
	"fmt"
	tlv "github.com/yzw-trt/simple-tlv"
)

const (
	key1 uint8 = iota + 1
	key2
	key3
)

func SimpleTLV() {

	var (
		value1 = "name"
		value2 = "age"
		value3 = "address"
	)

	// 创建tlv对象
	c := tlv.NewContainer()
	// 添加tlv成员 key不能为0
	c.SetBytes(key1, []byte(value1))
	c.SetBytes(key2, []byte(value2))
	c.SetBytes(key3, []byte(value3))
	// 将tlv对象编码为字节流
	data := c.Bytes()
	fmt.Println(data)

	// 将字节流转换为tlv对象
	c, err := tlv.UnmarshalTLVContainer(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取成员
	if v, ok := c.GetBytes(key1); !ok || string(v) != value1 {
		fmt.Printf("key1 error: ok:%v, value:%s\n", ok, v)
	}
	if v, ok := c.GetBytes(key2); !ok || string(v) != value2 {
		fmt.Printf("key2 error: ok:%v, value:%s\n", ok, v)
	}
	if v, ok := c.GetBytes(key3); !ok || string(v) != value3 {
		fmt.Printf("key3 error: ok:%v, value:%s\n", ok, v)
	}
	fmt.Println("done!")
}
