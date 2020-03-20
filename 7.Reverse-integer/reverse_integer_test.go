package reverse_integer

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	re := reverse(123)
	fmt.Printf("%v", re)
}