package str

import "testing"

func TestIsUnique(t *testing.T) {
	tests := map[string]struct {
		word     string
		expected bool
	}{
		"unique chars": {
			word:     "abcdef",
			expected: true,
		},
		"duplicate chars": {
			word:     "abcdeff",
			expected: false,
		},
		"empty string": {
			word:     "",
			expected: true,
		},
		"duplicated chars not sorted": {
			word:     "abca",
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := isUnique(tc.word)
			if actual != tc.expected {
				t.Errorf("got %v, want %v", actual, tc.expected)
			}
		})
	}
}
