package trainings

import (
	"testing"
	"time"

	"github.com/Sevacoming/go1fl-4-sprint-final/internal/personaldata"
)

func TestTraining_ActionInfo(t *testing.T) {
	pd := personaldata.PersonalData{Name: "Иван", Weight: 70, Height: 175}

	tests := []struct {
		name     string
		training Training
		wantErr  bool
	}{
		{
			name: "Ходьба",
			training: Training{
				Steps:        4000,
				TrainingType: "ходьба",
				Duration:     time.Hour,
				Personal:     pd,
			},
			wantErr: false,
		},
		{
			name: "Бег",
			training: Training{
				Steps:        5000,
				TrainingType: "бег",
				Duration:     time.Hour,
				Personal:     pd,
			},
			wantErr: false,
		},
		{
			name: "Неизвестный тип",
			training: Training{
				Steps:        3000,
				TrainingType: "плавание",
				Duration:     time.Hour,
				Personal:     pd,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.training.ActionInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("ActionInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got == "" {
				t.Errorf("ActionInfo() вернул пустую строку")
			}
		})
	}
}
