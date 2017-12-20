package strtools

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// Random returns a random string
func Random(l int) string {
	b, _ := randbytes(l)

	input := base64.StdEncoding.EncodeToString(b)
	r := strings.NewReplacer([]string{"/", "", "+", "", "=", ""}...)
	output := r.Replace(input)

	return Substr(output, 0, l)
}

func randbytes(len int) ([]byte, error) {
	var b = make([]byte, len)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}, err
	}
	return b[:], nil
}
