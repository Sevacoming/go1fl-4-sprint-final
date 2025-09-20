package trainings

import (
	"fmt"
	"time"

	"github.com/Sevacomimg/go1f1-4-sprint-final/internal/personaldata"
	"github.com/Sevacomimg/go1f1-4-sprint-final/internal/spentenergy"
)

type Training struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.PersonalData
	Action   string
}

// ActionInfo возвращает информацию о тренировке
func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 || t.Duration <= 0 {
		return "", fmt.Errorf("некорректные параметры тренировки")
	}

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(distance, t.Duration)

	var calories float64
	var err error

	switch t.Action {
	case "run":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration,
		)
	case "walk":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration,
		)
	default:
		return "", fmt.Errorf("неизвестное действие: %s", t.Action)
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Действие: %s, дистанция: %.2f км, скорость: %.2f км/ч, калории: %.1f",
		t.Action, distance, speed, calories,
	), nil
}
