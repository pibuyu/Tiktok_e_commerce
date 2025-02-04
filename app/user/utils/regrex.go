package utils

import "regexp"

// VerifyMobileFormat , Verify format of phone number
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// VerifyEmailFormat , Verify format of email
func VerifyEmailFormat(email string) bool {
	regex := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	return regex.MatchString(email)
}

// VerifyPasswordFormat , password format require:length between[6,20],must contain both digits and letters
func VerifyPasswordFormat(password string) bool {
	pattern := `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{6,20}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(password)
}
