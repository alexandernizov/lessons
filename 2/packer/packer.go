package packer

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrIncorrectInput = errors.New("incorrect input")
)

func UnpackString(input string) (string, error) {
	// Алгоритм:
	// 1. Пустая строка - пустая строка
	// 2. строка не начинается с цифры
	// 3. Разделяем строку по рунам
	// 4. Добавляем в буфер новые строчки, через repeat

	if input == "" {
		return "", nil
	}

	var res string
	var prev string
	for _, let := range input {
		if unicode.IsDigit(let) {
			if prev == "" {
				return "", ErrIncorrectInput
			}
			multiply := let - '0'
			res = res + strings.Repeat(prev, int(multiply))
			prev = ""
			continue
		}
		res = res + string(prev)
		prev = string(let)
	}
	res = res + prev
	return res, nil
}

// Через билдер строк - работает в 3 раза быстрее
func UnpackString2(input string) (string, error) {
	// Алгоритм:
	// 1. Пустая строка - пустая строка
	// 2. строка не начинается с цифры
	// 3. Разделяем строку по рунам
	// 4. Добавляем в буфер новые строчки, через repeat

	if input == "" {
		return "", nil
	}

	b := strings.Builder{}
	var prev rune = -1
	for _, let := range input {
		if unicode.IsDigit(let) {
			if prev == -1 {
				return "", ErrIncorrectInput
			}
			multiply := int(let - '0')
			for i := 0; i < int(multiply); i++ {
				b.WriteRune(prev)
			}
			prev = -1
			continue
		}
		if prev != -1 {
			b.WriteRune(prev)
		}
		prev = let
	}
	b.WriteRune(prev)
	return b.String(), nil
}
