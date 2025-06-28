package util

import (
	"fmt"

	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextTaskDate(now time.Time, dStart string, repeatRule string) (string, error) {
	var nextDate time.Time

	startDate, err := time.Parse(DateFormat, dStart)
	if err != nil {
		return "", fmt.Errorf("ошибка NextTaskDate - неверный формат даты")
	}

	if !validateStringRule(repeatRule) {
		return "", fmt.Errorf("ошибка NextTaskDate - неверный формат правила повторения задач")
	}

	nextDate = startDate
	period := repeatRule[0]

	if period == 'd' {
		parts := strings.Split(repeatRule, " ")
		interval, _ := strconv.Atoi(parts[1])

		for {
			nextDate = nextDate.AddDate(0, 0, interval)
			if AfterNow(nextDate, now) {
				break
			}
		}
	}

	if period == 'y' {
		for {
			nextDate = nextDate.AddDate(1, 0, 0)
			if AfterNow(nextDate, now) {
				break
			}
		}
	}

	return nextDate.Format(DateFormat), nil
}

func validateStringRule(s string) bool {
	if len(s) == 0 {
		return false
	}

	firstChar := s[0]

	switch firstChar {
	case 'd':
		// Проверяем формат <d><пробел><число>
		parts := strings.SplitN(s, " ", 3)
		if len(parts) != 2 {
			return false // должно быть ровно 2 части: "d" и число
		}

		if parts[0] != "d" {
			return false // первая часть должна быть "d"
		}

		num, err := strconv.Atoi(parts[1])
		if err != nil {
			return false // не является числом
		}

		return num >= 1 && num <= 400

	case 'y':
		return len(s) == 1 // только "y" без других символов

	default:
		return false // другие первые символы недопустимы
	}
}

func AfterNow(date time.Time, now time.Time) bool {
	return date.After(now)
}
