package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := "ffykfhsq"
	found := 0
	counter := 0

	for found < 8 {
		hasher := md5.New()
		io.WriteString(hasher, fmt.Sprintf("%s%d", input, counter))
		hash := fmt.Sprintf("%x", hasher.Sum(nil))
		if strings.HasPrefix(hash, "00000") {
			fmt.Printf("%s", hash[5:6])
			found++
		}
		counter++
	}
}
