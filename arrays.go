package arrays

import "reflect"

type Array struct {
	Data interface{}
}

func Arr(data interface{}) Array {
	return Array{data}
}

func (a Array) GetData() interface{} {
	return a.Data
}

func (a Array) isValid() bool {
	switch Data := reflect.ValueOf(a.GetData()).Kind(); Data {
	case reflect.Map, reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

func (a Array) Get(key interface{}) interface{} {
	if a.isValid() {
		switch arrayData := reflect.ValueOf(a.GetData()); arrayData.Kind() {
		case reflect.Map:
			if value := arrayData.MapIndex(reflect.ValueOf(key)); value.IsValid() {
				return value.String()
			}
		}
	}
	return nil
}

func (a Array) Index(key int) interface{} {
	if a.isValid() {
		arrayData := reflect.ValueOf(a.GetData())
		if value := arrayData.Index(key); value.IsValid() {
			return value.String()
		}
	}
	return nil
}

func (a Array) Exists(key string) bool {
	if a.isValid() {
		if a.Get(key) != nil {
			return true
		}
	}
	return false
}

func (a Array) First() interface{} {
	if a.isValid() {
		switch arrayData := reflect.ValueOf(a.GetData()); arrayData.Kind() {
		case reflect.Map:
			return a.Get(arrayData.MapKeys()[0].Interface())
		case reflect.Slice:
			return a.Index(0)
		}
	}
	return nil
}

func (a Array) Last() interface{} {
	if a.isValid() {
		switch arrayData := reflect.ValueOf(a.GetData()); arrayData.Kind() {
		case reflect.Map:
			lastKey := len(arrayData.MapKeys())
			return a.Get(arrayData.MapKeys()[lastKey-1].Interface())
		case reflect.Slice:
			return a.Index(arrayData.Len() - 1)
		}
	}
	return nil
}

func (a Array) Only(values ...string) interface{} {
	if a.isValid() {
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

func (a Array) AddMap(key interface{}, value interface{}) Array {
	arrayData := reflect.ValueOf(a.GetData())
	arrayData.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	a.Data = arrayData.Interface()
	return a
}

func (a Array) Has(key string) bool {
	if a.isValid() {
		switch arrayData := reflect.ValueOf(a.GetData()); arrayData.Kind() {
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
		case reflect.Map:
			for _, index := range arrayData.MapKeys() {
				if index.String() == key {
					return true
				} else if index.Kind() == reflect.Interface {
					if index.Interface() == key {
						return true
					}
				}
			}
		}
	}
	return false
}