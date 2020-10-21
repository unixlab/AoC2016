package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"sync"
)

// const SALT = "abc" // 22728 & 22551
// const SALT = "zpqevtbw" // 16106 & 22423
const SALT = "qzyelonm" // 15168 & 20864

const SIZE = 25000

func stretchedMd5Hash(data string) string {
	for i := 0; i < 2017; i++ {
		data = fmt.Sprintf("%x", md5.Sum([]byte(data)))
	}
	return data
}

func findKey64(hashes [SIZE]string) int {
	foundKeys := 0
	for id, hash := range hashes {
		for i := 0; i < len(hash)-2; i++ {
			char := hash[i : i+1]
			if char == hash[i+1:i+2] && char == hash[i+2:i+3] {
				for j := id + 1; j <= id+1000; j++ {
					if j >= SIZE {
						continue
					}
					if strings.Index(hashes[j], fmt.Sprintf("%s%s%s%s%s", char, char, char, char, char)) >= 0 {
						foundKeys++
						if foundKeys == 64 {
							return id
						}
						break
					}
				}
				break
			}
		}
	}
	return -1
}

func main() {
	var hashes [SIZE]string
	for i := 0; i < SIZE; i++ {
		hashes[i] = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", SALT, i))))
	}
	fmt.Printf("part 1 => %d\n", findKey64(hashes))

	var threads sync.WaitGroup
	for i := 0; i < SIZE; i += 5000 {
		threads.Add(1)
		go func(start int) {
			for j := start; j < start+5000; j++ {
				hashes[j] = stretchedMd5Hash(fmt.Sprintf("%s%d", SALT, j))
			}
			threads.Done()
		}(i)
	}
	threads.Wait()
	fmt.Printf("part 2 => %d\n", findKey64(hashes))
}
