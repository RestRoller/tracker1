package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/RestRoller/tracker1/internal/spentcalories"
)

const (
	stepLength = 0.65
	mInKm      = 1000.0
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format: %q", data)
	}

	stepsStr := parts[0]
	durationStr := parts[1]

	// если есть пробелы — это ошибка
	if strings.Contains(stepsStr, " ") {
		return 0, 0, fmt.Errorf("invalid steps: contains spaces")
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be > 0")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil || duration <= 0 {
		return 0, 0, fmt.Errorf("invalid duration")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

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
