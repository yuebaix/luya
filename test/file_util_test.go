package test

import (
	"github.com/yuebaix/luya/internal/pkg/util"
	"testing"
)

func TestCsv(t *testing.T) {
	util.CreateFie("test.csv")
	util.WriteCsv("test.csv", []string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})
}
