package spentcalories

import (
	"testing"
	"time"
)

// ✅ Тест WalkingSpentCalories
func TestWalkingSpentCalories(t *testing.T) {
	steps := 1000
	weight := 70.0
	height := 1.75
	duration := 30 * time.Minute

	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		t.Errorf("не ожидали ошибку, получили: %v", err)
	}

	if calories <= 0 {
		t.Errorf("ожидали положительное значение калорий, получили %v", calories)
	}
}

// ✅ Тест RunningSpentCalories
func TestRunningSpentCalories(t *testing.T) {
	steps := 2000
	weight := 70.0
	height := 1.75
	duration := 45 * time.Minute

	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		t.Errorf("не ожидали ошибку, получили: %v", err)
	}

	if calories <= 0 {
		t.Errorf("ожидали положительное значение калорий, получили %v", calories)
	}
}

// ✅ Тест parseTraining с корректными данными
func TestParseTraining_ValidData(t *testing.T) {
	input := "1000,ходьба,30m"
	steps, activity, duration, err := parseTraining(input)
	if err != nil {
		t.Errorf("не ожидали ошибку, получили: %v", err)
	}
	if steps != 1000 {
		t.Errorf("ожидали 1000 шагов, получили %d", steps)
	}
	if activity != "ходьба" {
		t.Errorf("ожидали 'ходьба', получили %s", activity)
	}
	if duration != 30*time.Minute {
		t.Errorf("ожидали 30 минут, получили %v", duration)
	}
}

// ✅ Тест parseTraining с некорректным форматом
func TestParseTraining_InvalidFormat(t *testing.T) {
	_, _, _, err := parseTraining("invalid,data")
	if err == nil {
		t.Error("ожидали ошибку при некорректных данных, получили nil")
	}
}

// ✅ Тест parseTraining с некорректными числами
func TestParseTraining_InvalidNumbers(t *testing.T) {
	input := "abc,бег,30m"
	_, _, _, err := parseTraining(input)
	if err == nil {
		t.Error("ожидали ошибку при некорректных числовых значениях, получили nil")
	}
}

// ✅ Тест parseTraining с некорректной длительностью
func TestParseTraining_InvalidDuration(t *testing.T) {
	input := "1000,ходьба,abc"
	_, _, _, err := parseTraining(input)
	if err == nil {
		t.Error("ожидали ошибку при некорректной длительности, получили nil")
	}
}
