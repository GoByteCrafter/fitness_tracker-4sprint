package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию

	// парсинг строки data в слайс
	dataParse := strings.Split(data, ",")
	if len(dataParse) != 3 {
		return 0, "", 0, fmt.Errorf("ошибка входных данных parseTraining: %s", data)
	}

	// преобразование шагов в int
	steps, err := strconv.Atoi(dataParse[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка parseTraining --> strconv.Atoi: %v", err)
	} else if steps <= 0 {
		return 0, "", 0, errors.New("ошибка parseTraining --> количество шагов не может быть отрицательным или нулевым значением")
	}

	duration, err := time.ParseDuration(dataParse[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка parseTraining --> time.ParseDuration: %v", err)
	} else if duration <= 0 {
		return 0, "", 0, errors.New("ошибка parseTraining --> продолжительность не может быть отрицательным или нулевым значением")
	}
	// используем strings.ToLower на случай исключения ошибки при передачи в data вида активности в нижнем регистре
	return steps, strings.ToLower(dataParse[1]), duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	return (stepLength * float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	return distance(steps, height) / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var result string
	langTag := cases.Title(language.Und) // для возврата первого символа Activity в result в верхний регистр
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch activity {
	case "бег":
		dist := distance(steps, height)
		spd := meanSpeed(steps, height, duration)
		cal, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", langTag.String(activity), duration.Hours(), dist, spd, cal)
	case "ходьба":
		dist := distance(steps, height)
		spd := meanSpeed(steps, height, duration)
		cal, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
		result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", langTag.String(activity), duration.Hours(), dist, spd, cal)
	default:
		return result, fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("один или несколько из параметров steps, weight, height, duration отрицательные или равны 0")
	}

	return (weight * meanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("один или несколько из параметров steps, weight, height, duration отрицательные или равны 0")
	}
	rsc, _ := RunningSpentCalories(steps, weight, height, duration)

	return rsc * walkingCaloriesCoefficient, nil
}
