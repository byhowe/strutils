package strutils

import "unicode/utf8"
import "errors"

/*
 * Basic string utility function for demo purposes.
 */

var ErrInvalidUTF8 = errors.New("input is not a valid UTF-8")

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, ErrInvalidUTF8
	}
	r := []rune(s)
	lr := len(r)
	ss := make([]rune, lr)

	for i := 0; i < lr; i++ {
		ss[lr-1-i] = r[i]
	}

	return string(ss), nil
}
