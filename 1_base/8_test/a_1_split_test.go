package main

import (
	"reflect"
	"testing"
)

/**
  Author: xym
  Date: 2021/4/11 12:31
  Project: Go_Learning
  Description:
*/
/*
Golang单元测试对文件名和方法名，参数都有很严格的要求。

    1、文件名必须以xx_test.go命名
    2、方法必须是Test[^a-z]开头
    3、方法参数必须 t *testing.T
    4、使用go test执行单元测试
*/

func TestSplit(t *testing.T) {
	/*got := Split("i:love:you", ":")
	want := []string{"i", "love", "you"}*/

	/*got := Split("我爱你", "爱")
	want := []string{"我", "你"}*/

	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"simple":     {input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"multi_sep":  {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"multi_sep2": {input: "abcdefg", sep: "cd", want: []string{"ab", "efg"}},
		"chinese":    {input: "代价代价返回代价代价", sep: "返回", want: []string{"代价代价", "代价代价"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %s failed, want:%v got:%v", name, tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	// b.N不是固定的数
	for i := 0; i < b.N; i++ {
		Split("可对角化大", "对角")
	}
}
