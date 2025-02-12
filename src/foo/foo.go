package foo

const (
	Max = 100 // 1文字目が大文字なので他のパッケージからも参照可能
	min = 1   // 1文字目が小文字なので他のパッケージからは参照できない
)

func ReturnMin() int {
	return min
}
