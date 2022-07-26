package Price

import (
	"testing"
)

func TestReturnPrice(t *testing.T) {
	t.Run("ReturnPrice" , func(t *testing.T) {
		t.Log("Enter Price")
		result := ReturnPrice("20")
		expect := 20
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})

	t.Run("Return0" , func(t *testing.T) {
		t.Log("test hoge")
		result := ReturnPrice("-")
		expect := 0
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})	
}

func TestCheckDiffPrice(t *testing.T) {
	t.Run("CheckDiffPrice" , func(t *testing.T) {
		t.Log("MPrice >= TPrice")
		result := CheckDiffPrice(100,10)
		expect := 90
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})
	t.Run("CheckDiffPrice" , func(t *testing.T) {
		t.Log("MPrice < TPrice")
		result := CheckDiffPrice(10,100)
		expect := 90
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})

	t.Run("CheckDiffPrice" , func(t *testing.T) {
		t.Log("MPrice == TPrice")
		result := CheckDiffPrice(100,100)
		expect := 0
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}
	})	
	
}

func TestCheckCost(t *testing.T) {
	t.Run("CheckCost", func(t *testing.T) {
		t.Log("Cost is filled")
		result := CheckCost("-")
		expect := "光熱費を入力してください"
		if result != expect {
			t.Error("\nActual： ", result, "\nExpectation： ", expect)
		}		
	})
}