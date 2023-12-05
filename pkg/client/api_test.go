package client

import (
	"testing"

	"strings"

	"github.com/prometheus/common/model"
)

func TestIntToStringReturnsString(t *testing.T) {
	t.Parallel()
	time := model.Now()
	got := IntToString(time)
	result := strings.Split(got, ".")
	if len(result) > 1 {
		t.Fatalf("string contains a decimal point: %s", got)
	}
}
