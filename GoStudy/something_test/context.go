/**
 * @Author: Gong Yanhui
 * @Description: 测试context
 * @Date: 2024-01-08 14:56
 */

package something_test

import (
	"context"
	"fmt"
	"time"
)

func TestContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done(): // 只要调用了cancel 就会一直执行这个case
				fmt.Println("done")
			default:
				fmt.Println("default")
			}
		}
	}(ctx)

	cancel()
	time.Sleep(3 * time.Second)
}
