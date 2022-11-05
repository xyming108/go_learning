package main

import (
	"fmt"
	"sort"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
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

提示：strs[i] 仅包含小写字母
*/

/*
排序法
*/
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string, 0)
	for _, v := range strs {
		strByte := []byte(v)
		sort.Slice(strByte, func(i, j int) bool {
			return strByte[i] < strByte[j]
		})
		strMap[string(strByte)] = append(strMap[string(strByte)], v)
	}
	result := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	anagrams := groupAnagrams(strs)
	fmt.Println(anagrams)
}
