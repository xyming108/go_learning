package main

import (
	"fmt"
	"strconv"
)

/*
3-3-3-3-4-4-5-5-6-7-8-9-10-J-Q-K-A
4-5-6-7-8-8-8

结果：

9-10-J-Q-k-A
*/

func AA(s, ss string) (r string) {
	tempMap := make(map[int]int, len(s)+len(ss))
	for i := 3; i <= 14; i++ {
		tempMap[i] = 0
	}
	for i := 0; i < len(s); i++ {
		if i > 0 && i < len(s) && s[i] == '1' && s[i+1] == '0' {
			tempMap[10]++
		}
		data, _ := strconv.Atoi(string(s[i]))
		if s[i] != '-' && (data >= 3 && data <= 10) {
			tempMap[data]++
		}

		datas := string(s[i])
		if s[i] != '-' && datas == "J" {
			tempMap[11]++
		}
		if s[i] != '-' && datas == "Q" {
			tempMap[12]++
		}
		if s[i] != '-' && datas == "K" {
			tempMap[13]++
		}
		if s[i] != '-' && datas == "A" {
			tempMap[14]++
		}
	}
	for i := 0; i < len(ss); i++ {
		if i > 0 && i < len(s) && s[i] == '1' && s[i+1] == '0' {
			tempMap[10]++
		}
		data, _ := strconv.Atoi(string(ss[i]))
		if s[i] != '-' && (data >= 3 && data <= 10) {
			tempMap[data]++
		}

		datas := string(ss[i])
		if ss[i] != '-' && datas == "J" {
			tempMap[11]++
		}
		if ss[i] != '-' && datas == "Q" {
			tempMap[12]++
		}
		if ss[i] != '-' && datas == "K" {
			tempMap[13]++
		}
		if ss[i] != '-' && datas == "A" {
			tempMap[14]++
		}
	}

	var ps []int
	for i := 3; i <= 14; i++ {
		if tempMap[i] != 4 {
			ps = append(ps, i)
		}
	}

	var index, maxLen int
	count := 1
	for i := 0; i < len(ps)-1; i++ {
		if ps[i+1]-ps[i] == 1 {
			count++
		} else {
			count = 1
		}

		if count > maxLen {
			maxLen = count
			index = i - maxLen + 2
		}
	}

	var rr string
	for m := index; m < index+maxLen; m++ {
		if m != index+maxLen-1 {
			switch ps[m] {
			case 11:
				rr += "J-"
			case 12:
				rr += "Q-"
			case 13:
				rr += "k-"
			case 14:
				rr += "A-"
			default:
				rr += strconv.Itoa(ps[m]) + "-"
			}
		} else {
			switch ps[m] {
			case 11:
				rr += "J"
			case 12:
				rr += "Q"
			case 13:
				rr += "k"
			case 14:
				rr += "A"
			default:
				rr += strconv.Itoa(ps[m])
			}
		}

	}

	return rr
}

func main() {
	var s string
	var ss string
	fmt.Scan(&s)
	fmt.Scan(&ss)
	fmt.Println(AA(s, ss))
}
