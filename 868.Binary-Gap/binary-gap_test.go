package Binary_Gap

import "testing"

func TestOutput(t *testing.T) {
	input := 22
	ret := binaryGap(input)
	if ret != 2 {
		t.Error("fail!")
	}
}
