package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	steps int
}

func (ds *DaySteps) Add(input string) error {
	steps, _, _, err := parseTraining(input)
	if err != nil {
		return err
	}
	ds.steps += steps
	return nil
}

func (ds *DaySteps) Steps() int {
	return ds.steps
}

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

//DayActionInfo
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, _, err := parseTraining(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distance := float64(steps) * 0.00075
	calories := (0.035 * weight + (4.35*4.35/height)*0.029*weight) * duration.Hours()

	result := fmt.Sprintf("Количество шагов: %d.\n", steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.", calories)

	return result
}
