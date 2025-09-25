package spentenergy

import (
	"testing"
	"time"
)

func TestDistanceMeanSpeed(t *testing.T) {
	d := Distance(1000, 175)
	if d <= 0 {
		t.Fatalf("expected distance > 0, got %v", d)
	}
	speed := MeanSpeed(d, 30*time.Minute)
	if speed <= 0 {
		t.Fatalf("expected speed > 0, got %v", speed)
	}
}

func TestCalories(t *testing.T) {
	_, err := WalkingSpentCalories(1000, 70, 175, 30*time.Minute)
	if err != nil {
		t.Fatalf("walk calories error %v", err)
	}
	_, err = RunningSpentCalories(1000, 70, 175, 30*time.Minute)
	if err != nil {
		t.Fatalf("run calories error %v", err)
	}
}
