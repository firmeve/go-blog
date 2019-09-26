package utils

func SliceStringUnqiue(items []string) []string {
	return items//SliceAnyUnique(items).([]string)
}

func SliceAnyUnique(items []interface{}) []interface{} {
	keyItems := make(map[interface{}]bool)
	newItems := make([]interface{}, 0)
	for _, item := range items {
		if _, ok := keyItems[item]; !ok {
			newItems = append(newItems, item)
			keyItems[item] = true
		}
	}

	return newItems
}

func InSlice(items []interface{},value interface{}) bool {
	for _,item := range items {
		if item == value{
			return true
		}
	}
	return false
}
