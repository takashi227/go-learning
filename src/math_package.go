package main

import (
	"fmt"
	"math"
)

func main() {
	// 数学的な定数
	// 円周率
	fmt.Println(math.Pi)

	// 2の平方根
	fmt.Println(math.Sqrt2)

	// 2の自然対数
	fmt.Println(math.Log2E)

	// 10の自然対数
	fmt.Println(math.Log10E)

	// 自然対数の底
	fmt.Println(math.Ln2)

	// 10の自然対数
	fmt.Println(math.Ln10)

	// 数学的な関数
	// 絶対値
	fmt.Println(math.Abs(-10))

	// 最大値
	fmt.Println(math.Max(10, 20))

	// 最小値
	fmt.Println(math.Min(10, 20))

	// 累乗
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Pow(2, 0))

	// 整数値に関する定数
	// float32の最大値
	fmt.Println(math.MaxFloat32)
	// float32の最小値(0以外)
	fmt.Println(math.SmallestNonzeroFloat32)
	// float64の最大値
	fmt.Println(math.MaxFloat64)
	// float64の最小値(0以外)
	fmt.Println(math.SmallestNonzeroFloat64)
	// int8の最大値
	fmt.Println(math.MaxInt8)
	// int8の最小値
	fmt.Println(math.MinInt8)
	// int16の最大値
	fmt.Println(math.MaxInt16)
	// int16の最小値
	fmt.Println(math.MinInt16)
	// int32の最大値
	fmt.Println(math.MaxInt32)
	// int32の最小値
	fmt.Println(math.MinInt32)
	// int64の最大値
	fmt.Println(math.MaxInt64)
	// int64の最小値
	fmt.Println(math.MinInt64)
	// intの最大値
	fmt.Println(math.MaxInt)
	// intの最小値
	fmt.Println(math.MinInt)

	// 小数点以下の切り捨て
	fmt.Println(math.Trunc(3.14))  // 3
	fmt.Println(math.Trunc(-3.14)) // -3

	// 小数点以下の切り上げ
	fmt.Println(math.Ceil(3.14))  // 4
	fmt.Println(math.Ceil(-3.14)) // -3

	// 小数点以下の切り下げ
	fmt.Println(math.Floor(3.14))  // 3
	fmt.Println(math.Floor(-3.14)) // -4

}
