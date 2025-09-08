package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// daysteps
	stepLength = 0.7
	mInKm      = 1000

	// spentcalories
	stepLengthCoefficient      = 0.415
	minInH                     = 60
	walkingCaloriesCoefficient = 0.035
)

//
// ====== DAYSTEPS ======
//

// parsePackage парсит строку "678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("некорректный формат строки: %s", data)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования времени: %v", err)
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

	distance := float64(steps) * stepLength / mInKm

	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	result := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distance, calories,
	)
	return result
}

//
// ====== SPENTCALORIES ======
//

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("некорректный формат строки: %s", data)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов: %v", err)
	}

	activity := strings.TrimSpace(parts[1])

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования времени: %v", err)
	}

	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	return (float64(steps) * stepLen) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные данные")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные данные")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := ((weight * speed * minutes) / minInH) * walkingCaloriesCoefficient
	return calories, nil
}

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
	case "бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", activity)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		activity, duration.Hours(), dist, speed, calories,
	)
	return result, nil
}

//
// ====== MAIN ======
//

func main() {

	fmt.Println("== Пример дневной активности ==")
	info := DayActionInfo("1000,30m", 70, 1.75)
	fmt.Println(info)

	fmt.Println("\n== Пример тренировки ==")
	training, err := TrainingInfo("3000,Бег,45m", 70, 1.75)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(training)
	}
}
