package main

import (
	"fmt"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/internal/actioninfo"
	"github.com/Sevacoming/go1fl-4-sprint-final/internal/daysteps"
	"github.com/Sevacoming/go1fl-4-sprint-final/personaldata"
)

func main() {
	pd, err := personaldata.New("Иван", 70, 1.75)
	if err != nil {
		panic(err)
	}

	ds := &daysteps.DaySteps{
		Duration: 30 * time.Minute,
		Personal: *pd,
	}

	data := []string{"2025-09-13;1000"}

	fmt.Println("== Пример дневной активности ==")
	actioninfo.Info(data, ds)
}
