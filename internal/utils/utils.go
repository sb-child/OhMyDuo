package utils

import (
	"context"
	"my-duo/internal/consts"
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

func SplitText(s string) []consts.SpiltTextPiece {
	var result []consts.SpiltTextPiece
	var current string
	for _, char := range s {
		if char <= 127 { // ASCII character
			if (unicode.IsSpace(char) || unicode.IsSymbol(char)) && current != "" {
				result = append(result, consts.SpiltTextPiece{Text: current, Unicode: false})
				result = append(result, consts.SpiltTextPiece{Text: string(char), Unicode: false})
				current = ""
			} else {
				current += string(char)
			}
		} else { // Unicode character
			if current != "" {
				result = append(result, consts.SpiltTextPiece{Text: current, Unicode: false})
				current = ""
			}
			result = append(result, consts.SpiltTextPiece{Text: string(char), Unicode: true})
		}
	}
	if current != "" {
		result = append(result, consts.SpiltTextPiece{Text: current, Unicode: false})
	}
	return result
}
