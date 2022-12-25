/**
 * @Author: Gong Yanhui
 * @Description: 客户端如何发起 HTTP 请求	https://laravelacademy.org/post/20989
 * @Date: 2022-12-25 17:35
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
	可以通过 net/http 包里面的 Client 类提供的如下方法发起 HTTP 请求
func (c *Client) Do(req *Request) (*Response, error)
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
*/

func main() {
	// 初始化客户端请求对象
	req, err := http.NewRequest("GET", "https://baidu.com", nil)
	// 只有 POST、PUT、DELETE 之类的请求才需要设置请求实体，对于 HEAD、GET 而言，传入 nil 即可
	if err != nil {
		fmt.Printf("请求初始化失败：%v", err)
		return
	}
	// 添加自定义请求头
	req.Header.Add("Custom-Header", "Custom-Value")
	// ... 其它请求头配置
	client := &http.Client{
		// ... 设置客户端属性
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("客户端发起请求失败：%v", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
