package currency

import (
	"os"
	"strings"
	"testing"
)

func TestFindUSDRateFromRealExampleJSON(t *testing.T) {
	// Read the test data from the file
	currencyRateFile, err := os.ReadFile("currency_rate_example_NBU.json")
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	// Create a reader with the test data
	data := strings.NewReader(string(currencyRateFile))

	rate, err := FindUSDRateNBU(data)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if rate != 39.427200 {
		t.Errorf("Expected rate 39.427200, got %f", rate)
	}
}

func TestFindUSDRate(t *testing.T) {
	var tests = []struct {
		name      string  // name of the test case
		jsonData  string  // input JSON data
		expected  float64 // expected rate
		expectErr bool    // whether an error is expected
	}{
		{
			name: "Valid USD Rate",
			jsonData: `[
                {"r030":840, "txt":"Долар США", "rate":27.11, "cc":"USD", "exchangedate":"20.05.2024"}
            ]`,
			expected:  27.11,
			expectErr: false,
		},
		{
			name: "No USD in List",
			jsonData: `[
                {"r030":978, "txt":"Euro", "rate":31.11, "cc":"EUR", "exchangedate":"20.05.2024"}
            ]`,
			expected:  0,
			expectErr: true,
		},
		{
			name: "Invalid JSON Format",
			jsonData: `[
                {"r030":840, "txt":"Долар США", "rate":, "cc":"USD", "exchangedate":"20.05.2024"}
            ]`,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "Empty JSON List",
			jsonData:  `[]`,
			expected:  0,
			expectErr: true,
		},
		{
			name: "Multiple Entries Including USD",
			jsonData: `[
                {"r030":978, "txt":"Euro", "rate":31.11, "cc":"EUR", "exchangedate":"20.05.2024"},
                {"r030":840, "txt":"Долар США", "rate":28.22, "cc":"USD", "exchangedate":"20.05.2024"}
            ]`,
			expected:  28.22,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.jsonData)
			rate, err := FindUSDRateNBU(reader)
			if (err != nil) != tt.expectErr {
				t.Fatalf("Test %s expected error %v, got %v", tt.name, tt.expectErr, err)
			}
			if !tt.expectErr && rate != tt.expected {
				t.Errorf("Test %s expected rate %f, got %f", tt.name, tt.expected, rate)
			}
		})
	}
}
