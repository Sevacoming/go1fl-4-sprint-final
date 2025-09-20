package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Sevacoming/go1f1-4-sprint-final/internal/personaldata"
	"github.com/Sevacoming/go1f1-4-sprint-final/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.PersonalData
}

func New(input string, pd personaldata.PersonalData) (*DaySteps, error) {
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("неверный формат входных данных: %s", input)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	duration := time.Hour // временно фиксировано

	return &DaySteps{
		Steps:    steps,
		Duration: duration,
		Personal: pd,
	}, nil
}

func (ds DaySteps) Info() (string, error) {
	kcal, err := spentenergy.WalkingSpentCalories(
		ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration,
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Шагов: %d, длительность %v, потрачено %.2f ккал",
		ds.Steps, ds.Duration, kcal,
	), nil
}
