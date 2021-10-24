package main
import "testing"
func test(t *testing.T){

	result := sum(2,3)

  if result != 5 {
		t.Error("The result most be 5 ")
	}
}
