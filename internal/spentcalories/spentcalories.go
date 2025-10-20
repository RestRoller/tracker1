package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	stepLengthCoefficient      = 0.414
	mInKm                      = 1000
	minInH                     = 60
	walkingCaloriesCoefficient = 0.789
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("неверное количество шагов")
	}

	activityType := parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return 0, "", 0, fmt.Errorf("неверная продолжительность")
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
		return "", err
	}

	var calories float64
	var dist float64
	var speed float64

	switch activityType {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	info := fmt.Sprintf("Тип тренировки: %s\n", activityType)
	info += fmt.Sprintf("Длительность: %.2f ч.\n", duration.Hours())
	info += fmt.Sprintf("Дистанция: %.2f км.\n", dist)
	info += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	info += fmt.Sprintf("Сожгли калорий: %.2f", calories)

	return info, nil
}
