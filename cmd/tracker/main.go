package main

import (
	"fmt"
	"time"

	"github.com/Sevacomimg/go11f-4-sprint-final/internal/actioninfo"
	"github.com/Sevacomimg/go11f-4-sprint-final/internal/daysteps"
	"github.com/Sevacomimg/go11f-4-sprint-final/internal/personaldata"
)

func main() {
	pd := personaldata.New("Иван", 70, 1.75)

	ds := &daysteps.DaySteps{
		Duration: 30 * time.Minute,
		Personal: *pd,
	}

	data := []string{"2025-09-13;1000"}

	fmt.Println("== Пример дневной активности ==")
	actioninfo.Info(data, ds)
}
