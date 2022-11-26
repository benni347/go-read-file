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
	regexCwd := regexp.MustCompile(`(?m)\./.*\n|\./.*`)                                     // regexCwd checks if it is in the current working directory.
	regexSys := regexp.MustCompile(`(?m)/home.*\n|/home.*|/dev.*\n|/dev.*|/sys.*\n|/sys.*`) // regexSys check if it is a system path.
	if regexCwd.Find([]byte(path)) != nil {
		parentPath, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		fullPath := filepath.Join(parentPath, path)
		file, err := os.Open(fullPath)
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
	bufioReader := bufio.NewReader(fdR)
	var buffer bytes.Buffer

	for {
		readLine, isPrefix, err := bufioReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buffer.Write(readLine)
		if !isPrefix {
			buffer.WriteByte('\n')
		}

	}
	return buffer.Bytes(), nil
}
