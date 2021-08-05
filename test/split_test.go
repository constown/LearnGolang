package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want:%v, but got:%v", want, ret)
	}
}

// 测试组

func TestGroupSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := []testCase{
		testCase{"babcbde", "b", []string{"", "a", "c", "de"}},
		testCase{"A:B:C", ":", []string{"A", "B", "C"}},
		testCase{"abcdef", "bc", []string{"a", "def"}},
		testCase{"沙河有沙有河水", "有", []string{"沙河", "沙", "河水"}},
	}
	for _, tc := range testGroup {
		got := Split(tc.str, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("want:%v, but got:%v", tc.want, got)
		}
	}
}

// 子测试
// 跑单个在终端中输入：
// go test -run=TestSonSplit/case1
func TestSonSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": testCase{"babcbde", "b", []string{"", "a", "c", "de"}},
		"case2": testCase{"A:B:C", ":", []string{"A", "B", "C"}},
		"case3": testCase{"abcdef", "bc", []string{"a", "def"}},
		"case4": testCase{"沙河有沙有河水", "有", []string{"沙河", "沙", "河水"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%v, but got:%v", tc.want, got)
			}
		})
	}
}

// 基准测试
// 终端中运行
//go test -bench=Split
//go test -bench=Split -benchmem

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
