package personaldata

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

type fakeParser struct {
	parseErr      error
	actionInfo    string
	actionInfoErr error
}

func (f *fakeParser) Parse(s string) error {
	return f.parseErr
}

func (f *fakeParser) ActionInfo() (string, error) {
	if f.actionInfoErr != nil {
		return "", f.actionInfoErr
	}
	return f.actionInfo, nil
}

func TestInfo_OK(t *testing.T) {
	data := []string{"2025-09-13;5000"}
	fp := &fakeParser{actionInfo: "шаги: 5000"}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	Info(data, fp) //что за функция?

	got := buf.String()
	if !strings.Contains(got, "шаги: 5000") {
		t.Errorf("ожидали в выводе шаги 5000, получили: %s", got)
	}
}

func TestInfo_ParseError(t *testing.T) {
	data := []string{"2025-09-13;5000"}
	fp := &fakeParser{parseErr: fmt.Errorf("тестовая ошибка")}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	Info(data, fp) //что за функция?

	got := buf.String()
	if !strings.Contains(got, "Ошибка парсинга") {
		t.Errorf("ожидали сообщение об ошибке парсинга, получили: %s", got)
	}
}

func TestInfo_ActionInfoError(t *testing.T) {
	data := []string{"2025-09-13;5000"}
	fp := &fakeParser{actionInfoErr: fmt.Errorf("тестовая ошибка")}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	Info(data, fp) //что за функция?

	got := buf.String()
	if !strings.Contains(got, "Ошибка ActionInfo") {
		t.Errorf("ожидали сообщение об ошибке ActionInfo, получили: %s", got)
	}
}
