package main

import (
	"fmt"
	"sort"
)

/*
给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
输入: strs = [""]
输出: [[""]]

示例 3:
输入: strs = ["a"]
输出: [["a"]]
*/

/*
排序法
排序后的作为key，对应的字符串作为value数组
*/
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		sbyte := []byte(str)
		sort.Slice(sbyte, func(i, j int) bool {
			return sbyte[i] < sbyte[j]
		})
		s := string(sbyte)
		mp[s] = append(mp[s], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
