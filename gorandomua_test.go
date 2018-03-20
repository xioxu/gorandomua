package gorandomua

import (
	"testing"
	"fmt"
)

func TestGetRandom(t *testing.T) {
	ua := GetRandom()
	fmt.Print(ua)
	if ua == ""{
		t.Fail()
	}
}