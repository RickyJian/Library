package util

import (
	"crypto/sha512"
	"fmt"
)

func StringToSHA512(data string) string {
	h := sha512.New()
	h.Write([]byte(data))
	return fmt.Sprintf("%x",h.Sum(nil))
}
