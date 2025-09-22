package personaldata

import (
	"errors"
	"fmt"
)

type info struct {
	Name   string
	Weight float64 // kg
	Height float64 // cm
}

// New возвращает указатель на PersonalData и ошибку при неверных параметрах.
func New(name string, weight, height float64) (*PersonalData, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if weight <= 0 {
		return nil, errors.New("weight must be > 0")
	}
	if height <= 0 {
		return nil, errors.New("height must be > 0")
	}
	return &PersonalData{
		Name:   name,
		Weight: weight,
		Height: height,
	}, nil
}

// Print возвращает краткую строку с данными о пользователе.
func (p PersonalData) Print() string {
	return fmt.Sprintf("%s, %.1f kg, %.1f cm", p.Name, p.Weight, p.Height)
}
