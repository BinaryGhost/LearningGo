package helpers

import (
	"log"
	"math/rand/v2"
	"net/http"
	"net/url"
	"strings"
)

func ShortURL(Url string) string {
	var prependium strings.Builder

	u, err := url.Parse(Url)
	if err != nil {
		log.Fatal(err)
	}

	protocol := u.Scheme
	prependium.WriteString(protocol)
	prependium.WriteString("://shortURL/")

	nameLen := len(u.Hostname()) / 2

	bytes := make([]byte, nameLen)
	for i := 0; i < nameLen; i++ {
		r := rand.N(uint8(255))

		if r == 0 {
			bytes[i] = 1
		} else {
			bytes[i] = r
		}
		prependium.WriteString(string(bytes[i]))
	}

	return prependium.String()
}

// To make things simple. I try to connect to this url,
// if it exists. it returns a successfull GET
func IsURL(url string) bool {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		//log.Fatal(err)
		return false
	}
	defer resp.Body.Close()

	return true
}
