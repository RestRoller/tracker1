package spentcalories

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

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

func distance(steps int, length float64) float64 {
    return float64(steps) * length
}

func meanSpeed(steps int, length float64, duration time.Duration) float64 {
    dist := distance(steps, length)
    hours := duration.Hours()
    if hours == 0 {
        return 0
    }
    return dist / hours
}

func WalkingSpentCalories(steps int, weight float64, height float64, duration time.Duration) (float64, float64) {
    if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
        return 0, 0
    }
    
    length := 0.00075
    dist := distance(steps, length)
    speed := meanSpeed(steps, length, duration)
    
    calories := (0.035 * weight + (speed*speed/height) * 0.029 * weight) * duration.Hours()
    return calories, dist
}

func RunningSpentCalories(steps int, weight float64, height float64, duration time.Duration) (float64, float64) {
    if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
        return 0, 0
    }
    
    length := 0.00075
    dist := distance(steps, length)
    speed := meanSpeed(steps, length, duration)
    
    calories := (0.035 * weight + (speed/1.5) * 0.035 * weight) * duration.Hours()
    return calories, dist
}

func TrainingInfo(trainingType string, input string, weight, height float64) (string, error) {
    steps, duration, _, err := parseTraining(input)
    if err != nil {
        return "", err
    }
    
    var calories, distance float64
    
    switch trainingType {
    case "Ходьба":
        calories, distance = WalkingSpentCalories(steps, weight, height, duration)
    case "Бег":
        calories, distance = RunningSpentCalories(steps, weight, height, duration)
    default:
        return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
    }
    
    speed := meanSpeed(steps, 0.00075, duration)
    
    info := fmt.Sprintf("Тип тренировки: %s\n", trainingType)
    info += fmt.Sprintf("Длительность: %.2f ч.\n", duration.Hours())
    info += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
    info += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
    info += fmt.Sprintf("Сожгли калорий: %.2f", calories)
    
    return info, nil
}
