package trainings

import (
	"strings"
	"testing"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

func TestTraining_Parse(t *testing.T) {
	tr := &Training{}
	if err := tr.Parse("4000,Ходьба,1h00m"); err != nil {
		t.Fatalf("ожидали успешный парсинг, получили ошибку: %v", err)
	}
	if tr.Steps != 4000 || tr.TrainingType != "Ходьба" || tr.Duration != time.Hour {
		t.Fatalf("неверный результат парсинга: %+v", tr)
	}
}

func TestTraining_Parse_Invalid(t *testing.T) {
	tr := &Training{}
	if err := tr.Parse("bad-data"); err == nil {
		t.Fatal("ожидали ошибку при неверном формате")
	}
}

func TestTraining_ActionInfo_Walking(t *testing.T) {
	tr := &Training{
		Steps:        4000,
		TrainingType: "Ходьба",
		Duration:     time.Hour,
		Personal:     personaldata.PersonalData{Name: "Test", Weight: 70, Height: 175},
	}
	res, err := tr.ActionInfo()
	if err != nil {
		t.Fatalf("ожидали success, получили ошибку: %v", err)
	}
	if !strings.Contains(res, "Тип тренировки: Ходьба") {
		t.Error("ActionInfo: нет строки с типом")
	}
	if !strings.Contains(res, "длительность: 1.00 ч.") {
		t.Error("ActionInfo: нет строки с длительностью")
	}
	if !strings.Contains(res, "Сожгли калорий:") {
		t.Error("ActionInfo: нет строки с калориями")
	}
}

func TestTraining_ActionInfo_UnknownType(t *testing.T) {
	tr := &Training{Steps: 1000, TrainingType: "Йога", Duration: time.Hour, Personal: personaldata.PersonalData{Name: "T", Weight: 70, Height: 175}}
	if _, err := tr.ActionInfo(); err == nil {
		t.Fatal("ожидали ошибку для неизвестного типа тренировки")
	}
}
