package defaults

import (
	"reflect"
	"strconv"
	"time"
)

type StructDef struct {
	DefaultTagName string // struct default tag
}

// FormatStruct Set default value for struct tag, Support type of uint, *uint, int, *int, string, *string, float,
// *float, *bool. Nonsupport nest struct.
// function input params type of pointer.
func (s *StructDef) FormatStruct(pointer interface{}) {
	objT := reflect.TypeOf(pointer)
	objV := reflect.ValueOf(pointer)

	// Judge struct type is *struct
	if objT.Kind() == reflect.Ptr && objT.Elem().Kind() == reflect.Struct {

		for i := 0; i < objT.Elem().NumField(); i++ {

			// Current field value
			currentField := objV.Elem().Field(i).Interface()

			// If value is pointer
			if objV.Elem().Field(i).Kind() == reflect.Ptr {
				if !objV.Elem().Field(i).IsNil() {
					currentField = objV.Field(i).Elem().Interface()
				} else {
					currentField = nil
				}
			}

			// Determines whether the field has a value
			if !IsSatisfied(currentField) {
				// Get field tag default value, if value is empty then pass.
				if defaultValue := objT.Elem().Field(i).Tag.Get(s.DefaultTagName); defaultValue != "" {
					// Field type is pointer
					if objT.Elem().Field(i).Type.Kind() == reflect.Ptr {
						// Type of pointer
						switch objT.Elem().Field(i).Type.Elem().Kind() {
						case reflect.Bool:
							if v, err := strconv.ParseBool(defaultValue); err == nil {
								objV.Elem().Field(i).Set(reflect.ValueOf(&v))
							}
						case reflect.Uint8:
							if v, err := strconv.ParseUint(defaultValue, 10, 8); err == nil {
								tmp := uint8(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Uint16:
							if v, err := strconv.ParseUint(defaultValue, 10, 16); err == nil {
								tmp := uint16(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Uint32:
							if v, err := strconv.ParseUint(defaultValue, 10, 32); err == nil {
								tmp := uint32(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Uint64:
							if v, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
								objV.Elem().Field(i).Set(reflect.ValueOf(&v))
							}
						case reflect.Uint:
							if v, err := strconv.ParseUint(defaultValue, 10, 32<<(^uint(0)>>63)); err == nil {
								tmp := uint(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Int8:
							if v, err := strconv.ParseInt(defaultValue, 10, 8); err == nil {
								tmp := int8(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Int16:
							if v, err := strconv.ParseInt(defaultValue, 10, 16); err == nil {
								tmp := int16(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Int32:
							if v, err := strconv.ParseInt(defaultValue, 10, 32); err == nil {
								tmp := int32(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Int64:
							if v, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
								objV.Elem().Field(i).Set(reflect.ValueOf(&v))
							}
						case reflect.Int:
							if v, err := strconv.ParseInt(defaultValue, 10, 32<<(^uint(0)>>63)); err == nil {
								tmp := int(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Float32:
							if v, err := strconv.ParseFloat(defaultValue, 32); err == nil {
								tmp := float32(v)
								objV.Elem().Field(i).Set(reflect.ValueOf(&tmp))
							}
						case reflect.Float64:
							if v, err := strconv.ParseFloat(defaultValue, 64); err == nil {
								objV.Elem().Field(i).Set(reflect.ValueOf(&v))
							}
						case reflect.String:
							objV.Elem().Field(i).Set(reflect.ValueOf(&defaultValue))
						}
					} else {
						// Field type is not pointer
						switch objT.Elem().Field(i).Type.Kind() {
						case reflect.Uint8:
							if v, err := strconv.ParseUint(defaultValue, 10, 8); err == nil {
								objV.Elem().Field(i).SetUint(v)
							}
						case reflect.Uint16:
							if v, err := strconv.ParseUint(defaultValue, 10, 16); err == nil {
								objV.Elem().Field(i).SetUint(v)
							}
						case reflect.Uint32:
							if v, err := strconv.ParseUint(defaultValue, 10, 32); err == nil {
								objV.Elem().Field(i).SetUint(v)
							}
						case reflect.Uint64:
							if v, err := strconv.ParseUint(defaultValue, 10, 64); err == nil {
								objV.Elem().Field(i).SetUint(v)
							}
						case reflect.Uint:
							if v, err := strconv.ParseUint(defaultValue, 10, 32<<(^uint(0)>>63)); err == nil {
								objV.Elem().Field(i).SetUint(v)
							}
						case reflect.Int8:
							if v, err := strconv.ParseInt(defaultValue, 10, 8); err == nil {
								objV.Elem().Field(i).SetInt(v)
							}
						case reflect.Int16:
							if v, err := strconv.ParseInt(defaultValue, 10, 16); err == nil {
								objV.Elem().Field(i).SetInt(v)
							}
						case reflect.Int32:
							if v, err := strconv.ParseInt(defaultValue, 10, 32); err == nil {
								objV.Elem().Field(i).SetInt(v)
							}
						case reflect.Int64:
							if v, err := strconv.ParseInt(defaultValue, 10, 64); err == nil {
								objV.Elem().Field(i).SetInt(v)
							}
						case reflect.Int:
							if v, err := strconv.ParseInt(defaultValue, 10, 32<<(^uint(0)>>63)); err == nil {
								objV.Elem().Field(i).SetInt(v)
							}
						case reflect.Float32:
							if v, err := strconv.ParseFloat(defaultValue, 32); err == nil {
								objV.Elem().Field(i).SetFloat(v)
							}
						case reflect.Float64:
							if v, err := strconv.ParseFloat(defaultValue, 64); err == nil {
								objV.Elem().Field(i).SetFloat(v)
							}
						case reflect.String:
							objV.Elem().Field(i).SetString(defaultValue)
						}
					}
				}
			}
		}
	}
}

// IsSatisfied judge object type is satisfied.
func IsSatisfied(obj interface{}) bool {
	if obj == nil {
		return false
	}
	if i, ok := obj.(int); ok {
		return i != 0
	}
	if i, ok := obj.(uint); ok {
		return i != 0
	}
	if i, ok := obj.(int8); ok {
		return i != 0
	}
	if i, ok := obj.(uint8); ok {
		return i != 0
	}
	if i, ok := obj.(int16); ok {
		return i != 0
	}
	if i, ok := obj.(uint16); ok {
		return i != 0
	}
	if i, ok := obj.(uint32); ok {
		return i != 0
	}
	if i, ok := obj.(int32); ok {
		return i != 0
	}
	if i, ok := obj.(int64); ok {
		return i != 0
	}
	if i, ok := obj.(uint64); ok {
		return i != 0
	}
	if i, ok := obj.(float32); ok {
		return i != 0
	}
	if i, ok := obj.(float64); ok {
		return i != 0
	}
	if t, ok := obj.(time.Time); ok {
		return !t.IsZero()
	}
	if v := reflect.ValueOf(obj); v.Kind() == reflect.Slice {
		return v.Len() > 0
	}
	if i, ok := obj.(string); ok {
		return i != ""
	}
	return true
}
