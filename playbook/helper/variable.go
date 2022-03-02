package helper

import "strings"

func FindVariable(val string) (bool, string) {
	fPos := strings.Index(val, "{{")
	if fPos == -1 {
		//fPos = strings.Index(val, "\"{{")
		//if fPos == -1 {
		//	return false, ""
		//}
		return false, ""
	}

	sPos := strings.Index(val, "}}")
	if sPos == -1 {
		//sPos = strings.Index(val, "}}\"")
		//if sPos == -1 {
		//	return false, ""
		//}
		return false, ""
	}

	word := strings.TrimSpace(val[fPos+2 : sPos])
	return true, word
}

func FindVariableAndReplaceWithValue(raw *string, variable *string, values []string) (bool, []string) {

	fPos := strings.Index(*raw, "{{")
	if fPos == -1 {
		return false, nil
	}
	sPos := strings.Index(*raw, "}}")
	if sPos == -1 {
		return false, nil
	}

	var rawVal string
	rawVal = *raw

	word := strings.TrimSpace(rawVal[fPos+2 : sPos])
	if word != *variable {
		return false, nil
	}

	var newValues []string
	for _, v := range values {
		newVal := strings.Replace(rawVal, rawVal[fPos:sPos+2], v, 1)
		newValues = append(newValues, newVal)
	}

	return true, newValues
}
