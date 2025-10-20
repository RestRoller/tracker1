package spentcalories_test

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// в тестах daysteps использует фиксированную длину шага 0.65 м
	stepLength = 0.65
	// метров в километре
	mInKm = 1000.0
)

// parsePackage парсит "678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format: %q", data)
	}

	stepsPart := parts[0]
	durationPart := parts[1]

	if strings.Contains(stepsPart, " ") {
		return 0, 0, fmt.Errorf("invalid steps: contains spaces")
	}

	steps, err := strconv.Atoi(stepsPart)
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be > 0")
	}

	duration, err := time.ParseDuration(durationPart)
	if err != nil || duration <= 0 {
		return 0, 0, fmt.Errorf("invalid duration")
	}

	return steps, duration, nil
}

// DayActionInfo возвращает строку-отчет с переводом строки в конце
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	// дистанция, считаем по фиксированной длине шага
	distMeters := float64(steps) * stepLength
	distKm := distMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distKm,
		calories,
	)

	return result
}
