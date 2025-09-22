package daysteps

import (
	"testing"

	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

func TestDaySteps_Info(t *testing.T) {
	pd := personaldata.PersonalData{Name: "Иван", Weight: 70, Height: 175}

	ds, err := New("4000 шагов", pd)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	got, err := ds.Info()
	if err != nil {
		t.Errorf("Info() error = %v", err)
	}
	if got == "" {
		t.Errorf("Info() вернул пустую строку")
	}
}
