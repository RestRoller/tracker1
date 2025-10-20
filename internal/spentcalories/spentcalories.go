package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	stepLengthCoefficient      = 0.414
	mInKm                      = 1000
	minInH                     = 60
	walkingCaloriesCoefficient = 0.029
)

func ParseTraining(input string) (steps int, duration time.Duration, err error) {
    parts := strings.Fields(input)
    if len(parts) != 3 {
        return 0, 0, fmt.Errorf("неверный формат данных")
    }
    
    steps, err = strconv.Atoi(parts[0])
    if err != nil || steps <= 0 {
        return 0, 0, fmt.Errorf("неверное количество шагов")
    }
	
    duration, err = time.ParseDuration(parts[1] + parts[2])
    if err != nil || duration <= 0 {
        return 0, 0, fmt.Errorf("неверная продолжительность")
    }
    
    return steps, duration, nil
}

func Distance(steps int) float64 {
    return float64(steps) * 0.00075
}

func MeanSpeed(distance float64, duration time.Duration) float64 {
    hours := duration.Hours()
    if hours == 0 {
        return 0
    }
    return distance / hours
}

func RunningSpentCalories(steps int, duration time.Duration, weight float64) float64 {
    if steps <= 0 || duration <= 0 || weight <= 0 {
        return 0
    }
    
    speed := MeanSpeed(Distance(steps), duration)
    return (0.035 * weight + (speed/1.5) * 0.035 * weight) * duration.Hours()
}

func WalkingSpentCalories(steps int, duration time.Duration, weight, height float64) float64 {
    if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
        return 0
    }
    
    speed := MeanSpeed(Distance(steps), duration)
    return (0.035 * weight + (speed*speed/height) * 0.029 * weight) * duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activityType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var result string
	var calories float64
	var dist float64
	var speed float64

	switch activityType {
	case "Бег", "бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		result = fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			duration.Hours(), dist, speed, calories)

	case "Ходьба", "ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		result = fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			duration.Hours(), dist, speed, calories)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return result, nil
}
