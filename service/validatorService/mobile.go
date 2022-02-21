package validatorService

import "regexp"

func (v *validService) IsMobile(data interface{}) bool {
	re, err := regexp.Compile(`^1[0-9]{10}$`)
	if err != nil {
		return false
	}
	switch data.(type) {
	case int, int32, int64, uint, uint32, uint64, string:
		matchString := re.MatchString(data.(string))
		return matchString
	default:
		return false
	}
}
