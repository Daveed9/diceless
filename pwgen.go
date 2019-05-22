package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// converts wordlist to slice
// precondition: input (file) one word per line
func fileSlice(r io.Reader) (wordlist []string, err error) {
	scan := bufio.NewScanner(r)
	wordList := make([]string, 0)
	for scan.Scan() {
		wordList = append(wordList, removeExtras(scan.Text()))
	}
	return wordList, scan.Err()
}

// removes all non letter characters from a string
// allows for diceware and non-diceware wordlists with and without extra chars
func removeExtras(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	check(err)
	return reg.ReplaceAllString(str, "")
}

var (
	flagFilePath string
	flagLength   int
	password     string
)

func init() {
	flag.StringVar(&flagFilePath, "file", "eff_large_wordlist.txt", "path/to/wordlist")
	flag.IntVar(&flagLength, "len", 6, "length of password in words")
	flag.Parse()
}

func main() {
	password := ""

	wordList, err := os.Open(flagFilePath)
	check(err)
	defer wordList.Close()

	words, err := fileSlice(wordList)
	check(err)

	len := big.NewInt(int64(len(words)))

	for i := 0; i < flagLength; i++ {
		num, err := rand.Int(rand.Reader, len)
		check(err)
		// mega sketchy big.Int to int conversion
		indexStr := num.String()
		index, err := strconv.ParseInt(indexStr, 10, 0)
		check(err)

		password += words[index] + " "
	}

	fmt.Println("Password:", password)
}
