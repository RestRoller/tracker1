package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	stepLengthCoefficient     = 0.45   // длина шага относительно роста
	walkingCaloriesCoefficient = 0.8
	mInKm                      = 1000.0
)

// parseTraining парсит строку формата "3456,Ходьба,3h00m"
func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps value")
	}

	activity := strings.TrimSpace(parts[1])
	durationStr := strings.TrimSpace(parts[2])

	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid duration")
	}

	return steps, activity, duration, nil
}

// distance возвращает дистанцию в км
func distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient
	distMeters := float64(steps) * stepLength
	return distMeters / mInKm
}

// meanSpeed возвращает среднюю скорость (км/ч)
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	hours := duration.Hours()
	if hours <= 0 {
		return 0
	}
	return dist / hours
}

// RunningSpentCalories — калории при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input values")
	}

	speed := meanSpeed(steps, height, duration)
	hours := duration.Hours()
	calories := (0.035*weight + (speed*speed/height)*0.029*weight) * hours
	return calories, nil
}

// WalkingSpentCalories — калории при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input values")
	}

	speed := meanSpeed(steps, height, duration)
	hours := duration.Hours()
	calories := (0.035*weight + (speed*speed/height)*0.029*weight) * hours
	calories *= walkingCaloriesCoefficient
	return calories, nil
}

// TrainingInfo — текстовый отчёт
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	var calories float64

	switch strings.ToLower(activity) {
	case "ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	case "бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		log.Println(err)
		return "", err
	}

	report := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	)
	return report, nil
}
