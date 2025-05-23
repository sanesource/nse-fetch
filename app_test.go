package nse

import (
	"testing"
)

func TestFetchAutoComplete(t *testing.T) {
	_, err := FetchAutoComplete("HDF")
	if err != nil {
		t.Errorf("FetchAutoComplete should not throw error: %v", err)
	}
}

func TestFetchEquityHistorical(t *testing.T) {
	_, err := FetchEquityHistorical("IDFCFIRSTB", "23-05-2024", "23-05-2025")
	if err != nil {
		t.Errorf("FetchEquityHistorical should not throw error: %v", err)
	}
}
