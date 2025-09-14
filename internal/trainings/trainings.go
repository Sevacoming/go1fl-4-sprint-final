package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/internal/personaldata"
	"github.com/Sevacoming/go1fl-4-sprint-final/internal/spentcalories"
	"github.com/Sevacoming/go1fl-4-sprint-final/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("ошибка парсинга: ожидали 3 элемента, получили %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать шаги: %w", err)
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать длительность: %w", err)
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps <= 0 {
		return "", fmt.Errorf("ошибка: количество шагов должно быть > 0")
	}
	if t.Duration <= 0 {
		return "", fmt.Errorf("ошибка: длительность должна быть > 0")
	}

	distKm := spentenergy.DistanceKm(t.Steps)
	speed := spentenergy.Speed(distKm, t.Duration)

	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "ходьба":
		calories, err = spentcalories.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "бег":
		calories, err = spentcalories.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}
	if err != nil {
		return "", fmt.Errorf("ошибка вычисления калорий: %w", err)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType,
		t.Duration.Hours(),
		distKm,
		speed,
		calories,
	)

	return result, nil
}
