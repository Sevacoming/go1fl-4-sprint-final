package personaldata

import "fmt"

type PersonalData struct {
	Name   string
	Weight float64
	Height float64
}

func New(name string, weight, height float64) (*PersonalData, error) {
	if name == "" {
		return nil, fmt.Errorf("имя не может быть пустым")
	}
	if weight <= 0 {
		return nil, fmt.Errorf("вес должен быть > 0")
	}
	if height <= 0 {
		return nil, fmt.Errorf("рост должен быть > 0")
	}

	return &PersonalData{
		Name:   name,
		Weight: weight,
		Height: height,
	}, nil
}
