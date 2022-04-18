package db

// Init all db resources
func Init() {
}

func selectWhere(equalCods map[string]string, greaterCods map[string]string, lessCods map[string]string, itemList []*Item) ([]*Item, int64, error) {
	for _, item := range shopItemList {
		if checkEqualItem(equalCods, item) && checkGreaterItem(greaterCods, item) && checkLessItem(lessCods, item) {
			itemList = append(itemList, item)
		}
	}
	return itemList, int64(len(itemList)), nil
}
