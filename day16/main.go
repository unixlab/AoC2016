package main

import (
	"fmt"
	"strings"
)

// from https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func replace01(s string) string {
	s = strings.Replace(s, "0", "2", -1)
	s = strings.Replace(s, "1", "0", -1)
	s = strings.Replace(s, "2", "1", -1)
	return s
}

func generateData(s string, length int) string {
	var data strings.Builder
	data.WriteString(s)
	for data.Len() < length {
		data.WriteString("0")
		data.WriteString(replace01(reverseString(data.String()[:data.Len()-1])))
	}
	return data.String()[:length]
}

func getChecksum(s string) string {
	var checksum strings.Builder
	s = s
	even := true
	for even {
		for i := 0; i < len(s); i += 2 {
			if s[i] == s[i+1] {
				checksum.WriteString("1")
			} else {
				checksum.WriteString("0")
			}
		}
		if checksum.Len()%2 == 1 {
			even = false
		}
		s = checksum.String()
		checksum.Reset()
	}
	return s
}

func main() {
	fmt.Printf("example => %s\n", getChecksum(generateData("10000", 20)))
	fmt.Printf("part 1  => %s\n", getChecksum(generateData("11110010111001001", 272)))
	fmt.Printf("part 2  => %s\n", getChecksum(generateData("11110010111001001", 35651584)))
}
