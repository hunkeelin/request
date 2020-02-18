package request

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	rb := RequestBuilder{}
	rb.SetURL("wrongurl")
	err := rb._check()
	if err != nil {
		fmt.Println(err)
	}

}
