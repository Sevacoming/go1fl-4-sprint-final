package activity

import "testing"

func TestNew_Valid(t *testing.T) {
	a, err := New("Бег", 30, 5)
	if err != nil {
		t.Errorf("ожидали nil, получили ошибку: %v", err)
	}
	if a == nil {
		t.Errorf("ожидали объект Activity, получили nil")
	}
}

func TestNew_InvalidType(t *testing.T) {
	_, err := New("", 30, 5)
	if err == nil {
		t.Errorf("ожидали ошибку для пустого типа активности, получили nil")
	}
}

func TestNew_InvalidDuration(t *testing.T) {
	_, err := New("Ходьба", 0, 3)
	if err == nil {
		t.Errorf("ожидали ошибку для длительности <= 0, получили nil")
	}
}

func TestNew_InvalidDistance(t *testing.T) {
	_, err := New("Ходьба", 30, -1)
	if err == nil {
		t.Errorf("ожидали ошибку для отрицательной дистанции, получили nil")
	}
}
