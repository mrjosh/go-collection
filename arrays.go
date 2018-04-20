package arrays

// using reflect package for manage array data
import "reflect"

// create a type of array with array data
type Arr struct {
	Data interface{}
}

// make a new instance from Arr type
func Array(data interface{}) Arr {
	return Arr{data}
}

// get data from array type
func (a Arr) GetData() interface{} {
	return a.Data
}

// check valid data from data
func (a Arr) isValid() bool {
	switch Data := reflect.ValueOf(a.GetData()).Kind(); Data {
	case reflect.Map, reflect.Slice, reflect.Array:
		return true
	default:
		return false
	}
}

// get a index from map array data
func (a Arr) Get(key interface{}) interface{} {
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

// get index from slice array data
func (a Arr) Index(key int) interface{} {
	if a.isValid() {
		arrayData := reflect.ValueOf(a.GetData())
		if value := arrayData.Index(key); value.IsValid() {
			return value.String()
		}
	}
	return nil
}

// check exists index of array data
func (a Arr) Exists(key string) bool {
	if a.isValid() {
		if a.Get(key) != nil {
			return true
		}
	}
	return false
}

// get first index of array data
func (a Arr) First() interface{} {
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

// get last index from array data
func (a Arr) Last() interface{} {
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

// get only keys need in one array in return
func (a Arr) Only(values ...string) interface{} {
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

// add a map into array data
func (a Arr) AddMap(key interface{}, value interface{}) Arr {
	arrayData := reflect.ValueOf(a.GetData())
	arrayData.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	a.Data = arrayData.Interface()
	return a
}

// check array has key in slice and map array data
func (a Arr) Has(key string) bool {
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
