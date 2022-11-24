package lib

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

// ReadFile takes the path parameter which is a string of the path were the file is located
// return a byte array with the content and if necessary an error.
func ReadFile(path string) ([]byte, error) {
	regexCwd := regexp.MustCompile(`(?m)\./.*\n`) // regexCwd checks if it is in the current working directory.
	regexSys := regexp.MustCompile(`(?m)/home.*\n|/dev.*\n|/sys/.*\n`)
	// FIXME: At the moment this always returns the error.
	if regexCwd.Find([]byte(path)) != nil {
		parentPath, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		pullPath := filepath.Join(parentPath, path)
		file, err := os.Open(pullPath)
		if err != nil {
			return nil, err
		}

		defer func(file *os.File) {
			err2 := file.Close()
			if err2 != nil {
				fmt.Println(err2)
			}
		}(file)
		return read(file)
	} else if regexSys.Find([]byte(path)) != nil {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		defer func(file *os.File) {
			err2 := file.Close()
			if err2 != nil {
				fmt.Println(err2)
			}
		}(file)
		return read(file)
	} else {
		err := errors.New("the given path isn't valid to read from")
		return nil, err
	}
}

func read(fdR io.Reader) ([]byte, error) {
	br := bufio.NewReader(fdR)
	var buf bytes.Buffer

	for {
		ba, isPrefix, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buf.Write(ba)
		if !isPrefix {
			buf.WriteByte('\n')
		}

	}
	return buf.Bytes(), nil
}
