package spentenergy

import (
	"errors"
	"time"
)

// Distance возвращает примерную дистанцию в километрах на основе числа шагов и роста (см).
// Простая модель: шаг = height * 0.45% (height в см -> m), затем km.
func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0.0
	}
	// примерная длина шага = height * 0.415 (см -> m)
	stepM := height / 100.0 * 0.415
	distanceMeters := float64(steps) * stepM
	return distanceMeters / 1000.0
}

// MeanSpeed возвращает среднюю скорость (км/ч) через distance (км) и duration (time.Duration).
func MeanSpeed(distanceKm float64, duration time.Duration) float64 {
	if duration <= 0 || distanceKm <= 0 {
		return 0.0
	}
	hours := duration.Hours()
	return distanceKm / hours
}

// WalkingSpentCalories - простая модель вычисления потраченных калорий при ходьбе.
// Возвращает (kcal, error)
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid parameters")
	}
	// модель: MET ~ 3.5 для ходьбы, kcal = MET * weight(kg) * hours
	met := 3.5
	kcal := met * weight * duration.Hours()
	return kcal, nil
}

// RunningSpentCalories - модель для бега.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid parameters")
	}
	// модель: MET ~ 9.0 для бега
	met := 9.0
	kcal := met * weight * duration.Hours()
	return kcal, nil
}
