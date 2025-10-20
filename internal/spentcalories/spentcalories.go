package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseTraining(input string) (int, time.Duration, float64, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, 0, 0, fmt.Errorf("неверное количество шагов")
	}

	duration, err := time.ParseDuration(parts[1] + parts[2])
	if err != nil || duration <= 0 {
		return 0, 0, 0, fmt.Errorf("неверная продолжительность")
	}

	distance := float64(steps) * 0.00075

	return steps, duration, distance, nil
}

func distance(steps int, length float64) float64 {
	return float64(steps) * length
}

func meanSpeed(distance float64, duration time.Duration) float64 {
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return distance / hours
}

func WalkingSpentCalories(steps int, duration time.Duration, weight float64, height float64) float64 {
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0
	}

	dist := float64(steps) * 0.00075
	speed := meanSpeed(dist, duration)

	return (0.035 * weight + (speed*speed/height)*0.029*weight) * duration.Hours()
}

func RunningSpentCalories(steps int, duration time.Duration, weight float64) float64 {
	if steps <= 0 || duration <= 0 || weight <= 0 {
		return 0
	}

	dist := float64(steps) * 0.00075
	speed := meanSpeed(dist, duration)

	return (0.035 * weight + (speed/1.5)*0.035*weight) * duration.Hours()
}

func TrainingInfo(trainingType string, input string, weight, height float64) (string, error) {
	steps, duration, dist, err := parseTraining(input)
	if err != nil {
		return "", err
	}

	var calories float64

	switch trainingType {
	case "Ходьба":
		calories = WalkingSpentCalories(steps, duration, weight, height)
	case "Бег":
		calories = RunningSpentCalories(steps, duration, weight)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
	}

	speed := meanSpeed(dist, duration)

	info := fmt.Sprintf("Тип тренировки: %s\n", trainingType)
	info += fmt.Sprintf("Длительность: %.2f ч.\n", duration.Hours())
	info += fmt.Sprintf("Дистанция: %.2f км.\n", dist)
	info += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
	info += fmt.Sprintf("Сожгли калорий: %.2f", calories)

	return info, nil
}
//update NUMBER 32843
