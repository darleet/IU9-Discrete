package main

import (
	"fmt"
)

func encode(utf32 []rune) []byte {
	// Слайс для хранения интерпретированного текста
	utf8 := make([]byte, 0)

	for _, symbol := range utf32 {
		// Определим количество байт в utf-8
		var octets int
		switch {
		case symbol >= 1<<16:
			octets = 4
		case symbol >= 1<<11:
			octets = 3
		case symbol >= 1<<7:
			octets = 2
		default:
			octets = 1
		}

		// Слайс для хранения интерпретации символа в []byte
		result := make([]byte, 0)

		// Пройдемся по разрядам в двоичной записи
		for i := 0; i < octets; i++ {
			// Добавим в начало предыдущего байта 10...
			if len(result) > 0 {
				result[0] += 1 << 7
			}
			// Добавим новый байт
			result = append(make([]byte, 1), result...)
			// В k будет храниться 2 в степени нужного разряда
			var k byte = 1
			for j := 0; j < 6; j++ {
				result[0] += byte(symbol&1) * k
				k = k << 1
				symbol = symbol >> 1
			}
		}

		// Если в результате всего один октет
		if octets == 1 {
			result[0] += byte(symbol) * 1 << 6
		} else {
			var k byte = 1 << 7
			for i := 0; i < octets; i++ {
				result[0] += k
				k = k >> 1
			}
		}

		// Добавим результат для символа в слайс
		// со всеми результатами
		utf8 = append(utf8, result...)
	}

	return utf8
}

func decode(utf8 []byte) []rune {
	// Слайс для хранения интерпретированного текста
	utf32 := make([]rune, 0)
	octetCounter := 0

	for _, element := range utf8 {
		if element&(1<<7) == 0 { // Если байт начинается с нуля
			utf32 = append(utf32, rune(element))
		} else if element&(1<<7) == 1<<7 && element&(1<<6) == 0 { // Если байт начинается с 10
			interpretation := rune(element) & (1<<7 - 1) << (octetCounter * 6)
			utf32[len(utf32)-1] += interpretation
			octetCounter--
		} else { // Если это начальный байт
			// k будет указывать на ноль после единиц
			k := 5
			// будем уменьшать k пока не встретим ноль
			for ; rune(element)&(1<<k) != 0; k-- {
			}
			// octetCounter будет хранить количество оставшихся октет
			octetCounter = 6 - k
			interpretation := rune(element) & (1<<k - 1) << (octetCounter * 6)
			utf32 = append(utf32, interpretation)
			octetCounter--
		}
	}

	return utf32
}

func main() {
	fmt.Print([]byte("𐍈"))
	fmt.Print([]rune("𐍈"))
	fmt.Print(encode([]rune("𐍈")))
	fmt.Println(decode([]byte("𐍈")))

	fmt.Print([]byte("$"))
	fmt.Print([]rune("$"))
	fmt.Print(encode([]rune("$")))
	fmt.Println(decode([]byte("$")))

	fmt.Print([]byte("¢"))
	fmt.Print([]rune("¢"))
	fmt.Print(encode([]rune("¢")))
	fmt.Println(decode([]byte("¢")))

	fmt.Print([]byte("€"))
	fmt.Print([]rune("€"))
	fmt.Print(encode([]rune("€")))
	fmt.Println(decode([]byte("€")))
}
