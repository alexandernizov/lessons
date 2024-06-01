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

	if input == "" {
		return "", nil
	}

	b := strings.Builder{}

	//Не пройдет, если в строке будет русская буква последней, по-хорошему бы сделать это так же как в запаковке строк:
	//строка в массив рун, и дальше уже получать предыдущие, текущие значения по рунам. Но пока тесты проходят - не переделываю))

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

// Дополнительное задание: поддержка экранирования через \
func UnpackStringSlash(input string) (string, error) {

	if input == "" {
		return "", nil
	}

	b := strings.Builder{}
	// -1 - означает, что предыдущий символ пустой, -2 - что предыдущий символ - это обратный слэш
	var prev rune = -1

	//Обходим строку побуквенно: если текущий символ - это цифра - то воспринимаем его как множитель предыдущего. Если
	//это не цифра - то записываем предыдущий символ без множителя.
	for _, let := range input {
		//Если предыдущий символ == обратный слеш - то это экранироваие текущего символа. Возможно экранировать только цифры или
		//обратный слэш
		if prev == 92 {
			if unicode.IsDigit(let) {
				prev = let
				continue
			} else if let == 92 {
				prev = -2
				continue
			} else {
				return "", ErrIncorrectInput
			}
		}
		if unicode.IsDigit(let) {
			if prev == -1 {
				return "", ErrIncorrectInput
			}
			multiply := int(let - '0')
			for i := 0; i < int(multiply); i++ {
				if prev == -2 {
					prev = 92
				}
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
	if prev != -1 {
		b.WriteRune(prev)
	}
	return b.String(), nil
}

func PackString(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	runes := []rune(input)

	b := strings.Builder{}
	var prev rune = -1
	count := 0

	for i := 0; i < len(runes)+1; i++ {
		var cur rune
		if i != len(runes) {
			cur = runes[i]
		}
		if prev == -1 {
			prev = cur
			count = 1
			continue
		}
		if prev == cur {
			count++
			continue
		}
		if prev != cur {
			b.WriteRune(prev)
			if count > 1 {
				b.WriteRune('0' + rune(count))
			}
			count = 1
			prev = cur
		}
	}

	return b.String(), nil
}
