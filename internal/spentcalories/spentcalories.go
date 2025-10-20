package spentcalories

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

// Distance вычисляет дистанцию в километрах по количеству шагов
func Distance(steps int) float64 {
    return float64(steps) * 0.00075
}

// MeanSpeed вычисляет среднюю скорость в км/ч
func MeanSpeed(distance float64, duration time.Duration) float64 {
    hours := duration.Hours()
    if hours == 0 {
        return 0
    }
    return distance / hours
}

// ParseTraining парсит строку с данными тренировки
func ParseTraining(input string) (steps int, duration time.Duration, err error) {
    parts := strings.Fields(input)
    if len(parts) != 3 {
        return 0, 0, fmt.Errorf("неверный формат данных")
    }
    
    // Парсим шаги
    steps, err = strconv.Atoi(parts[0])
    if err != nil || steps <= 0 {
        return 0, 0, fmt.Errorf("неверное количество шагов")
    }
    
    // Парсим продолжительность
    duration, err = time.ParseDuration(parts[1] + parts[2])
    if err != nil || duration <= 0 {
        return 0, 0, fmt.Errorf("неверная продолжительность")
    }
    
    return steps, duration, nil
}

// WalkingSpentCalories вычисляет количество сожженных калорий при ходьбе
func WalkingSpentCalories(steps int, duration time.Duration, weight, height float64) float64 {
    if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
        return 0
    }
    
    distance := Distance(steps)
    speed := MeanSpeed(distance, duration)
    
    // Формула для ходьбы
    return (0.035 * weight + (speed*speed/height) * 0.029 * weight) * duration.Hours()
}

// RunningSpentCalories вычисляет количество сожженных калорий при беге
func RunningSpentCalories(steps int, duration time.Duration, weight float64) float64 {
    if steps <= 0 || duration <= 0 || weight <= 0 {
        return 0
    }
    
    distance := Distance(steps)
    speed := MeanSpeed(distance, duration)
    
    // Формула для бега
    return (0.035 * weight + (speed/1.5) * 0.035 * weight) * duration.Hours()
}

// TrainingInfo возвращает информацию о тренировке в виде строки
func TrainingInfo(trainingType string, input string, weight, height float64) (string, error) {
    steps, duration, err := ParseTraining(input)
    if err != nil {
        return "", err
    }
    
    distance := Distance(steps)
    speed := MeanSpeed(distance, duration)
    
    var calories float64
    switch trainingType {
    case "Ходьба":
        calories = WalkingSpentCalories(steps, duration, weight, height)
    case "Бег":
        calories = RunningSpentCalories(steps, duration, weight)
    default:
        return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
    }
    
    info := fmt.Sprintf("Тип тренировки: %s\n", trainingType)
    info += fmt.Sprintf("Длительность: %.2f ч.\n", duration.Hours())
    info += fmt.Sprintf("Дистанция: %.2f км.\n", distance)
    info += fmt.Sprintf("Скорость: %.2f км/ч\n", speed)
    info += fmt.Sprintf("Сожгли калорий: %.2f", calories)
    
    return info, nil
}
