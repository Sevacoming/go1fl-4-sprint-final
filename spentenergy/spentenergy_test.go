package spentenergy

import (
	"testing"
	"time"
)

func TestDistance(t *testing.T) {
	dist := Distance(2000, 1.75)
	if dist <= 0 {
		t.Errorf("ожидали положительную дистанцию, получили %.2f", dist)
	}
}

func TestMeanSpeed(t *testing.T) {
	speed := MeanSpeed(2000, 1.75, 30*time.Minute)
	if speed <= 0 {
		t.Errorf("ожидали положительную скорость, получили %.2f", speed)
	}
}

func TestRunningSpentCalories_Valid(t *testing.T) {
	calories, err := RunningSpentCalories(3000, 70, 1.75, 45*time.Minute)
	if err != nil {
		t.Errorf("не ожидали ошибку, но получили: %v", err)
	}
	if calories <= 0 {
		t.Errorf("ожидали положительное число калорий, получили %.2f", calories)
	}
}

func TestRunningSpentCalories_Invalid(t *testing.T) {
	_, err := RunningSpentCalories(-500, 70, 1.75, 20*time.Minute)
	if err == nil {
		t.Error("ожидали ошибку при некорректных параметрах, но получили nil")
	}
}

func TestWalkingSpentCalories_Valid(t *testing.T) {
	calories, err := WalkingSpentCalories(1500, 65, 1.70, 30*time.Minute)
	if err != nil {
		t.Errorf("не ожидали ошибку, но получили: %v", err)
	}
	if calories <= 0 {
		t.Errorf("ожидали положительное число калорий, получили %.2f", calories)
	}
}

func TestWalkingSpentCalories_Invalid(t *testing.T) {
	_, err := WalkingSpentCalories(0, 65, 1.70, 30*time.Minute)
	if err == nil {
		t.Error("ожидали ошибку при steps=0, но получили nil")
	}
}
