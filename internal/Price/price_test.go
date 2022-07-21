package Price

import (
	"testing"
)

func TestPrice(t *testing.T) {
	result := PrintPrice(1)
	expext := 1
	if result != expext {
	  t.Error("\n実際： ", result, "\n理想： ", expext)
	}
  
	t.Log("TestHello終了")
}