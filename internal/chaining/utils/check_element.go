package utils

func AllElementExist(slice1, slice2 []string) bool {
	slice1Map := make(map[string]bool)
	for _, element := range slice1 {
		slice1Map[element] = true
	}

	for _, element := range slice2 {
		if !slice1Map[element] {
			return false
		}
	}

	return true
}

func ContainsElement(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}

	return false
}

func OneUniqueElement(slice1, slice2, slice_check []string) (string, bool) {
	sliceInputMap := make(map[string]bool)

	for _, element := range slice1 {
		sliceInputMap[element] = true
	}
	for _, element := range slice2 {
		sliceInputMap[element] = true
	}

	for _, element := range slice_check {
		if !sliceInputMap[element] {
			return element, true
		}
	}

	return "", false
}
