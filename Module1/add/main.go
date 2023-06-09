package main

import "fmt"

func add(a, b []int32, p int) []int32 {
	// Найдем максимальное по длине число
	var maxNum, minNum []int32
	if len(a) >= len(b) {
		maxNum, minNum = a, b
	} else {
		maxNum, minNum = b, a
	}

	// В overNum будет храниться "переполнение" от сложения
	// предыдущих разрядов двух чисел
	var overNum, stepResult int32
	// В result будет храниться результат сложения двух чисел
	result := make([]int32, 0)

	for i := 0; i < len(maxNum); i++ {
		result = append(result, 0)

		if i < len(minNum) {
			stepResult = maxNum[i] + minNum[i]
		} else {
			stepResult = maxNum[i]
		}

		// Прибавим переполнение, полученное на пред. шагу
		stepResult += overNum
		result[i] += stepResult % int32(p)
		overNum = stepResult / int32(p)
	}

	// Если переполнение осталось после цикла
	if overNum != 0 {
		result = append(result, overNum)
	}

	return result
}

func main() {
	a := []int32{1, 1}
	b := []int32{1}
	c := add(a, b, 2)
	fmt.Println(c)
}
