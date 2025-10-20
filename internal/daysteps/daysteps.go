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
    steps, _, _, err := parsePackage(input)
    if err != nil {
        return err
    }
    ds.steps += steps
    return nil
}

func (ds *DaySteps) Steps() int {
    return ds.steps
}

// Используем parsePackage вместо parseTraining для daysteps
func parsePackage(input string) (int, time.Duration, float64, error) {
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
