package encrypt

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func HashingWithSalt(data string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("data : %s, salt : %s", data, salt)
	fmt.Println(saltedText)

	sha := sha1.New()
	sha.Write([]byte(saltedText))

	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func hash() {
	// hash
	// ecrypt = one way

	data := "rahasia"

	hash1, salt1 := HashingWithSalt(data)

	hash2, salt2 := HashingWithSalt(data)

	fmt.Println(hash1, salt1)
	fmt.Println(hash2, salt2)

	// data := "rahasia akan tetap rahasia"

	// sha := sha1.New()

	// sha.Write([]byte(data))

	// encrypted := sha.Sum(nil)
	// encryptedStr := fmt.Sprintf("%x", encrypted)

	// fmt.Println(encryptedStr)
}
