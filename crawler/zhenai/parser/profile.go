package parser

import (
	"regexp"
	"strconv"
	"xyy/learngo/crawler/engine"
	"xyy/learngo/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xingzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业：</span><span field="">([^<]+)</span></td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	// name inherits from upper-class parser

	// age
	if age, err := strconv.Atoi(extractString(contents, ageRe)); err != nil {
		profile.Age = age
	}
	// height
	if height, err := strconv.Atoi(extractString(contents, heightRe)); err != nil {
		profile.Height = height
	}
	// weight
	if weight, err := strconv.Atoi(extractString(contents, weightRe)); err != nil {
		profile.Weight = weight
	}
	// income
	profile.Income = extractString(contents, incomeRe)
	// gender
	profile.Gender = extractString(contents, genderRe)
	// car
	profile.Car = extractString(contents, carRe)
	// education
	profile.Education = extractString(contents, educationRe)
	// hukou
	profile.Hukou = extractString(contents, hukouRe)
	// house
	profile.House = extractString(contents, houseRe)
	// marriage
	profile.Marriage = extractString(contents, marriageRe)
	// occupation
	profile.Occupation = extractString(contents, occupationRe)
	// xingzuo
	profile.Xingzuo = extractString(contents, xingzuoRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
