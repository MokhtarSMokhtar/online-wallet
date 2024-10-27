package utile

func ContainsValue[T comparable](slice []T, key T) bool {
	for _, k := range slice {
		if k == key {
			return true
		}
	}
	return false
}
