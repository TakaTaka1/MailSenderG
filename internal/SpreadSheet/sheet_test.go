package spreadSheet

import (
	"testing"
)

func TestRetSheet(t *testing.T) {
	t.Run("RetSheetNameStruct", func(t *testing.T) {
		t.Log("Sample test")
		result := RetSheetNameStruct()
		expext := "食費"
		if result.FoodCost != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})
}
