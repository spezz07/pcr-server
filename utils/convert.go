package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func StructToJson(obj interface{}) string {
	jsons, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(jsons)
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func IntArrayToString(v []int, splitStr string) string {
	if len(v) == 0 {
		return ""
	}
	vv := make([]string, len(v))
	for i, v := range v {
		vv[i] = strconv.Itoa(v)
	}
	return strings.Join(vv, splitStr)
}
func StringArrayToIntArray(v []string) []int {
	if len(v) == 0 {
		return nil
	}
	vv := make([]int, len(v))
	for i, v := range v {
		vv[i], _ = strconv.Atoi(v)
	}
	return vv
}

func IsExistInStringArray(v string, list []string) (bool, int) {
	if len(list) == 0 {
		return false, 0
	}
	for index, val := range list {
		if val == v {
			return true, index
		}
	}
	return false, 0
}

// 遍历struct并且自动进行赋值
func StructByReflect(data map[string]interface{}, inStructPtr interface{}) {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}
	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if v, ok := data[t.Name]; ok {
			if reflect.ValueOf != nil {
				f.Set(reflect.ValueOf(v))
			}
		} else {
			fmt.Printf(t.Name + " not found\n")
		}
	}
}

func BindValue(configMap map[string]interface{}, result interface{}) error {
	// 被绑定的结构体非指针错误返回
	if reflect.ValueOf(result).Kind() != reflect.Ptr {
		return errors.New("input not point")
	}
	// 被绑定的结构体指针为 null 错误返回
	if reflect.ValueOf(result).IsNil() {
		return errors.New("input is null")
	}
	v := reflect.ValueOf(result).Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("json")
		// map 中没该变量有则跳过
		tag = camel2Case(tag)
		value, ok := configMap[tag]
		if !ok {
			continue
		}
		// 跳过结构体中不可 set 的私有变量
		if !v.Field(i).CanSet() {
			continue
		}
		switch v.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			res, err := strconv.ParseInt(string(value.(int)), 10, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetInt(res)
		case reflect.String:
			v.Field(i).SetString(value.(string))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			res, err := strconv.ParseUint(string(value.(uint)), 10, 64)
			if err != nil {
				return err
			}
			v.Field(i).SetUint(res)
		case reflect.Float32:
			panic("float32-不能转为string")
			//res, err := strconv.ParseFloat(value.(float32), 32)
			//if err != nil {
			//	return err
			//}
			//v.Field(i).SetFloat(res)
		case reflect.Float64:
			//panic("float36-不能转为string")
			//res, err := strconv.ParseFloat(value.(string), 64)
			//if err != nil {
			//	return err
			//}
			//v.Field(i).SetFloat(res)
		case reflect.Slice:
			var strArray []string
			var valArray []reflect.Value
			var valArr reflect.Value
			elemKind := t.Field(i).Type.Elem().Kind()
			elemType := t.Field(i).Type.Elem()
			value = strings.Trim(strings.Trim(value.(string), "["), "]")
			strArray = strings.Split(value.(string), ",")
			switch elemKind {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				for _, e := range strArray {
					ee, err := strconv.ParseInt(e, 10, 64)
					if err != nil {
						return err
					}
					valArray = append(valArray, reflect.ValueOf(ee).Convert(elemType))
				}
			case reflect.String:
				for _, e := range strArray {
					valArray = append(valArray, reflect.ValueOf(strings.Trim(e, "\"")).Convert(elemType))
				}
			}
			valArr = reflect.Append(v.Field(i), valArray...)
			v.Field(i).Set(valArr)
		}
	}
	return nil
}

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	b.WriteString(s)
	return b
}
func camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}
