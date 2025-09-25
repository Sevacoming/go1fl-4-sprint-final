package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {

		if err := dp.Parse(data); err != nil {
			log.Println("parse error:", err)
			continue
		}

		s, err := dp.ActionInfo()
		if err != nil {
			log.Println("info error:", err)
			continue
		}

		fmt.Println(s)
	}
}
