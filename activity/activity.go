package activity

import "fmt"

type Activity struct {
	Type     string
	Duration float64
	Distance float64
}

func New(activityType string, duration, distance float64) (*Activity, error) {
	if activityType == "" {
		return nil, fmt.Errorf("тип активности не может быть пустым")
	}
	if duration <= 0 {
		return nil, fmt.Errorf("длительность должна быть > 0")
	}
	if distance < 0 {
		return nil, fmt.Errorf("дистанция не может быть отрицательной")
	}
	return &Activity{
		Type:     activityType,
		Duration: duration,
		Distance: distance,
	}, nil
}
