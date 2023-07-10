package questions_test

import (
	"testing"

	"github.com/irhamsahbana/cmlabs-backend-fulltime-test/questions"
)

func TestPalindrome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "SAIPPUAKIVIKAUPPIAS is palindrome",
			arg:  "SAIPPUAKIVIKAUPPIAS",
			want: true,
		},
		{
			name: "Aibohphobia is palindrome",
			arg:  "Aibohphobia",
			want: true,
		},
		{
			name: "Anna is palindrome",
			arg:  "Anna",
			want: true,
		},
		{
			name: "My gym not palindrome",
			arg:  "My gym",
			want: false,
		},
		{
			name: "No lemon, no melon not palindrome",
			arg:  "No lemon, no melon",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := questions.Palindrome(tt.arg); got != tt.want {
				t.Errorf("Palindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
