package trainings

import (
	"strings"
	"testing"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

func TestTraining_Parse(t *testing.T) {
	t1 := &Training{}

	err := t1.Parse("5000,Ходьба,1h30m")
	if err != nil {
		t.Fatalf("Parse вернул ошибку: %v", err)
	}

	if t1.Steps != 5000 {
		t.Errorf("ожидали шаги 5000, получили %d", t1.Steps)
	}
	if t1.TrainingType != "Ходьба" {
		t.Errorf("ожидали тип 'Ходьба', получили %s", t1.TrainingType)
	}
	if t1.Duration != time.Hour+30*time.Minute {
		t.Errorf("ожидали 1h30m, получили %v", t1.Duration)
	}
}

func TestTraining_ActionInfo_Walking(t *testing.T) {
	tr := Training{
		Steps:        4000,
		TrainingType: "Ходьба",
		Duration:     time.Hour,
		Personal:     personaldata.PersonalData{Name: "Test", Weight: 70, Height: 175},
	}

	res, err := tr.ActionInfo()
	if err != nil {
		t.Fatalf("ActionInfo вернул ошибку: %v", err)
	}

	if !strings.Contains(res, "Ходьба") {
		t.Errorf("ожидали в выводе 'Ходьба', получили: %s", res)
	}
	if !strings.Contains(res, "Дистанция") {
		t.Errorf("ожидали наличие дистанции, получили: %s", res)
	}
}

func TestTraining_ActionInfo_Running(t *testing.T) {
	tr := Training{
		Steps:        6000,
		TrainingType: "Бег",
		Duration:     time.Hour,
		Personal:     personaldata.PersonalData{Name: "Test", Weight: 70, Height: 175},
	}

	res, err := tr.ActionInfo()
	if err != nil {
		t.Fatalf("ActionInfo вернул ошибку: %v", err)
	}

	if !strings.Contains(res, "Бег") {
		t.Errorf("ожидали в выводе 'Бег', получили: %s", res)
	}
}

func TestTraining_ActionInfo_UnknownType(t *testing.T) {
	tr := Training{
		Steps:        5000,
		TrainingType: "Йога",
		Duration:     time.Hour,
		Personal:     personaldata.PersonalData{Name: "Test", Weight: 70, Height: 175},
	}

	_, err := tr.ActionInfo()
	if err == nil {
		t.Errorf("ожидали ошибку для неизвестного типа, но её не было")
	}
}
