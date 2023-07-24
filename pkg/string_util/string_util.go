package string_util

import "github.com/valyala/fasttemplate"

const (
	startTag = "${"
	endTag   = "}"
)

func IsEmpty(s string) bool {
	return s == ""
}

// StrSubstitutor https://darjun.github.io/2021/05/24/godailylib/fasttemplate/
func StrSubstitutor(content string, valueMap map[string]interface{}) string {
	return fasttemplate.ExecuteString(content, startTag, endTag, valueMap)
}
