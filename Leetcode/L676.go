/**
 * @Author: Gong Yanhui
 * @Description: 676. 实现一个魔法字典
 * @Date: 2024-08-12 21:36
 */

package main

type MagicDictionary struct {
	words []string
}

func MagicDictionaryConstructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.words = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
next:
	for _, word := range this.words {
		if len(word) != len(searchWord) {
			continue
		}
		diff := false
		for i := 0; i < len(word); i++ {
			if word[i] != searchWord[i] {
				if diff {
					continue next
				}
				diff = true
			}
		}
		if diff {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
