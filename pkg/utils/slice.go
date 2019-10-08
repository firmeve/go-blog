package utils

func SliceStringUnique(items []string) []string {
	keyItems := make(map[string]bool)
	newItems := make([]string, 0)
	for _, item := range items {
		if _, ok := keyItems[item]; !ok {
			newItems = append(newItems, item)
			keyItems[item] = true
		}
	}

	return newItems
}

func SliceStringExcept(items, excepts []string) []string {
	newSlice := make([]string, 0)
	for _, item := range items {
		if !SliceStringIn(excepts, item) {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}

//func SliceAnyUnique(items []interface{}) []interface{} {
//	keyItems := make(map[interface{}]bool)
//	newItems := make([]interface{}, 0)
//	for _, item := range items {
//		if _, ok := keyItems[item]; !ok {
//			newItems = append(newItems, item)
//			keyItems[item] = true
//		}
//	}
//
//	return newItems
//}

func SliceStringIn(items []string, value string) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

func SliceIntIn(items []int, value int) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

func SliceUintIn(items []uint, value uint) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}
