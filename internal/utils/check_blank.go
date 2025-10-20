package utils

func CheckBlank(parameter string) string {

	if parameter == "" {
		return "err"
	}
	return parameter
}
