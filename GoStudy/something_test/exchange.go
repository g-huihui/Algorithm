/**
 * @Author: Gong Yanhui
 * @Description:
 * @Date: 2024-01-09 14:42
 */

package something_test

func ExchangeInt(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func ExchangeByte(a, b *byte) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func ExchangeRune(a, b *rune) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
