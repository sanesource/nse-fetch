package nse

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogResponse(message string, data map[string]any, maxLines int) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(message)
	lines := make([]string, 0)
	currentLine := ""
	for _, b := range jsonData {
		currentLine += string(b)
		if b == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
			if len(lines) == maxLines {
				break
			}
		}
	}
	// If less than 5 lines and last line is not empty, add it
	if len(lines) < maxLines && currentLine != "" {
		lines = append(lines, currentLine)
	}
	for _, line := range lines {
		fmt.Print(line)
	}
	fmt.Println("...")
}

// ValidateNoErrorInData checks if the response contains the "Missing index or key." error
// and fails the test if it does
func ValidateNoErrorInData(t *testing.T, data map[string]any) {
	if data == nil {
		t.Fatal("Response data is nil")
	}

	// Check if data contains nested error structure: {"data": {"data": "Missing index or key."}}
	if dataField, ok := data["data"].(map[string]any); ok {
		if nestedData, ok := dataField["data"].(string); ok {
			if nestedData == "Missing index or key." {
				t.Fatal("API returned error: Missing index or key.")
			}
		}
	}

	// Also check for direct error message in data field
	if dataField, ok := data["data"].(string); ok {
		if dataField == "Missing index or key." {
			t.Fatal("API returned error: Missing index or key.")
		}
	}
}

func TestFetchAutoComplete(t *testing.T) {
	data, err := FetchAutoComplete("HDF")
	LogResponse("[Test FetchAutoComplete] Response: ", data, 5)
	if err != nil {
		t.Errorf("FetchAutoComplete should not throw error: %v", err)
	}
	ValidateNoErrorInData(t, data)
}

func TestFetchEquityHistorical(t *testing.T) {
	data, err := FetchEquityHistorical("IDFCFIRSTB", "23-05-2024", "23-05-2025")
	LogResponse("[Test FetchEquityHistorical] Response: ", data, 5)
	if err != nil {
		t.Errorf("FetchEquityHistorical should not throw error: %v", err)
	}
	ValidateNoErrorInData(t, data)
}

func TestFetchNifty50Performance(t *testing.T) {
	data, err := FetchNifty50Performance()
	LogResponse("[Test FetchNifty50Performance] Response: ", data, 5)
	if err != nil {
		t.Errorf("FetchNifty50Performance should not throw error: %v", err)
	}
	ValidateNoErrorInData(t, data)
}

func TestFetchTopGainers(t *testing.T) {
	data, err := FetchTopGainers()
	LogResponse("[Test FetchTopGainers] Response: ", data, 5)
	if err != nil {
		t.Errorf("FetchTopGainers should not throw error: %v", err)
	}
	ValidateNoErrorInData(t, data)
}

func TestFetchTopLosers(t *testing.T) {
	data, err := FetchTopLosers()
	LogResponse("[Test FetchTopLosers] Response: ", data, 5)
	if err != nil {
		t.Errorf("FetchTopLosers should not throw error: %v", err)
	}
	ValidateNoErrorInData(t, data)
}
