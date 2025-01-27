package url

import (
	"crypto/sha256"
	"encoding/hex"
)

func Shorten(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	//fmt.Println(h.Sum(nil))
	hash := hex.EncodeToString(h.Sum(nil))
	shortURL := hash[:8]
	return shortURL
}
