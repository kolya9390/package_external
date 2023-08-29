package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) int {
    return int(math.Sqrt(x))
}