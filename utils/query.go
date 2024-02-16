package utils

// GetSortBy returns the value from the mapping based on the given sort key.
//
// Parameters:
//
//	sort string - The key to look up in the mapping.
//	mapping map[string]string - The mapping of keys to values.
//
// Return:
//
//	string - The value corresponding to the sort key in the mapping.
func GetSortBy(sort string, mapping map[string]string) string {
	if sort == "" {
		return "id"
	}

	res, ok := mapping[sort]
	if !ok {
		return "id"
	}

	return res
}
