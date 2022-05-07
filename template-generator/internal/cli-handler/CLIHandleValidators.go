package clihandler

func ValueInList(values []string) func(string) bool {
	return func(input string) bool {
		// input exist in list
		for _, v := range values {
			if v == input {
				return true
			}
		}
		return false
	}
}

func NotEmpty() func(string) bool {
	return func(input string) bool {

		if input == "" {
			return false
		}

		return true
	}
}
