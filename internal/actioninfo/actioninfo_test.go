package actioninfo

import (
	"strings"
	"testing"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/internal/daysteps"
	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

// Test корректного парсинга формата "steps,duration"
func TestDaySteps_Parse_OK(t *testing.T) {
	ds := &daysteps.DaySteps{
		Personal: personaldata.PersonalData{
			Name:   "Test",
			Weight: 70,
			Height: 175,
		},
	}
	err := ds.Parse("1000,30m")
	if err != nil {
		t.Fatalf("ожидали, что парсинг успешен, получили ошибку: %v", err)
	}
	if ds.Steps != 1000 {
		t.Fatalf("ожидали Steps=1000, получили %d", ds.Steps)
	}
	if ds.Duration != 30*time.Minute {
		t.Fatalf("ожидали Duration=30m, получили %v", ds.Duration)
	}
}

// Test некорректного формата
func TestDaySteps_Parse_Invalid(t *testing.T) {
	ds := &daysteps.DaySteps{}
	err := ds.Parse("bad-format")
	if err == nil {
		t.Fatal("ожидали ошибку при парсинге, но её нет")
	}
}

// Test формата "date;steps" (используется в actioninfo)
func TestDaySteps_Parse_DateSteps_OK(t *testing.T) {
	ds := &daysteps.DaySteps{}
	err := ds.Parse("2025-09-13;5000")
	if err != nil {
		t.Fatalf("ожидали успешный парсинг date;steps, получили ошибку: %v", err)
	}
	if ds.Steps != 5000 {
		t.Fatalf("ожидали Steps=5000, получили %d", ds.Steps)
	}
}

// Test ActionInfo возвращает строку, содержащую количество шагов
func TestDaySteps_ActionInfo(t *testing.T) {
	ds := &daysteps.DaySteps{
		Personal: personaldata.PersonalData{
			Name:   "Test",
			Weight: 70,
			Height: 175,
		},
	}
	if err := ds.Parse("1000,30m"); err != nil {
		t.Fatalf("парсинг не удался: %v", err)
	}
	out, err := ds.ActionInfo()
	if err != nil {
		t.Fatalf("ожидали ActionInfo без ошибки, получили: %v", err)
	}
	if !strings.Contains(out, "1000") {
		t.Fatalf("ожидали в выводе '1000', получили: %s", out)
	}
}
