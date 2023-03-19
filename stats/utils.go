package stats

func isInSlice(value string, slice []string) bool {
	for _, a := range slice {
		if a == value {
			return true
		}
	}
	return false
}
