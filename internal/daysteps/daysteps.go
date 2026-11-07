package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию

	// парсинг строки data в слайс {шаги продолжительность}
	strParse := strings.Split(data, ",")
	if len(strParse) != 2 {
		return 0, 0, fmt.Errorf("ошибка входных данных parsePackage: %s", data)
	}
	// преобразование шагов в int
	steps, err := strconv.Atoi(strParse[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка parsePackage --> strconv.Atoi: %v", err)
	} else if steps <= 0 {
		return 0, 0, errors.New("ошибка parsePackage --> количество шагов не может быть отрицательным или нулевым значением")
	}
	// преобразование продолжительности в time.Duration в формате XXhXXm
	duration, err := time.ParseDuration(strParse[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка parsePackage --> time.ParseDuration: %v", err)
	} else if duration <= 0 {
		return 0, 0, errors.New("ошибка parsePackage --> продолжительность не может быть отрицательной")
	}
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	var result string

	// получение количество шагов и продолжитиельности прогулки
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Printf("ошибка DayActionInfo --> parsePackage: %v", err)
		return result
	}
	if steps <= 0 {
		fmt.Printf("ошибка DayActionInfo --> parsePackage: количество шагов не может быть отрицательным или нулевым значением")
		return result
	}
	if weight <= 0 || height <= 0 {
		fmt.Printf("ошибка DayActionInfo --> вес и(или) рост не могут быть отрицательными или нулевыми значениями")
		return result
	}
	// подсчет дистанции
	distance := float64(steps) * stepLength / mInKm

	// подстчет калорий
	cal, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Printf("ошибка DayActionInfo --> WalkingSpentCalories: %v", err)
		return result
	}

	result = fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, cal)
	return result
}
