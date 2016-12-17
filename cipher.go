package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	args := os.Args

	if len(args) < 3 {
		log.Fatal("not enough argument")
	}

	var path string
	var message string

	switch args[2] {
	case "d":
		path = "decrypt.txt"
		message = "encrypted"
	case "e":
		path = "encrypt.txt"
		message = "dedcrypted"
	default:
		log.Fatal("unknown mode, mode flag must be d or e")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter cipher key: ")
	key, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	strLines, err := readLines(args[1])
	if err != nil {
		panic(err)
	}
	k, err := strconv.Atoi(string(key))
	if err != nil {
		panic(err)
	}
	err = writeLines(strLines, path, args[2], k)

	if err == nil {
		log.Printf("%v", message)
	} else {
		log.Printf("%v", "error")
	}

}

//EncryptCipher encrypts text string with optimized algorithm based on caesar cipher
func EncryptCipher(v string, key int) (string, error) {

	var enStr string
	var enCipher int
	var shift int

	for i := 0; i < len(v); i++ {
		if v[i] < 32 || v[i] > 126 {
			return enStr, errors.New("not a valid character")
		}
		r := key + i + 1

		if (int(v[i]) + r) > 126 {
			shift = (int(v[i]) + r)
			enCipher = shift - 95
		} else {

			enCipher = (int(v[i]) + r)
		}
		enStr = enStr + string(enCipher)
	}

	return enStr, nil
}

//DecryptCipher decrypt text string
func DecryptCipher(v string, key int) (string, error) {

	var decStr string
	var decCipher int
	var shift int
	for i := 0; i < len(v); i++ {
		if v[i] < 32 || v[i] > 126 {
			return decStr, errors.New("not a valid character")
		}
		r := key + i + 1

		if (int(v[i]) - r) < 32 {
			shift = (int(v[i]) - r)
			decCipher = shift + 95
		} else {

			decCipher = (int(v[i]) - r)
		}
		decStr = decStr + string(decCipher)
	}

	return decStr, nil
}

//readLines reads lines from a file and returns slice of string
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

//writeLines writes to a new file
func writeLines(lines []string, path string, t string, key int) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var str string
	w := bufio.NewWriter(file)

	if t == "e" {
		for _, line := range lines {

			str, err = EncryptCipher(line, key)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(w, str)

		}
	} else {
		for _, line := range lines {

			str, err = DecryptCipher(line, key)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(w, str)

		}
	}

	return w.Flush()
}
