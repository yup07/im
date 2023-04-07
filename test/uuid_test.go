package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	u := uuid.NewV4()

	S := fmt.Sprintf("%x", u)
	fmt.Println(S)
}
