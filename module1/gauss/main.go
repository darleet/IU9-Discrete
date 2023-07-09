package main

import (
	"fmt"
	"strconv"
)

// fraction описывает рациональную дробь
type fraction struct {
	numerator   int
	denominator int
}

// abs находит модуль числа
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// simplify упрощает дробь
func simplify(f *fraction) *fraction {
	if f.numerator < 0 && f.denominator < 0 ||
		f.numerator > 0 && f.denominator < 0 {
		f.numerator, f.denominator = -f.numerator, -f.denominator
	}

	if f.numerator != 0 {
		a, b := abs(f.numerator), abs(f.denominator)
		for a != b {
			if a > b {
				a = a - b
			} else {
				b = b - a
			}
		}

		f.numerator /= a
		f.denominator /= b
	} else {
		f.denominator = 1
	}

	return f
}

// Subtract вычитает из дроби другую дробь
func (f *fraction) Subtract(other *fraction, repeats *fraction) {
	subtractedNumerator := other.numerator * f.denominator * repeats.numerator
	subtractedDenominator := other.denominator * repeats.denominator
	f.numerator *= subtractedDenominator
	f.denominator *= subtractedDenominator
	f.numerator -= subtractedNumerator
	*f = *simplify(f)
}

// subtractRow вычитает из одной строки другую строку
func subtractRow(this, other []*fraction, repeats *fraction) {
	for i := range this {
		this[i].Subtract(other[i], repeats)
	}
}

// Divide делит одну дробь на другую
func (f *fraction) Divide(other *fraction) {
	f.numerator *= other.denominator
	f.denominator *= other.numerator
	*f = *simplify(f)
}

// divideRow делит одну строку на другую
func divideRow(row []*fraction, rowDivider *fraction) {
	for _, el := range row {
		if el.numerator != 0 && rowDivider.numerator != 0 {
			el.Divide(rowDivider)
		}
	}
}

// swapRows меняет местами две строки
func swapRows(a, b []*fraction) {
	for i := range a {
		*a[i], *b[i] = *b[i], *a[i]
	}
}

// findSwap ищет замену для строки среди последующих строк
func findSwap(index int, matrix [][]*fraction) bool {
	for i := index + 1; i < len(matrix); i++ {
		if matrix[i][index].numerator != 0 {
			swapRows(matrix[index], matrix[i])
			return true
		}
	}
	return false
}

// solveGauss ищет решение СЛАУ методом Гаусса
func solveGauss(matrix [][]*fraction) []*fraction {
	// Получаем первый треугольник из нулей
	for i := range matrix {
		// проверим, что на диагонали не стоит нуль, а если стоит -
		// попробуем найти замену среди последующих строк
		if !(matrix[i][i].numerator == 0 && !findSwap(i, matrix)) {
			rowDivider := &fraction{
				numerator:   matrix[i][i].numerator,
				denominator: matrix[i][i].denominator,
			}
			// приведем элемент на диагонали к единице
			divideRow(matrix[i], rowDivider)
			for j := i + 1; j < len(matrix); j++ {
				repeats := &fraction{
					numerator:   matrix[j][i].numerator,
					denominator: matrix[j][i].denominator,
				}
				// вычтем строку с единицей необходимое кол-во раз
				subtractRow(matrix[j], matrix[i], repeats)
			}
		}
	}

	// Получаем второй треугольник из нулей
	for i := range matrix {
		for j := i + 1; j < len(matrix); j++ {
			repeats := &fraction{
				numerator:   matrix[i][j].numerator,
				denominator: matrix[i][j].denominator,
			}
			// вычтем строку с единицей необходимое кол-во раз
			subtractRow(matrix[i], matrix[j], repeats)
		}
	}

	results := make([]*fraction, 0)

	for i, row := range matrix {
		if row[i].numerator == 0 {
			return nil
		}
		results = append(results, simplify(row[len(row)-1]))
	}

	return results
}

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	matrix := make([][]*fraction, 0)

	for i := 0; i < n; i++ {
		row := make([]*fraction, 0)
		for j := 0; j < n+1; j++ {
			var x int
			_, _ = fmt.Scanf("%d", &x)
			row = append(row, &fraction{x, 1})
		}
		matrix = append(matrix, row)
	}

	result := solveGauss(matrix)

	if result == nil {
		fmt.Println("No solution")
	} else {
		for _, res := range result {
			output := strconv.Itoa(res.numerator) + "/" +
				strconv.Itoa(res.denominator)
			fmt.Println(output)
		}
	}
}
