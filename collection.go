package main

// using reflect package for manage array data
import (
	"encoding/json"
	"reflect"
)

// create a type of array with array data
type Collection struct {
	Data interface{}
}

// make a new instance from Arr type
func NewCollection(data interface{}) *Collection {
	return &Collection{data}
}

// check valid data
func (a *Collection) IsValid() bool {
	switch Data := reflect.ValueOf(a.Data).Kind(); Data {
	case reflect.Map, reflect.Slice, reflect.Array, reflect.Struct:
		return true
	default:
		return false
	}
}

// get a index from map array data
func (a *Collection) Get(key interface{}) interface{} {
	if a.IsValid() {
		switch arrayData := reflect.ValueOf(a.Data); arrayData.Kind() {
		case reflect.Map:
			if value := arrayData.MapIndex(reflect.ValueOf(key)); value.IsValid() {
				return value.Interface()
			}
		case reflect.Slice:
			return reflect.ValueOf(a.Data).Index(key.(int)).Interface()
		case reflect.Struct:
			return reflect.ValueOf(a.Data).FieldByName(key.(string)).Interface()
		}
	}
	return nil
}

// get index from slice array data
func (a *Collection) Index(key int) interface{} {
	if a.IsValid() {
		arrayData := reflect.ValueOf(a.Data)
		if value := arrayData.Index(key); value.IsValid() {
			return value.String()
		}
	}
	return nil
}

// check exists index of array data
func (a *Collection) Exists(key string) bool {
	if a.IsValid() {

		switch dataType := reflect.ValueOf(a.Data); dataType.Kind() {
		case reflect.Map:
			if a.Get(key) != nil {
				return true
			}
		case reflect.Slice:
			return a.Has(key)
		}
	}
	return false
}

// get first index of array data
func (a *Collection) First() interface{} {
	if a.IsValid() {
		switch arrayData := reflect.ValueOf(a.Data); arrayData.Kind() {
		case reflect.Map:
			return a.Get(arrayData.MapKeys()[0].Interface())
		case reflect.Slice:
			return a.Index(0)
		}
	}
	return nil
}

// get last index from array data
func (a *Collection) Last() interface{} {
	if a.IsValid() {
		switch arrayData := reflect.ValueOf(a.Data); arrayData.Kind() {
		case reflect.Map:
			lastKey := len(arrayData.MapKeys()) - 1
			return a.Get(arrayData.MapKeys()[lastKey].Interface())
		case reflect.Slice:
			return a.Index(arrayData.Len() - 1)
		}
	}
	return nil
}

// get only keys need in one array in return
func (a *Collection) Only(values ...interface{}) interface{} {
	if a.IsValid() {
		dataValues := map[interface{}]interface{}{}
		for _, key := range values {
			if value := a.Get(key); value != nil {
				dataValues[key] = value
			}
		}
		return dataValues
	}
	return nil
}

// add a map into array data
func (a *Collection) AddMap(key interface{}, value interface{}) *Collection {
	arrayData := reflect.ValueOf(a.Data)
	arrayData.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	a.Data = arrayData.Interface()
	return a
}

// check array has key in slice and map array data
func (a *Collection) Has(key string) bool {
	if a.IsValid() {
		switch arrayData := reflect.ValueOf(a.Data); arrayData.Kind() {
		case reflect.Slice:
			for i := 0; i < arrayData.Len(); i++ {
				if arrayData.Index(i).String() == key {
					return true
				} else if arrayData.Index(i).Kind() == reflect.Interface {
					if arrayData.Index(i).Interface() == key {
						return true
					}
				}
			}
		case reflect.Map, reflect.Struct:
			if a.Get(key) != nil {
				return true
			}
		}
	}
	return false
}

// Convert the given data to json string
func (a *Collection) ToJson() string {
	res, _ := json.Marshal(a.Data)
	return string(res)
}

// Convert the given data to json string bytes
func (a *Collection) ToBytes() []byte {
	return []byte(a.ToJson())
}
