package daysteps

import (
	"strings"
	"testing"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

func TestDayStepsActionInfo_OK(t *testing.T) {
	ds := &DaySteps{
		Steps:    1000,
		Duration: 30 * time.Minute,
		Personal: personaldata.PersonalData{Name: "Иван", Weight: 70, Height: 175},
	}

	got, err := ds.ActionInfo()
	if err != nil {
		t.Errorf("ActionInfo вернул ошибку: %v", err)
	}

	if got == "" {
		t.Errorf("ActionInfo вернул пустую строку")
	}

	if !strings.Contains(got, "Иван") {
		t.Errorf("ожидали имя в результате, получили: %s", got)
	}
}

func TestDayStepsActionInfo_Error(t *testing.T) {
	ds := &DaySteps{
		Steps:    0, // некорректное значение
		Duration: 30 * time.Minute,
		Personal: personaldata.PersonalData{Name: "Иван", Weight: 70, Height: 175},
	}

	_, err := ds.ActionInfo()
	if err == nil {
		t.Errorf("ожидали ошибку, но её не было")
	}

	if err != nil && !strings.Contains(err.Error(), "шагов нет") {
		t.Errorf("ожидали ошибку 'шагов нет', получили: %v", err)
	}
}
