package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	stepLength = 0.65
	mInKm      = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверный формат данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("неверное количество шагов")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil || duration <= 0 {
		return 0, 0, fmt.Errorf("неверная продолжительность")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories := calculateWalkingCalories(steps, weight, height, duration)

	result := fmt.Sprintf("Количество шагов: %d.\n", steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distanceKm)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.", calories)

	return result
}

func calculateWalkingCalories(steps int, weight, height float64, duration time.Duration) float64 {
	speed := (float64(steps) * stepLength / mInKm) / duration.Hours()
	durationMinutes := duration.Minutes()
	calories := (weight * speed * durationMinutes) / 60
	return calories * 0.789 // walkingCaloriesCoefficient
}
