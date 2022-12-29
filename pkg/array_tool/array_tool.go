package array_tool

import "reflect"

// 值是否在数组中
func Inarray[T comparable](arr []T, value T) (bool, int) {
	for i, v := range arr {
		if value == v {
			return true, i
		}
	}
	return false, -1
}

func ArrayKeyExists[T any](arr []T, index int) bool {
	for i, _ := range arr {
		if index == i {
			return true
		}
	}
	return false
}

// 根据索引删除值,返回值为乱序
func UnsetArrayByIndex(arr interface{}, indexs []int) (interface{}, error, bool) {
	index_map := map[int]int{}
	for _, iv := range indexs {
		index_map[iv] = 0
	}
	val, ok := isSlice(arr)
	if !ok {
		return arr, nil, false
	}
	sliceLen := val.Len()
	base := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		base[i] = val.Index(i).Interface()
	}
	j := 0
	for i, v := range base {
		if _, ok := index_map[i]; ok {
			base[j], base[i] = v, base[j]
			j++
		}
	}
	arr = base[j:]
	return arr, nil, true
}

// 判断是否为slcie数据
func isSlice(arg interface{}) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == reflect.Slice {
		ok = true
	}
	return
}

//以a为基准，输出a有，b没有的值。
func StrDiff(a []string, b []string) []string {
	var c []string
	temp := map[string]struct{}{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{} // 空struct 不占内存空间
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}
	return c
}
