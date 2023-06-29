// package util
// @Author: dsg
// @Description:
// @File:  validate.go
// @Date: 2023/6/30 2:47

package util

import (
	"regexp"
	"strconv"
	"strings"
)

func ValidateIDCard(card string) bool {

	regRuler := `^[1-9]\d{5}(18|19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])\d{3}([0-9Xx])$`
	//regRuler := "(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)"

	regex := regexp.MustCompile(regRuler)

	if !regex.MatchString(card) {
		return false
	}

	// 校验身份证号码校验码
	factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	sum := 0
	for i := 0; i < 17; i++ {
		num, _ := strconv.Atoi(string(card[i]))
		sum += num * factor[i]
	}

	mod := sum % 11
	checkCode := checkCodes[mod]

	if strings.ToUpper(string(card[17])) != checkCode {
		return false
	}

	return true
}

func ValidatePhone(phone string) bool {
	reg := `^(13[0-9]|14[01456879]|15[0-3,5-9]|16[2567]|17[0-8]|18[0-9]|19[0-3,5-9])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

// 统一社会信用代码
func ValidateSocialCode(code string) bool {
	// 统一社会信用代码正则表达式
	pattern := `^[0-9ABCDEFGHJKLMNPQRTUWXY]{2}\d{6}[0-9ABCDEFGHJKLMNPQRTUWXY]{10}$`
	match, err := regexp.MatchString(pattern, code)
	if err != nil {
		return false
	}
	// 校验码校验
	if match {
		weights := []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
		codes := "0123456789ABCDEFGHJKLMNPQRTUWXY"
		sum := 0
		for i, c := range code[:17] {
			w := weights[i]
			idx := strings.IndexRune(codes, c)
			sum += idx * w
		}
		check := 31 - sum%31
		if check == 31 {
			check = 0
		}
		return string(codes[check]) == code[17:]
	}
	return false
}
