package casbin

import (
	"strings"
	"regexp"
	"bytes"
)

func escape(s string) string {
	return strings.Replace(s, ".", "_", -1)
}

func fixAttribute(s string) string {
	reg := regexp.MustCompile("r\\.sub\\.([A-Za-z0-9]*)")
	res := reg.ReplaceAllString(s, "subAttr(r.sub, \"$1\")")

	reg = regexp.MustCompile("r\\.obj\\.([A-Za-z0-9]*)")
	res = reg.ReplaceAllString(res, "objAttr(r.obj, \"$1\")")

	return res
}

func arrayEquals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func array2DEquals(a [][]string, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !arrayEquals(v, b[i]) {
			return false
		}
	}
	return true
}

func arrayRemoveDuplicates(s *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *s {
		if !found[x] {
			found[x] = true
			(*s)[j] = (*s)[i]
			j++
		}
	}
	*s = (*s)[:j]
}

func arrayToString(s []string) string {
	var tmp bytes.Buffer
	for i, v := range s {
		if i != len(s)-1 {
			tmp.WriteString(v + ", ")
		} else {
			tmp.WriteString(v)
		}
	}
	return tmp.String()
}