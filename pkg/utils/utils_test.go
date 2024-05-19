package utils

import "testing"

func TestValidEmail(t *testing.T) {
	var tests = []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid Email Simple",
			email:    "example@example.com",
			expected: true,
		},
		{
			name:     "Valid Email with Underscore and some Domain",
			email:    "test_email@example.co.org",
			expected: true,
		},
		{
			name:     "Valid Email with Hyphen and Numeric Subdomain",
			email:    "test-email@123.abc.com",
			expected: true,
		},
		{
			name:     "Invalid Email Missing Domain",
			email:    "invalid-email@",
			expected: false,
		},
		{
			name:     "Invalid Email Missing @",
			email:    "email-domain.com",
			expected: false,
		},
		{
			name:     "Invalid Email Ending with Dot",
			email:    "email@domain.com.",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidEmail(tt.email); got != tt.expected {
				t.Errorf("%s: isValidEmail(%q) = %v, want %v", tt.name, tt.email, got, tt.expected)
			}
		})
	}
}
