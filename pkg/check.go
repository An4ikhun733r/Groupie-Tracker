package pkg

func IsValid(id string) bool {
	if id[0] == '0' || id == "" {
		return false
	}

	return true
}

func IsRange(id int) bool {
	if id < 1 || id > len(Artist) {
		return false
	}

	return true
}
