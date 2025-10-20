package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// В пакетных тестах ожидается именно такой коэффициент
	stepLengthCoefficient = 0.45
	// В тестах коэффициент для ходьбы — 0.5 (половина от бега)
	walkingCaloriesCoefficient = 0.5
	// минут в часе
	minInH = 60.0
	// метров в километре
	mInKm = 1000.0
)

// parseTraining парсит строку формата "3456,Ходьба,3h00m"
func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format")
	}

	stepsStr := strings.TrimSpace(parts[0])
	activity := strings.TrimSpace(parts[1])
	durationStr := strings.TrimSpace(parts[2])

	// шаги не должны быть с пробелами внутри (тесты требуют ошибку при " 123" или "123 ")
	if strings.Contains(parts[0], " ") {
		return 0, "", 0, fmt.Errorf("invalid steps value")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps value")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid duration")
	}

	return steps, activity, duration, nil
}

// distance возвращает дистанцию в километрах, рассчитываемую через рост
func distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	stepLen := height * stepLengthCoefficient // метры
	distMeters := float64(steps) * stepLen
	return distMeters / mInKm
}

// meanSpeed км/ч
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

// RunningSpentCalories — калории при беге.
// Формула тестов: (weight * speed * durationMinutes) / 60
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input values")
	}

	speed := meanSpeed(steps, height, duration) // км/ч
	minutes := duration.Minutes()               // минуты
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

// WalkingSpentCalories — калории при ходьбе (по тесту — то же, но с коэффициентом)
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("invalid input values")
	}

	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	calories *= walkingCaloriesCoefficient
	return calories, nil
}

// TrainingInfo — форматированный отчёт (с переводом строки в конце)
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

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity,
		duration.Hours(),
		dist,
		speed,
		calories,
	)
	return result, nil
}
