/**
 * @Author: Gong Yanhui
 * @Description: 2296. 设计一个文本编辑器
 * @Date: 2025-02-27 12:50
 */

package main

// 当前解法 超时

type TextEditor struct {
	index int    // 当前光标位置 指向当前字符串下一位 比如 "abc" index:3 也代表左边的长度
	text  string // 文本内容
}

func TextEditorConstructor() TextEditor {
	return TextEditor{
		index: 0,
		text:  "",
	}
}

func (this *TextEditor) AddText(text string) {
	this.text = this.text[:this.index] + text + this.text[this.index:]
	this.index += len(text)
}

func (this *TextEditor) DeleteText(k int) int {
	if k > this.index {
		k = this.index
	}
	this.text = this.text[:this.index-k] + this.text[this.index:]
	this.index -= k

	return k
}

func (this *TextEditor) CursorLeft(k int) string {
	if k > this.index {
		k = this.index
	}
	this.index -= k

	return this.getLeft()
}

func (this *TextEditor) CursorRight(k int) string {
	if k > len(this.text)-this.index {
		k = len(this.text) - this.index
	}
	this.index += k

	return this.getLeft()
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */

func (this *TextEditor) getLeft() string {
	var resLen int
	if this.index > 10 {
		resLen = 10
	} else {
		resLen = this.index
	}

	return this.text[this.index-resLen : this.index]
}
