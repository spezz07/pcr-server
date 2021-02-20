package utils

func RemoveItemArray(listPt *[]string, index int) {
	list := *listPt
	list[len(list)-1], list[index] = list[index], list[len(list)-1]
	*listPt = list[:len(list)-1]
}
