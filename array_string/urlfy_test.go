package str

import "testing"

func TestUrlfy(t *testing.T) {
	tests := []struct {
		input string
		url   string
	}{
		{"Mr John Smith    ", "Mr%20John%20Smith"},
		{"", ""},
	}

	for _, test := range tests {
		if got := urlfy(test.input); got != test.url {
			t.Errorf("urlfy(%q) = %q, want %q", test.input, got, test.url)
		}
	}
}
