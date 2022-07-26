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

func TestCheckDiffPrice(t *testing.T) {
	t.Run("CheckDiffPrice" , func(t *testing.T) {
		t.Log("MPrice >= TPrice")
		result := CheckDiffPrice(100,10)
		expext := 90
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})
	t.Run("CheckDiffPrice" , func(t *testing.T) {
		t.Log("MPrice < TPrice")
		result := CheckDiffPrice(10,100)
		expext := 90
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})

	t.Run("CheckDiffPrice" , func(t *testing.T) {
		t.Log("MPrice == TPrice")
		result := CheckDiffPrice(100,100)
		expext := 0
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}
	})	
	
}

func TestCheckCost(t *testing.T) {
	t.Run("CheckCost", func(t *testing.T) {
		t.Log("Cost is filled")
		result := CheckCost("-")
		expect := "光熱費を入力してください"
		if result != expext {
			t.Error("\nActual： ", result, "\nExpectation： ", expext)
		}		
	})
}