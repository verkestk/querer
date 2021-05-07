// Package querer provides a way to take a url.Values (struct representation of a query string)
// and populate a struct. this is all based on struct tags matching querystring keys.
package querer

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

const (
	dateLayout     = "2006-01-02"
	dateTimeLayout = "2006-01-02T15:04:05"
)

type params struct {
	MinDOB   time.Time `query:"min_dob"`
	MaxDOB   time.Time `query:"max_dob"`
	MaxAge   uint      `query:"max_age"`
	LastName string    `query:"last_name"`
}

var (
	valueKindFunctions = map[reflect.Kind]func(strValue string) (*reflect.Value, error){
		reflect.Bool:    getBoolValue,
		reflect.Int:     getIntValue,
		reflect.Uint:    getUintValue,
		reflect.Float64: getFloatValue,
		reflect.String:  getStringValue,

		// TODO: support the following:
		// reflect.Int8:    getIntValue,
		// reflect.Int16:   getIntValue,
		// reflect.Int32:   getIntValue,
		// reflect.Int64:   getIntValue,
		// reflect.Uint8:   getUintValue,
		// reflect.Uint16:  getUintValue,
		// reflect.Uint32:  getUintValue,
		// reflect.Uint64:  getUintValue,
		// reflect.Uintptr: getUintValue,
		// reflect.Float32: getFloatValue,
	}
	valueStructFunctions = map[string]func(strValue string) (*reflect.Value, error){
		"time.Time": getStructTimeValue,
	}
)

// UnmarshalQuery takes the values from a url query string and populates a struct based on `query` tags
func UnmarshalQuery(toStruct interface{}, query url.Values) error {
	// make sure this is a pointer
	toKind := reflect.TypeOf(toStruct).Kind()
	if toKind != reflect.Ptr {
		return fmt.Errorf("query unmarshal - cannot marshal into non pointer: %s", toKind)
	}

	// get the reference of the pointer
	toValue := getReference(reflect.ValueOf(toStruct))
	return unmarshalQueryToValue(toValue, query)
}

// unmarshals the url query string into the value
func unmarshalQueryToValue(value *reflect.Value, query url.Values) error {
	// make sure we've got a pointer to a struct
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("query unmarshal - cannot marshal into a pointer to a non-struct: %s", value.Kind())
	}

	// go through the fields
	for i := 0; i < value.NumField(); i++ {

		// for embedded structs, recurse
		if value.Type().Field(i).Anonymous {
			embedded := getReference(value.Field(i))
			unmarshalQueryToValue(embedded, query)
			continue
		}

		// for other members, look for a query tag and set values
		key := value.Type().Field(i).Tag.Get("query")
		if key != "" {

			// get the string value from the query string
			strValue := query.Get(key)
			if strValue == "" {
				// empty, treat as a zero value
				continue
			}

			// get the value to populate
			valueToSet, err := getValue(strValue, value.Type().Field(i).Type)
			if err != nil {
				return err
			}

			// set the value
			if valueToSet != nil {
				value.Field(i).Set(*valueToSet)
			}
		}
	}

	return nil
}

// walks the pointers until a non pointer is found
func getReference(value reflect.Value) *reflect.Value {
	reference := value

	for reference.IsValid() && reference.Type().Kind() == reflect.Ptr {
		// if the reference is nil, make a new one
		if reference.IsNil() {
			reference.Set(reflect.New(reference.Type().Elem()))
		}
		reference = reference.Elem()
	}

	return &reference
}

// turns a string into a Value based on type
func getValue(strValue string, valueType reflect.Type) (*reflect.Value, error) {
	var (
		valueFunction func(strValue string) (*reflect.Value, error)
		ok            bool
	)

	// get a function for converting the string to a value based on kind/name
	if valueType.Kind() == reflect.Struct {
		name := fmt.Sprintf("%s.%s", valueType.PkgPath(), valueType.Name())
		valueFunction, ok = valueStructFunctions[name]
		if !ok {
			return nil, fmt.Errorf("query unmarshal - invalid field type: %s", name)
		}
	} else {
		valueFunction, ok = valueKindFunctions[valueType.Kind()]
		if !ok {
			return nil, fmt.Errorf("query unmarshal - invalid field type: %s", valueType.Name())
		}
	}

	return valueFunction(strValue)
}

func getBoolValue(strValue string) (*reflect.Value, error) {
	boolValue, err := strconv.ParseBool(strValue)
	if err != nil {
		return nil, err
	}

	value := reflect.ValueOf(boolValue)
	return &value, nil
}

func getIntValue(strValue string) (*reflect.Value, error) {
	intValue, err := strconv.ParseInt(strValue, 10, 64)
	if err != nil {
		return nil, err
	}

	value := reflect.ValueOf(int(intValue))
	return &value, nil
}

func getUintValue(strValue string) (*reflect.Value, error) {
	uintValue, err := strconv.ParseUint(strValue, 10, 64)
	if err != nil {
		return nil, err
	}
	value := reflect.ValueOf(uint(uintValue))

	return &value, nil
}

func getFloatValue(strValue string) (*reflect.Value, error) {
	floatValue, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		return nil, err
	}
	value := reflect.ValueOf(floatValue)

	return &value, nil
}

func getStringValue(strValue string) (*reflect.Value, error) {
	value := reflect.ValueOf(strValue)

	return &value, nil
}

func getStructTimeValue(strValue string) (*reflect.Value, error) {
	timeLayout := ""

	switch len(strValue) {
	case 10:
		timeLayout = dateLayout
	case 19:
		timeLayout = dateTimeLayout
	default:
		// invalid date string
		return nil, fmt.Errorf("query unmarshal - invalid date value: %s", strValue)
	}

	timeValue, err := time.Parse(timeLayout, strValue)
	if err != nil {
		// invalid date string
		return nil, fmt.Errorf("query unmarshal - error parsing date value: %s - %s", strValue, err)
	}

	value := reflect.ValueOf(timeValue)
	return &value, nil
}
