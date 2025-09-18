package spentcalories

import (
	"testing"
	"time"
)

func TestWalkingSpentCalories_Valid(t *testing.T) {
	cal, err := WalkingSpentCalories(1000, 70.0, 1.75, 30*time.Minute)
	if err != nil {
		t.Fatalf("ожидали без ошибки, получили: %v", err)
	}
	if cal <= 0 {
		t.Fatalf("ожидали положительное значение калорий, получили: %v", cal)
	}
}

func TestWalkingSpentCalories_InvalidParams(t *testing.T) {

	if _, err := WalkingSpentCalories(0, 70.0, 1.75, 30*time.Minute); err == nil {
		t.Error("ожидалась ошибка при нулевых шагах")
	}

	if _, err := WalkingSpentCalories(1000, -70.0, 1.75, 30*time.Minute); err == nil {
		t.Error("ожидалась ошибка при отрицательном весе")
	}

	if _, err := WalkingSpentCalories(1000, 70.0, 0.0, 30*time.Minute); err == nil {
		t.Error("ожидалась ошибка при нулевом росте")
	}

	if _, err := WalkingSpentCalories(1000, 70.0, 1.75, 0*time.Minute); err == nil {
		t.Error("ожидалась ошибка при нулевой длительности")
	}
}

func TestRunningSpentCalories_Valid(t *testing.T) {
	cal, err := RunningSpentCalories(1500, 70.0, 1.75, 45*time.Minute)
	if err != nil {
		t.Fatalf("ожидали без ошибки, получили: %v", err)
	}
	if cal <= 0 {
		t.Fatalf("ожидали положительное значение калорий, получили: %v", cal)
	}
}

func TestRunningSpentCalories_InvalidParams(t *testing.T) {

	if _, err := RunningSpentCalories(0, 70.0, 1.75, 30*time.Minute); err == nil {
		t.Error("ожидалась ошибка при нулевых шагах")
	}

	if _, err := RunningSpentCalories(1000, -70.0, 1.75, 30*time.Minute); err == nil {
		t.Error("ожидалась ошибка при отрицательном весе")
	}

	if _, err := RunningSpentCalories(1000, 70.0, 0.0, 30*time.Minute); err == nil {
		t.Error("ожидалась ошибка при нулевом росте")
	}

	if _, err := RunningSpentCalories(1000, 70.0, 1.75, 0*time.Minute); err == nil {
		t.Error("ожидалась ошибка при нулевой длительности")
	}
}
