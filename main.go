package black_it

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type biObj struct {
	text  string
	level int
	fill  string
}

// [BI:█#123]TEXT[:BI]

// Mark just encrypt text
// can be used to save encrypted text to a database or document
// set pattern to default █
// description -> des[BI:█#1]cri[:BI]ption
//
//	text:the text that needs to be processed
//	level:the level at which it can be displayed
func Mark(text string, level int) string {
	return fmt.Sprintf("[BI:█#%d]%s[:BI]", level, text)
}

// MarkWithFill encrypt text with pattern
// can be used to save encrypted text to a database or document
// description -> des[BI:🙈]cri[:BI]ption
//
//	text:the text that needs to be processed
//	level:the level at which it can be displayed
//	fill:the pattern of the fill,emoji support
func MarkWithFill(text string, level int, fill string) string {
	return fmt.Sprintf("[BI:%s#%d]%s[:BI]", fill, level, text)
}

// Unmark just decrypt text
// des███ption -> description
//
//	text:the text that needs to be processed
//	level:the level at which it can be displayed
func Unmark(text string, level int) string {
	cp := regexp.MustCompile(`\[BI:.*?#\d+].*?\[:BI]`)
	return cp.ReplaceAllStringFunc(text, func(s string) string {
		o := getBiObj(s)
		if o.level <= level {
			return o.text
		} else {
			return s
		}
	})
}

func BlackIt(text string, level int) string {
	cp := regexp.MustCompile(`\[BI:.*?#\d+].*?\[:BI]`)
	return cp.ReplaceAllStringFunc(text, func(s string) string {
		o := getBiObj(s)
		if o.level <= level {
			return o.text
		} else {
			ns := ""
			for i := 0; i < len(o.text); i++ {
				ns += o.fill
			}
			return ns
		}
	})
}

func getBiObj(text string) *biObj {
	pattern := `^\[BI:.*?#\d+].*?\[:BI]$`
	b, _ := regexp.Match(pattern, []byte(text))
	if !b {
		return nil
	}

	text = strings.TrimPrefix(text, "[BI:")
	text = strings.TrimSuffix(text, "[:BI]")
	ss := strings.SplitN(text, "]", 2)

	ss1 := strings.Split(ss[0], "#")
	atoi, _ := strconv.Atoi(ss1[1])

	return &biObj{
		text:  ss[1],
		level: atoi,
		fill:  ss1[0],
	}
}
