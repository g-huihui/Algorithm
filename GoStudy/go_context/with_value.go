/**
 * @Author: Gong Yanhui
 * @Description: WithValue 传递上下文 携带数据
 * @Date: 2023-08-05 15:11
 */

package main

import (
	"context"
	"fmt"
	"time"
)

const key = "log_id"

func NewRequestId() string {
	// return strings.Replace(uuid.New().String(), "-", "", -1)
	return "my_request_id"
}

func NewContextWithTraceId() context.Context {
	return context.WithValue(context.Background(), key, NewRequestId())
}

func GetContextValue(ctx context.Context, key string) string {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return value
}

func PrintLog(ctx context.Context, msg string) {
	fmt.Printf("time:%s, log_id:%s, msg:%s\n", time.Now().Format("2006-01-02 15:04:05"), GetContextValue(ctx, key), msg)
}

func ProcessEnter(ctx context.Context) {
	PrintLog(ctx, "some error info...")
}

func main() {
	ProcessEnter(NewContextWithTraceId())
}

/**
 * 1. 不建议使用context值传递关键参数 关键参数应该显示的声明出来 不应该隐式处理 context中最好是携带签名 trace_id这类值
 * 2. 因为携带value也是key-value的形式 为了避免context因多个包同时使用context而带来冲突 key建议采用内置类型
 * 3. context传递的数据中key-value都是interface类型 这种类型编译期无法确定类型 所以不是很安全 所以在类型断言时别忘了保证程序的健壮性
 */
