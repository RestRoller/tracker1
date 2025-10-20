package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	stepLengthCoefficient      = 0.414
	mInKm                      = 1000
	minInH                     = 60
	walkingCaloriesCoefficient = 0.029
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка парсинга шагов: %v", err)
	}

	activityType := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка парсинга времени: %v", err)
	}

	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength
	return distanceMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	hours := duration.Hours()
	return dist / hours
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные параметры")
	}

	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	calories := (weight * speed * durationMinutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные параметры")
	}

	speed := meanSpeed(steps, height, duration)
	durationMinutes := duration.Minutes()

	calories := (weight * speed * durationMinutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var result string
	var calories float64
	var dist float64
	var speed float64

	switch activityType {
	case "Бег", "бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		result = fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			duration.Hours(), dist, speed, calories)

	case "Ходьба", "ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		result = fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			duration.Hours(), dist, speed, calories)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return result, nil
}
