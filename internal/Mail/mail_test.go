package Mail

import (
	"testing"
)

func TestRetMail(t *testing.T) {
	t.Run("RetMail", func(t *testing.T) {
		t.Log("Sample test")
		result := RetMail()
		expext := "Mail"
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})
}
