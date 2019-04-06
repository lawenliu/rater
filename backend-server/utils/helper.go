package utils

func CheckDtype(dtype string) bool {
	if dtype == DtypeAvatar || dtype == DtypeImage || dtype == DtypeVideo {
		return true
	}

	return false
}