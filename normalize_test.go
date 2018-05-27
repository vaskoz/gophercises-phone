package phone

import "testing"

func TestNormalize(t *testing.T) {
	testcases := []struct {
		in, out string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}

	for _, test := range testcases {
		t.Run(test.in, func(t *testing.T) {
			t.Parallel()
			if out := Normalize(test.in); test.out != out {
				t.Errorf("Expected %s but got %s", test.out, out)
			}
		})
	}
}
