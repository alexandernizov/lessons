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
