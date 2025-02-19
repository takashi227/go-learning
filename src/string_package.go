package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列を結合する
	s1 := strings.Join([]string{"a", "b", "c"}, "-")
	s2 := strings.Join([]string{"1", "2", "3"}, "")
	fmt.Println(s1)
	fmt.Println(s2)

	// 文字列に含まれる部分文字列を検索する
	i1 := strings.Index("abcde", "b")
	i2 := strings.Index("abcde", "d")
	fmt.Println(i1)
	fmt.Println(i2)

	i3 := strings.LastIndex("abcdeabcde", "cd")
	fmt.Println(i3)

	i4 := strings.IndexAny("abc", "abc")
	i5 := strings.LastIndexAny("abc", "abc")
	fmt.Println(i4)
	fmt.Println(i5)

	b1 := strings.HasPrefix("abc", "a")
	b2 := strings.HasSuffix("abc", "c")
	fmt.Println(b1, b2)

	b3 := strings.Contains("abc", "b")
	b4 := strings.ContainsAny("abcde", "bd")
	fmt.Println(b3, b4)

	i6 := strings.Count("abcabcabc", "b")
	i7 := strings.Count("abcabcabc", "") // 文字列の長さ+1
	fmt.Println(i6, i7)

	// 文字列を繰り返して結合する
	s3 := strings.Repeat("abc", 3)
	s4 := strings.Repeat("abc", 0)
	fmt.Println(s3)
	fmt.Println(s4)

	// 文字列を置換する
	s5 := strings.Replace("abcabcabc", "b", "x", 2)
	s6 := strings.Replace("abcabcabc", "b", "x", -1)
	fmt.Println(s5)
	fmt.Println(s6)

	// 文字列を分割する
	s7 := strings.Split("a,b,c", ",")
	fmt.Println(s7)
	s8 := strings.SplitAfter("a,b,c", ",")
	fmt.Println(s8)
	s9 := strings.SplitN("a,b,c", ",", 2)
	fmt.Println(s9)
	s10 := strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println(s10)

	// 文字列の前後の空白を削除する
	s11 := strings.TrimSpace("  -  abc  -  ")
	fmt.Println(s11)
	s12 := strings.TrimSpace("　　abc　　") // 全角スペース
	fmt.Println(s12)

	// 文字列からスペースで区切られたフィールドを取り出す
	s13 := strings.Fields(" a b c ")
	fmt.Println(s13)
	fmt.Println(s13[1])

	// 大文字・小文字に変換する
	s14 := strings.ToUpper("abc")
	s15 := strings.ToLower("ABC")
	fmt.Println(s14)
	fmt.Println(s15)
}
