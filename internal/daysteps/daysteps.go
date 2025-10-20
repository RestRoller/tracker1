package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/RestRoller/tracker1/internal/spentcalories"
)

const (
	// средняя длина шага в метрах
	stepLength = 0.65
	// метров в километре
	mInKm = 1000.0
)

// parsePackage парсит строку вида "678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format: %q", data)
	}

	stepsStr := strings.TrimSpace(parts[0])
	durationStr := strings.TrimSpace(parts[1])

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps: %w", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be > 0")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration: %w", err)
	}

	return steps, duration, nil
}

// DayActionInfo возвращает форматированный отчет о прогулке
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	// дистанция в метрах = steps * stepLength
	distMeters := float64(steps) * stepLength
	distKm := distMeters / mInKm

	// калории рассчитываются функцией из пакета spentcalories (ходьба)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		// если ошибка — вывести и вернуть пустую строку
		fmt.Println(err)
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps,
		distKm,
		calories,
	)

	return result
}
