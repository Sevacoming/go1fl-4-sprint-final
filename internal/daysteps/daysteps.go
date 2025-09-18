package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/internal/spentcalories"
	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Personal personaldata.PersonalData
}

func (ds *DaySteps) Parse(input string) error {
	parts := strings.Split(input, ";")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат входных данных: %s", input)
	}

	steps, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать шаги: %w", err)
	}

	ds.Steps = steps
	ds.Duration = 30 * time.Minute

	return nil
}

func (ds *DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("шагов нет")
	}

	durationMinutes := ds.Duration.Minutes()
	calories, err := spentcalories.WalkingSpentCalories(ds.Steps,
		ds.Personal.Weight, ds.Personal.Height,
		ds.Duration,
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"%s прошёл %d шагов за %.0f минут, потратил %.2f ккал",
		ds.Personal.Name,
		ds.Steps,
		durationMinutes,
		calories,
	), nil
}
