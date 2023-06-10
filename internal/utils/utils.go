package utils

import (
	"context"
	"unicode"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func GetResource(path string) []byte {
	if gres.Contains(path) {
		// g.Log().Warningf(context.Background(), "file found")
		return gres.GetContent(path)
	}
	if gfile.IsFile(path) {
		g.Log().Warningf(context.Background(), "utils.GetResource: %s is not exist in resource object, but found in filesystem.", path)
		return gfile.GetBytes(path)
	}
	return nil
}
func CapitalizeFirst(s string) string {
	for i, c := range s {
		if unicode.IsLetter(c) {
			return string(unicode.ToUpper(c)) + s[i+1:]
		}
	}
	return s
}
