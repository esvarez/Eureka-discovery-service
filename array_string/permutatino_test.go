package str

import "testing"

func TestIsPermutation(t *testing.T) {
	tests := map[string]struct {
		str, str2 string
		want      bool
	}{
		"1": {"", "", true},
		"2": {"a", "a", true},
		"3": {"a", "b", false},
		"4": {"a", "ab", false},
		"5": {"abc", "bca", true},
		"6": {"abc", "adc", false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := isPermutation(tc.str, tc.str2)
			if got != tc.want {
				t.Errorf("isPermutation(%q, %q) = %v, want %v", tc.str, tc.str2, got, tc.want)
			}
		})
	}

}
