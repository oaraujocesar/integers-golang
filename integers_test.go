package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	result := Add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Expected %d, got %d ", expected, result)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
