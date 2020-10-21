package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const SALT = "qzyelonm" //
const SIZE = 35000

func main() {
	var hashes [SIZE]string
	for i := 0; i < SIZE; i++ {
		hashes[i] = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", SALT, i))))
	}
	foundKeys := 0
	for id, hash := range hashes {
		for i := 0; i < len(hash)-2; i++ {
			if hash[i:i+1] == hash[i+1:i+2] && hash[i:i+1] == hash[i+2:i+3] {
				for j := id + 1; j <= id+1000; j++ {
					if j >= SIZE {
						continue
					}
					char := hash[i : i+1]
					if strings.Index(hashes[j], fmt.Sprintf("%s%s%s%s%s", char, char, char, char, char)) >= 0 {
						foundKeys++
						if foundKeys == 64 {
							fmt.Printf("part 1 => %d\n", id)
						}
					}
				}
				break
			}
		}
	}
}
