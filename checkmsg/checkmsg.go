package checkmsg

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func CheckExclusionWord(fn, txt string) error {
	existsExclusionWord := errors.New("Found Exclusion Word")

	var exclusionWords []string
	file, err := os.Open(fn)

	if err != nil {
		return err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		exclusionWords = append(exclusionWords, sc.Text())
	}

	for _, e := range exclusionWords {
		if strings.Index(txt, e) != -1 {
			return existsExclusionWord
		}
	}

	return nil
}
