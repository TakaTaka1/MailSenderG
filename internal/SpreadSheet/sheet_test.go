package SpreadSheet

import (
	"testing"
)

func TestRetSheet (t *testing.T) {	
	t.Run("RetSheet" , func(t *testing.T) {
		t.Log("Sample test")
		result := RetSheet()
		expext := "Sheet"
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})	
}