package utils

func SliceWithUniqueValues(slice []string) []string {
	uniqueValues := make([]string, 0)
	encountered := make(map[string]bool)

	for _, v := range slice {
		if !encountered[v] {
			encountered[v] = true
			uniqueValues = append(uniqueValues, v)
		}
	}

	return uniqueValues
}
