package main

import (
	"fmt"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/internal/daysteps"
	"github.com/Sevacoming/go1fl-4-sprint-final/internal/trainings"
	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

func main() {
	// Создаём данные о пользователе
	pd := personaldata.PersonalData{
		Name:   "Иван",
		Weight: 70,
		Height: 175,
	}

	// Проверяем DaySteps
	ds, err := daysteps.New("4000 шагов", pd)
	if err != nil {
		fmt.Println("Ошибка при создании DaySteps:", err)
		return
	}
	dayInfo, _ := ds.Info()
	fmt.Println("Инфо по шагам:", dayInfo)

	// Проверяем Training (ходьба)
	trWalk := trainings.Training{
		Steps:        4000,
		TrainingType: "ходьба",
		Duration:     time.Hour,
		Personal:     pd,
	}
	infoWalk, err := trWalk.ActionInfo()
	if err != nil {
		fmt.Println("Ошибка тренировки (ходьба):", err)
	} else {
		fmt.Println("Инфо по тренировке (ходьба):", infoWalk)
	}

	// Проверяем Training (бег)
	trRun := trainings.Training{
		Steps:        5000,
		TrainingType: "бег",
		Duration:     time.Hour,
		Personal:     pd,
	}
	infoRun, err := trRun.ActionInfo()
	if err != nil {
		fmt.Println("Ошибка тренировки (бег):", err)
	} else {
		fmt.Println("Инфо по тренировке (бег):", infoRun)
	}
}
