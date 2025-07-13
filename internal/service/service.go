package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvStr(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("input string is empty")
	}

	if isMorseCode(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}

func isMorseCode(input string) bool {
	allowed := ".- "
	disallowed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZабвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

	return strings.ContainsAny(input, allowed) &&
		!strings.ContainsAny(input, disallowed) &&
		len(strings.TrimSpace(input)) > 0
}
