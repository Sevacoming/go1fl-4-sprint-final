package spentcalories

import (
	"testing"
	"time"
)

func TestWalkingSpentCalories(t *testing.T) {
	tests := []struct {
		steps    int
		weight   float64
		duration float64
		want     float64
	}{
		{1000, 70, 30, (1000 * 70) / (1000.0 * (30.0 / 60.0))},
		{0, 70, 30, 0},
		{1000, 0, 30, 0},
	}

	for _, tt := range tests {
		got := WalkingSpentCalories(tt.steps, tt.weight, tt.duration)
		if got != tt.want {
			t.Errorf("WalkingSpentCalories(%d, %.1f, %.1f) = %.2f, want %.2f",
				tt.steps, tt.weight, tt.duration, got, tt.want)
		}
	}
}

func TestRunningSpentCalories(t *testing.T) {
	tests := []struct {
		steps    int
		weight   float64
		duration float64
		want     float64
	}{
		{2000, 70, 20, (2000 * 70 * 2) / (1000.0 * (20.0 / 60.0))},
		{0, 70, 20, 0},
		{2000, 0, 20, 0},
	}

	for _, tt := range tests {
		got := RunningSpentCalories(tt.steps, tt.weight, tt.duration)
		if got != tt.want {
			t.Errorf("RunningSpentCalories(%d, %.1f, %.1f) = %.2f, want %.2f",
				tt.steps, tt.weight, tt.duration, got, tt.want)
		}
	}
}

func TestTrainingInfo(t *testing.T) {
	tests := []struct {
		data    string
		weight  float64
		height  float64
		wantErr bool
	}{
		{"5000,бег,1h0m0s", 70, 175, false},
		{"6000,ходьба,30m0s", 80, 180, false},
		{"bad_data", 70, 175, true},
	}

	for _, tt := range tests {
		_, err := TrainingInfo(tt.data, tt.weight, tt.height)
		if (err != nil) != tt.wantErr {
			t.Errorf("TrainingInfo(%q) error = %v, wantErr %v",
				tt.data, err, tt.wantErr)
		}
	}
}

func TestParseTraining(t *testing.T) {
	tests := []struct {
		data      string
		wantSteps int
		wantType  string
		wantDur   time.Duration
		wantErr   bool
	}{
		{"5000,бег,1h0m0s", 5000, "бег", time.Hour, false},
		{"6000,ходьба,30m0s", 6000, "ходьба", 30 * time.Minute, false},
		{"0,бег,1h0m0s", 0, "бег", time.Hour, true},
		{"5000,бег,", 0, "", 0, true},
		{"abc,бег,1h0m0s", 0, "", 0, true},
		{"5000,,1h0m0s", 0, "", 0, true},
		{"5000,бег,xyz", 0, "", 0, true},
	}

	for _, tt := range tests {
		got, err := parseTraining(tt.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseTraining(%q) error = %v, wantErr %v",
				tt.data, err, tt.wantErr)
			continue
		}
		if err == nil {
			if got.Steps != tt.wantSteps {
				t.Errorf("parseTraining(%q) steps = %v, want %v",
					tt.data, got.Steps, tt.wantSteps)
			}
			if got.TrainingType != tt.wantType {
				t.Errorf("parseTraining(%q) type = %v, want %v",
					tt.data, got.TrainingType, tt.wantType)
			}
			if got.Duration != tt.wantDur {
				t.Errorf("parseTraining(%q) duration = %v, want %v",
					tt.data, got.Duration, tt.wantDur)
			}
		}
	}
}
