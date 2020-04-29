package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	passwordPart2 := make([]string, 8)
	input := "ffykfhsq"
	foundPart1 := 0
	foundPart2 := 0
	counter := 0

	for foundPart1 < 8 || foundPart2 < 8 {
		hasher := md5.New()
		io.WriteString(hasher, fmt.Sprintf("%s%d", input, counter))
		hash := fmt.Sprintf("%x", hasher.Sum(nil))
		if strings.HasPrefix(hash, "00000") {
			if foundPart1 < 8 {
				fmt.Printf("%s", hash[5:6])
				foundPart1++
			}

			if foundPart2 < 8 {
				pos, err := strconv.Atoi(hash[5:6])
				if err == nil && pos < 8 && passwordPart2[pos] == "" {
					passwordPart2[pos] = hash[6:7]
					foundPart2++
				}
			}
		}
		counter++
	}
	fmt.Println()
	fmt.Println(strings.Join(passwordPart2, ""))
}
