package Price

import (
	"testing"
)

func TestReturnPrice(t *testing.T) {
	t.Run("ReturnPrice" , func(t *testing.T) {
		t.Log("Enter Price")
		result := ReturnPrice("20")
		expext := 20
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})

	t.Run("Return0" , func(t *testing.T) {
		t.Log("test hoge")
		result := ReturnPrice("-")
		expext := 0
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})	
}