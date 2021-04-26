package encrypt

import (
	"encoding/base64"
	"fmt"
)

func encodeDecode() {
	//code
	var code = "rahasia"
	// casting dari string ke byte
	byteCode := []byte(code)
	// encode base64
	encodeStr := base64.StdEncoding.EncodeToString(byteCode)
	fmt.Println(encodeStr)

	// decode (convert dari encode ke data awal)
	decodeStr, err := base64.StdEncoding.DecodeString(encodeStr)
	if err != nil {
		fmt.Println("error")
		return
	}
	// casting dari byte ke string
	strDecode := string(decodeStr)
	fmt.Println(strDecode)

	//url(URLEncoding)

	var link = "https://www.google.com"

	encodeStr2 := base64.URLEncoding.EncodeToString([]byte(link))

	fmt.Println(encodeStr)

	// decoding url
	decodeStr2, err := base64.RawStdEncoding.DecodeString(encodeStr2)

	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(string(decodeStr2))
}
