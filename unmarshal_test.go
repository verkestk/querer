package querer

import (
	"net/url"
	"testing"
	"time"
)

func Test_Empty(t *testing.T) {
	type paramStruct struct {
	}

	query := url.Values(map[string][]string{})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}
}

func Test_NonPtr(t *testing.T) {
	type paramStruct struct {
	}

	query := url.Values(map[string][]string{})

	params := paramStruct{}
	err := UnmarshalQuery(params, query)
	if err == nil {
		t.Errorf("expected error - got nil")
	}
}

func Test_Bool(t *testing.T) {

	type paramStruct struct {
		Bool1       bool `query:"bool_1"`
		Boolt       bool `query:"bool_t"`
		BoolT       bool `query:"bool_T"`
		BoolTRUE    bool `query:"bool_TRUE"`
		Booltrue    bool `query:"bool_true"`
		BoolTrue    bool `query:"bool_True"`
		Bool0       bool `query:"bool_0"`
		Boolf       bool `query:"bool_f"`
		BoolF       bool `query:"bool_F"`
		BoolFALSE   bool `query:"bool_FALSE"`
		Boolfalse   bool `query:"bool_false"`
		BoolFalse   bool `query:"bool_False"`
		BoolMissing bool `query:"bool_Missing"`
		BoolEmpty   bool `query:"bool_Empty"`
		BoolInvalid bool `query:"bool_Invalid"`
	}

	query := url.Values(map[string][]string{
		"bool_1":     []string{"1"},
		"bool_t":     []string{"t"},
		"bool_T":     []string{"T"},
		"bool_TRUE":  []string{"TRUE"},
		"bool_true":  []string{"true"},
		"bool_True":  []string{"True"},
		"bool_0":     []string{"0"},
		"bool_f":     []string{"f"},
		"bool_F":     []string{"F"},
		"bool_FALSE": []string{"FALSE"},
		"bool_false": []string{"false"},
		"bool_False": []string{"False"},
		"bool_Empty": []string{""},
	})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if !params.Bool1 {
		t.Errorf("unexpected value - expected %v - got %v", true, params.Bool1)
	}
	if !params.Boolt {
		t.Errorf("unexpected value - expected %v - got %v", true, params.Boolt)
	}
	if !params.BoolT {
		t.Errorf("unexpected value - expected %v - got %v", true, params.BoolT)
	}
	if !params.BoolTRUE {
		t.Errorf("unexpected value - expected %v - got %v", true, params.BoolTRUE)
	}
	if !params.Booltrue {
		t.Errorf("unexpected value - expected %v - got %v", true, params.Booltrue)
	}
	if !params.BoolTrue {
		t.Errorf("unexpected value - expected %v - got %v", true, params.BoolTrue)
	}
	if params.Bool0 {
		t.Errorf("unexpected value - expected %v - got %v", false, params.Bool0)
	}
	if params.Boolf {
		t.Errorf("unexpected value - expected %v - got %v", false, params.Boolf)
	}
	if params.BoolF {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolF)
	}
	if params.BoolFALSE {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolFALSE)
	}
	if params.Boolfalse {
		t.Errorf("unexpected value - expected %v - got %v", false, params.Boolfalse)
	}
	if params.BoolFalse {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolFalse)
	}
	if params.BoolMissing {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolMissing)
	}
	if params.BoolInvalid {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolInvalid)
	}
	if params.BoolEmpty {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolEmpty)
	}

	query = url.Values(map[string][]string{
		"bool_Invalid": []string{"yes"},
	})

	params = new(paramStruct)
	err = UnmarshalQuery(params, query)

	if err == nil {
		t.Errorf("expected error - got nil")
	}

	if params.BoolInvalid {
		t.Errorf("unexpected value - expected %v - got %v", false, params.BoolInvalid)
	}
}

func Test_Int(t *testing.T) {
	type paramStruct struct {
		Int0       int `query:"int_0"`
		Int1       int `query:"int_1"`
		IntN1      int `query:"int_n1"`
		IntMissing int `query:"int_Missing"`
		IntEmpty   int `query:"int_Empty"`
		IntInvalid int `query:"int_Invalid"`
	}

	query := url.Values(map[string][]string{
		"int_0":     []string{"0"},
		"int_1":     []string{"1"},
		"int_n1":    []string{"-1"},
		"int_Empty": []string{""},
	})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if params.Int0 != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.Int0)
	}
	if params.Int1 != 1 {
		t.Errorf("unexpected value - expected %v - got %v", 1, params.Int1)
	}
	if params.IntN1 != -1 {
		t.Errorf("unexpected value - expected %v - got %v", -1, params.IntN1)
	}
	if params.IntMissing != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.IntMissing)
	}
	if params.IntEmpty != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.IntEmpty)
	}
	if params.IntInvalid != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.IntInvalid)
	}

	query = url.Values(map[string][]string{
		"int_Invalid": []string{"yes"},
	})

	params = new(paramStruct)
	err = UnmarshalQuery(params, query)

	if err == nil {
		t.Errorf("expected error - got nil")
	}

	if params.IntInvalid != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.IntInvalid)
	}
}

func Test_Uint(t *testing.T) {
	type paramStruct struct {
		UInt0       uint `query:"uint_0"`
		UInt1       uint `query:"uint_1"`
		UIntMissing uint `query:"uint_Missing"`
		UIntEmpty   uint `query:"uint_Empty"`
		UIntInvalid uint `query:"uint_Invalid"`
	}

	query := url.Values(map[string][]string{
		"uint_0":     []string{"0"},
		"uint_1":     []string{"1"},
		"uint_Empty": []string{""},
	})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if params.UInt0 != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.UInt0)
	}
	if params.UInt1 != 1 {
		t.Errorf("unexpected value - expected %v - got %v", 1, params.UInt1)
	}
	if params.UIntMissing != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.UIntMissing)
	}
	if params.UIntEmpty != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.UIntEmpty)
	}
	if params.UIntInvalid != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.UIntInvalid)
	}

	query = url.Values(map[string][]string{
		"uint_Invalid": []string{"-1"},
	})

	params = new(paramStruct)
	err = UnmarshalQuery(params, query)

	if err == nil {
		t.Errorf("expected error - got nil")
	}

	if params.UIntInvalid != 0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.UIntInvalid)
	}
}

func Test_Float64(t *testing.T) {
	type paramStruct struct {
		Float640       int `query:"float64_0"`
		Float641       int `query:"float64_1"`
		Float64N1      int `query:"float64_n1"`
		Float64Missing int `query:"float64_Missing"`
		Float64Empty   int `query:"float64_Empty"`
		Float64Invalid int `query:"float64_Invalid"`
	}

	query := url.Values(map[string][]string{
		"float64_0":     []string{"0"},
		"float64_1":     []string{"1"},
		"float64_n1":    []string{"-1"},
		"float64_Empty": []string{""},
	})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if params.Float640 != 0.0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.Float640)
	}
	if params.Float641 != 1.0 {
		t.Errorf("unexpected value - expected %v - got %v", 1, params.Float641)
	}
	if params.Float64N1 != -1.0 {
		t.Errorf("unexpected value - expected %v - got %v", -1, params.Float64N1)
	}
	if params.Float64Missing != 0.0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.Float64Missing)
	}
	if params.Float64Empty != 0.0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.Float64Empty)
	}
	if params.Float64Invalid != 0.0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.Float64Invalid)
	}

	query = url.Values(map[string][]string{
		"float64_Invalid": []string{"yes"},
	})

	params = new(paramStruct)
	err = UnmarshalQuery(params, query)

	if err == nil {
		t.Errorf("expected error - got nil")
	}

	if params.Float64Invalid != 0.0 {
		t.Errorf("unexpected value - expected %v - got %v", 0, params.Float64Invalid)
	}
}

func Test_Time(t *testing.T) {
	type paramStruct struct {
		TimeDate     time.Time `query:"time_Date"`
		TimeDateTime time.Time `query:"time_Datetime"`
		TimeMissing  time.Time `query:"time_Missing"`
		TimeEmpty    time.Time `query:"time_Empty"`
		TimeInvalid  time.Time `query:"time_Invalid"`
	}

	query := url.Values(map[string][]string{
		"time_Date":     []string{"2014-08-18"},
		"time_Datetime": []string{"2014-08-18T21:30:17"},
		"time_Empty":    []string{""},
	})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if params.TimeDate.Format(dateLayout) != "2014-08-18" {
		t.Errorf("unexpected value - expected %s - got %s", "2014-08-18", params.TimeDate.Format(dateLayout))
	}
	if params.TimeDateTime.Format(dateTimeLayout) != "2014-08-18T21:30:17" {
		t.Errorf("unexpected value - expected %s - got %s", "2014-08-18T21:30:17", params.TimeDateTime.Format(dateTimeLayout))
	}
	if !params.TimeMissing.Equal(time.Time{}) {
		t.Errorf("unexpected value - expected %v - got %v", time.Time{}, params.TimeMissing)
	}
	if !params.TimeEmpty.Equal(time.Time{}) {
		t.Errorf("unexpected value - expected %v - got %v", time.Time{}, params.TimeEmpty)
	}
	if !params.TimeInvalid.Equal(time.Time{}) {
		t.Errorf("unexpected value - expected %v - got %v", time.Time{}, params.TimeInvalid)
	}

	query = url.Values(map[string][]string{
		"time_Invalid": []string{"yes"},
	})

	params = new(paramStruct)
	err = UnmarshalQuery(params, query)

	if err == nil {
		t.Errorf("expected error - got nil")
	}

	if !params.TimeInvalid.Equal(time.Time{}) {
		t.Errorf("unexpected value - expected %v - got %v", time.Time{}, params.TimeInvalid)
	}
}

func Test_Anonmyous(t *testing.T) {

	type timeStruct struct {
		Time time.Time `query:"time"`
	}

	type floatStruct struct {
		*timeStruct
		Float64 float64 `query:"float64"`
	}

	type uintStruct struct {
		*floatStruct
		UInt uint `query:"uint"`
	}

	type intStruct struct {
		*uintStruct
		Int int `query:"int"`
	}

	type paramStruct struct {
		intStruct
		Bool bool `query:"bool"`
	}

	query := url.Values(map[string][]string{
		"bool":    []string{"1"},
		"int":     []string{"1"},
		"uint":    []string{"1"},
		"float64": []string{"1"},
		"time":    []string{"2014-08-18"},
	})

	params := new(paramStruct)
	err := UnmarshalQuery(params, query)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if !params.Bool {
		t.Errorf("unexpected value - expected %v - got %v", true, params.Bool)
	}
	if params.Int != 1 {
		t.Errorf("unexpected value - expected %v - got %v", 1, params.Int)
	}
	if params.UInt != 1 {
		t.Errorf("unexpected value - expected %v - got %v", 1, params.UInt)
	}
	if params.Float64 != 1.0 {
		t.Errorf("unexpected value - expected %v - got %v", 1, params.Float64)
	}
	if params.Time.Format(dateLayout) != "2014-08-18" {
		t.Errorf("unexpected value - expected %s - got %s", "2014-08-18", params.Time.Format(dateLayout))
	}
}
