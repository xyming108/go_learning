package main

import "fmt"

/*
给定一个字符串s和一个字符串数组words。words中所有字符串 长度相同。
s中的 串联子串 是指一个包含words中所有字符串以任意顺序排列连接起来的子串。
例如，如果words = ["ab","cd","ef"]， 那么"abcdef"，"abefcd"，"cdabef"，"cdefab"，
"efabcd"，和"efcdab" 都是串联子串。"acdbef" 不是串联子串，因为他不是任何words排列的连接。
返回所有串联字串在s中的开始索引。你可以以 任意顺序 返回答案。

示例 1：
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
输出顺序无关紧要。返回 [9,0] 也是可以的。

示例 2：
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
所以我们返回一个空数组。

示例 3：
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。

提示：
1 <= s.length <= 104
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] 和 s 由小写英文字母组成
*/

/*
map，自己写的，时间复杂度偏高
时间复杂度：O(ls*m*n), ls 是输入 s 的长度，m 是 word 的单词, n 是 words 中每个单词的长度
空间复杂度：O(m*n), m 是 word 的单词数，n 是 words 中每个单词的长度
*/
func findSubstring(s string, words []string) []int {
	wordFirstMap := make(map[string]bool, 0)
	for _, v := range words {
		firstStr := string(v[0])
		wordFirstMap[firstStr] = true
	}
	wordLength := len(words[0])
	amount := len(words)
	length := amount * wordLength
	result := make([]int, 0)
	for i, v := range s {
		if i+length > len(s) {
			break
		}
		vStr := string(v)
		if wordFirstMap[vStr] && i+length <= len(s) {
			tempStr := s[i : i+length]
			tempWords := make([]string, 0, amount)
			tempWordMap := make(map[string]int, 0)
			for _, u := range words {
				tempWordMap[u] += 1
			}
			for j := 0; j < length; {
				value := tempStr[j : j+wordLength]
				if tempWordMap[value] > 0 {
					tempWords = append(tempWords, value)
					tempWordMap[value]--
				} else {
					break
				}
				j += wordLength
			}
			if len(tempWords) == amount {
				result = append(result, i)
			}
		}
	}
	return result
}

/*
滑动窗口
时间复杂度：O(ls×n), ls 是输入 s 的长度，n 是 words 中每个单词的长度
空间复杂度：O(m×n), m 是 word 的单词数，n 是 words 中每个单词的长度
*/
func findSubstring2(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	for i := 0; i < n && i+m*n <= ls; i++ {
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{"word", "good", "best", "good"}
	fmt.Println(findSubstring(s, words))
	fmt.Println(findSubstring2(s, words))
}
