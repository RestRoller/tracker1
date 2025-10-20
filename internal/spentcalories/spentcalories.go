package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// коэффициент длины шага (height * stepLengthCoefficient = длина шага в метрах)
	stepLengthCoefficient    = 0.415
	// коэффициент для расчёта калорий при ходьбе
	walkingCaloriesCoefficient = 0.8
	// минут в часе
	minInH = 60.0
	// метров в километре
	mInKm = 1000.0
)

// parseTraining парсит строку вида "3456,Ходьба,3h00m"
func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid training data format: %q", data)
	}

	stepsStr := strings.TrimSpace(parts[0])
	activity := strings.TrimSpace(parts[1])
	durationStr := strings.TrimSpace(parts[2])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps: %w", err)
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration: %w", err)
	}

	return steps, activity, duration, nil
}

// distance возвращает дистанцию в километрах по шагам и росту (в метрах)
func distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLength := height * stepLengthCoefficient // в метрах
	distMeters := float64(steps) * stepLength
	return distMeters / mInKm
}

// meanSpeed возвращает среднюю скорость в км/ч
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distKm := distance(steps, height)
	hours := duration.Hours()
	if hours == 0 {
		return 0
	}
	return distKm / hours
}

// RunningSpentCalories рассчитывает калории при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be > 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be > 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be > 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be > 0")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

// WalkingSpentCalories рассчитывает калории при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("steps must be > 0")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("weight must be > 0")
	}
	if height <= 0 {
		return 0, fmt.Errorf("height must be > 0")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("duration must be > 0")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()

	calories := (weight * speed * minutes) / minInH
	calories = calories * walkingCaloriesCoefficient
	return calories, nil
}

// TrainingInfo возвращает форматированный отчёт по тренировке
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// нормальная проверка входных параметров
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return "", fmt.Errorf("invalid input values")
	}

	dist := distance(steps, height)                // в км
	speed := meanSpeed(steps, height, duration)    // км/ч
	calories := 0.0

	switch strings.ToLower(activity) {
	case strings.ToLower("Бег"), strings.ToLower("бег"):
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case strings.ToLower("Ходьба"), strings.ToLower("ходьба"):
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", fmt.Errorf("unknown training type: %q", activity)
	}

	// Форматирование отчёта: значения с двумя знаками после запятой
	hours := duration.Hours()
	report := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity,
		hours,
		dist,
		speed,
		calories,
	)

	return report, nil
}
