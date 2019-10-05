package MapAndPackage

type Dictionary map[string]string

//字典类型
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

//判断单词是否存在字典中
func UnkownWord(dictionary map[string]string, word string) bool {
	_, ok := dictionary[word]
	if !ok {
		return false
	}
	return true
}

//map 是引用类型不需要使用指针就可以直接修改
func (d Dictionary) AddWord(word, definition string) {
	d[word] = definition
}

//删除指定的键制
func (d Dictionary) DeleteWord(word string) {
	delete(d, word)
}
