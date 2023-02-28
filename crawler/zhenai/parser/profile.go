package parser

import (
	"regexp"
	"strconv"
	"xyy/learngo/crawler/engine"
	"xyy/learngo/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	// age
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
		profile.Age = age
	}

	// marriage
	profile.Marriage = extractString(contents, marriageRe)

	return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
