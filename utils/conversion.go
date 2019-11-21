package utils

func ConvertOutput(output []byte, err error) (string, error) {
	return string(output), err
}