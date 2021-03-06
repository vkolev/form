package form

import (
	"errors"
	"net/url"
	"testing"
	"time"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//
//
// go test -cpuprofile cpu.out
// ./validator.test -test.bench=. -test.cpuprofile=cpu.prof
// go tool pprof validator.test cpu.prof
//
//
// go test -memprofile mem.out

func TestInt(t *testing.T) {

	type TestInt struct {
		Int              int
		Int8             int8
		Int16            int16
		Int32            int32
		Int64            int64
		IntPtr           *int
		Int8Ptr          *int8
		Int16Ptr         *int16
		Int32Ptr         *int32
		Int64Ptr         *int64
		IntArray         []int
		IntPtrArray      []*int
		IntArrayArray    [][]int
		IntPtrArrayArray [][]*int
		IntMap           map[int]int
		IntPtrMap        map[*int]*int
		NoURLValue       int
	}

	values := url.Values{
		"Int":                    []string{"3"},
		"Int8":                   []string{"3"},
		"Int16":                  []string{"3"},
		"Int32":                  []string{"3"},
		"Int64":                  []string{"3"},
		"IntPtr":                 []string{"3"},
		"Int8Ptr":                []string{"3"},
		"Int16Ptr":               []string{"3"},
		"Int32Ptr":               []string{"3"},
		"Int64Ptr":               []string{"3"},
		"IntArray":               []string{"1", "2", "3"},
		"IntPtrArray[0]":         []string{"1"},
		"IntPtrArray[2]":         []string{"3"},
		"IntArrayArray[0][0]":    []string{"1"},
		"IntArrayArray[0][2]":    []string{"3"},
		"IntArrayArray[2][0]":    []string{"1"},
		"IntPtrArrayArray[0][0]": []string{"1"},
		"IntPtrArrayArray[0][2]": []string{"3"},
		"IntPtrArrayArray[2][0]": []string{"1"},
		"IntMap[1]":              []string{"3"},
		"IntPtrMap[1]":           []string{"3"},
	}

	var test TestInt

	test.IntArray = make([]int, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Int, int(3))
	Equal(t, test.Int8, int8(3))
	Equal(t, test.Int16, int16(3))
	Equal(t, test.Int32, int32(3))
	Equal(t, test.Int64, int64(3))

	Equal(t, *test.IntPtr, int(3))
	Equal(t, *test.Int8Ptr, int8(3))
	Equal(t, *test.Int16Ptr, int16(3))
	Equal(t, *test.Int32Ptr, int32(3))
	Equal(t, *test.Int64Ptr, int64(3))

	Equal(t, len(test.IntArray), 4)
	Equal(t, test.IntArray[0], int(1))
	Equal(t, test.IntArray[1], int(2))
	Equal(t, test.IntArray[2], int(3))
	Equal(t, test.IntArray[3], int(0))

	Equal(t, len(test.IntPtrArray), 3)
	Equal(t, *test.IntPtrArray[0], int(1))
	Equal(t, test.IntPtrArray[1], nil)
	Equal(t, *test.IntPtrArray[2], int(3))

	Equal(t, len(test.IntArrayArray), 3)
	Equal(t, len(test.IntArrayArray[0]), 3)
	Equal(t, len(test.IntArrayArray[1]), 0)
	Equal(t, len(test.IntArrayArray[2]), 1)
	Equal(t, test.IntArrayArray[0][0], int(1))
	Equal(t, test.IntArrayArray[0][1], int(0))
	Equal(t, test.IntArrayArray[0][2], int(3))
	Equal(t, test.IntArrayArray[2][0], int(1))

	Equal(t, len(test.IntPtrArrayArray), 3)
	Equal(t, len(test.IntPtrArrayArray[0]), 3)
	Equal(t, len(test.IntPtrArrayArray[1]), 0)
	Equal(t, len(test.IntPtrArrayArray[2]), 1)
	Equal(t, *test.IntPtrArrayArray[0][0], int(1))
	Equal(t, test.IntPtrArrayArray[0][1], nil)
	Equal(t, *test.IntPtrArrayArray[0][2], int(3))
	Equal(t, *test.IntPtrArrayArray[2][0], int(1))

	Equal(t, len(test.IntMap), 1)
	Equal(t, len(test.IntPtrMap), 1)

	v, ok := test.IntMap[1]
	Equal(t, ok, true)
	Equal(t, v, int(3))

	Equal(t, test.NoURLValue, int(0))
}

func TestUint(t *testing.T) {

	type TestUint struct {
		Uint              uint
		Uint8             uint8
		Uint16            uint16
		Uint32            uint32
		Uint64            uint64
		UintPtr           *uint
		Uint8Ptr          *uint8
		Uint16Ptr         *uint16
		Uint32Ptr         *uint32
		Uint64Ptr         *uint64
		UintArray         []uint
		UintPtrArray      []*uint
		UintArrayArray    [][]uint
		UintPtrArrayArray [][]*uint
		UintMap           map[uint]uint
		UintPtrMap        map[*uint]*uint
		NoURLValue        uint
	}

	values := url.Values{
		"Uint":                    []string{"3"},
		"Uint8":                   []string{"3"},
		"Uint16":                  []string{"3"},
		"Uint32":                  []string{"3"},
		"Uint64":                  []string{"3"},
		"UintPtr":                 []string{"3"},
		"Uint8Ptr":                []string{"3"},
		"Uint16Ptr":               []string{"3"},
		"Uint32Ptr":               []string{"3"},
		"Uint64Ptr":               []string{"3"},
		"UintArray":               []string{"1", "2", "3"},
		"UintPtrArray[0]":         []string{"1"},
		"UintPtrArray[2]":         []string{"3"},
		"UintArrayArray[0][0]":    []string{"1"},
		"UintArrayArray[0][2]":    []string{"3"},
		"UintArrayArray[2][0]":    []string{"1"},
		"UintPtrArrayArray[0][0]": []string{"1"},
		"UintPtrArrayArray[0][2]": []string{"3"},
		"UintPtrArrayArray[2][0]": []string{"1"},
		"UintMap[1]":              []string{"3"},
		"UintPtrMap[1]":           []string{"3"},
	}

	var test TestUint

	test.UintArray = make([]uint, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Uint, uint(3))
	Equal(t, test.Uint8, uint8(3))
	Equal(t, test.Uint16, uint16(3))
	Equal(t, test.Uint32, uint32(3))
	Equal(t, test.Uint64, uint64(3))

	Equal(t, *test.UintPtr, uint(3))
	Equal(t, *test.Uint8Ptr, uint8(3))
	Equal(t, *test.Uint16Ptr, uint16(3))
	Equal(t, *test.Uint32Ptr, uint32(3))
	Equal(t, *test.Uint64Ptr, uint64(3))

	Equal(t, len(test.UintArray), 4)
	Equal(t, test.UintArray[0], uint(1))
	Equal(t, test.UintArray[1], uint(2))
	Equal(t, test.UintArray[2], uint(3))
	Equal(t, test.UintArray[3], uint(0))

	Equal(t, len(test.UintPtrArray), 3)
	Equal(t, *test.UintPtrArray[0], uint(1))
	Equal(t, test.UintPtrArray[1], nil)
	Equal(t, *test.UintPtrArray[2], uint(3))

	Equal(t, len(test.UintArrayArray), 3)
	Equal(t, len(test.UintArrayArray[0]), 3)
	Equal(t, len(test.UintArrayArray[1]), 0)
	Equal(t, len(test.UintArrayArray[2]), 1)
	Equal(t, test.UintArrayArray[0][0], uint(1))
	Equal(t, test.UintArrayArray[0][1], uint(0))
	Equal(t, test.UintArrayArray[0][2], uint(3))
	Equal(t, test.UintArrayArray[2][0], uint(1))

	Equal(t, len(test.UintPtrArrayArray), 3)
	Equal(t, len(test.UintPtrArrayArray[0]), 3)
	Equal(t, len(test.UintPtrArrayArray[1]), 0)
	Equal(t, len(test.UintPtrArrayArray[2]), 1)
	Equal(t, *test.UintPtrArrayArray[0][0], uint(1))
	Equal(t, test.UintPtrArrayArray[0][1], nil)
	Equal(t, *test.UintPtrArrayArray[0][2], uint(3))
	Equal(t, *test.UintPtrArrayArray[2][0], uint(1))

	Equal(t, len(test.UintMap), 1)
	Equal(t, len(test.UintPtrMap), 1)

	v, ok := test.UintMap[1]
	Equal(t, ok, true)
	Equal(t, v, uint(3))

	Equal(t, test.NoURLValue, uint(0))
}

func TestString(t *testing.T) {

	type TestString struct {
		String              string
		StringPtr           *string
		StringArray         []string
		StringPtrArray      []*string
		StringArrayArray    [][]string
		StringPtrArrayArray [][]*string
		StringMap           map[string]string
		StringPtrMap        map[*string]*string
		NoURLValue          string
	}

	values := url.Values{
		"String":                    []string{"3"},
		"StringPtr":                 []string{"3"},
		"StringArray":               []string{"1", "2", "3"},
		"StringPtrArray[0]":         []string{"1"},
		"StringPtrArray[2]":         []string{"3"},
		"StringArrayArray[0][0]":    []string{"1"},
		"StringArrayArray[0][2]":    []string{"3"},
		"StringArrayArray[2][0]":    []string{"1"},
		"StringPtrArrayArray[0][0]": []string{"1"},
		"StringPtrArrayArray[0][2]": []string{"3"},
		"StringPtrArrayArray[2][0]": []string{"1"},
		"StringMap[1]":              []string{"3"},
		"StringPtrMap[1]":           []string{"3"},
	}

	var test TestString

	test.StringArray = make([]string, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.String, "3")

	Equal(t, *test.StringPtr, "3")

	Equal(t, len(test.StringArray), 4)
	Equal(t, test.StringArray[0], "1")
	Equal(t, test.StringArray[1], "2")
	Equal(t, test.StringArray[2], "3")
	Equal(t, test.StringArray[3], "")

	Equal(t, len(test.StringPtrArray), 3)
	Equal(t, *test.StringPtrArray[0], "1")
	Equal(t, test.StringPtrArray[1], nil)
	Equal(t, *test.StringPtrArray[2], "3")

	Equal(t, len(test.StringArrayArray), 3)
	Equal(t, len(test.StringArrayArray[0]), 3)
	Equal(t, len(test.StringArrayArray[1]), 0)
	Equal(t, len(test.StringArrayArray[2]), 1)
	Equal(t, test.StringArrayArray[0][0], "1")
	Equal(t, test.StringArrayArray[0][1], "")
	Equal(t, test.StringArrayArray[0][2], "3")
	Equal(t, test.StringArrayArray[2][0], "1")

	Equal(t, len(test.StringPtrArrayArray), 3)
	Equal(t, len(test.StringPtrArrayArray[0]), 3)
	Equal(t, len(test.StringPtrArrayArray[1]), 0)
	Equal(t, len(test.StringPtrArrayArray[2]), 1)
	Equal(t, *test.StringPtrArrayArray[0][0], "1")
	Equal(t, test.StringPtrArrayArray[0][1], nil)
	Equal(t, *test.StringPtrArrayArray[0][2], "3")
	Equal(t, *test.StringPtrArrayArray[2][0], "1")

	Equal(t, len(test.StringMap), 1)
	Equal(t, len(test.StringPtrMap), 1)

	v, ok := test.StringMap["1"]
	Equal(t, ok, true)
	Equal(t, v, "3")

	Equal(t, test.NoURLValue, "")
}

func TestFloat(t *testing.T) {

	type TestFloat struct {
		Float32              float32
		Float32Ptr           *float32
		Float64              float64
		Float64Ptr           *float64
		Float32Array         []float32
		Float32PtrArray      []*float32
		Float32ArrayArray    [][]float32
		Float32PtrArrayArray [][]*float32
		Float32Map           map[float32]float32
		Float32PtrMap        map[*float32]*float32
		NoURLValue           float32
	}

	values := url.Values{
		"Float32":                    []string{"3.3"},
		"Float32Ptr":                 []string{"3.3"},
		"Float64":                    []string{"3.3"},
		"Float64Ptr":                 []string{"3.3"},
		"Float32Array":               []string{"1.1", "2.2", "3.3"},
		"Float32PtrArray[0]":         []string{"1.1"},
		"Float32PtrArray[2]":         []string{"3.3"},
		"Float32ArrayArray[0][0]":    []string{"1.1"},
		"Float32ArrayArray[0][2]":    []string{"3.3"},
		"Float32ArrayArray[2][0]":    []string{"1.1"},
		"Float32PtrArrayArray[0][0]": []string{"1.1"},
		"Float32PtrArrayArray[0][2]": []string{"3.3"},
		"Float32PtrArrayArray[2][0]": []string{"1.1"},
		"Float32Map[1.1]":            []string{"3.3"},
		"Float32PtrMap[1.1]":         []string{"3.3"},
	}

	var test TestFloat

	test.Float32Array = make([]float32, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Float32, float32(3.3))
	Equal(t, test.Float64, float64(3.3))

	Equal(t, *test.Float32Ptr, float32(3.3))
	Equal(t, *test.Float64Ptr, float64(3.3))

	Equal(t, len(test.Float32Array), 4)
	Equal(t, test.Float32Array[0], float32(1.1))
	Equal(t, test.Float32Array[1], float32(2.2))
	Equal(t, test.Float32Array[2], float32(3.3))
	Equal(t, test.Float32Array[3], float32(0.0))

	Equal(t, len(test.Float32PtrArray), 3)
	Equal(t, *test.Float32PtrArray[0], float32(1.1))
	Equal(t, test.Float32PtrArray[1], nil)
	Equal(t, *test.Float32PtrArray[2], float32(3.3))

	Equal(t, len(test.Float32ArrayArray), 3)
	Equal(t, len(test.Float32ArrayArray[0]), 3)
	Equal(t, len(test.Float32ArrayArray[1]), 0)
	Equal(t, len(test.Float32ArrayArray[2]), 1)
	Equal(t, test.Float32ArrayArray[0][0], float32(1.1))
	Equal(t, test.Float32ArrayArray[0][1], float32(0.0))
	Equal(t, test.Float32ArrayArray[0][2], float32(3.3))
	Equal(t, test.Float32ArrayArray[2][0], float32(1.1))

	Equal(t, len(test.Float32PtrArrayArray), 3)
	Equal(t, len(test.Float32PtrArrayArray[0]), 3)
	Equal(t, len(test.Float32PtrArrayArray[1]), 0)
	Equal(t, len(test.Float32PtrArrayArray[2]), 1)
	Equal(t, *test.Float32PtrArrayArray[0][0], float32(1.1))
	Equal(t, test.Float32PtrArrayArray[0][1], nil)
	Equal(t, *test.Float32PtrArrayArray[0][2], float32(3.3))
	Equal(t, *test.Float32PtrArrayArray[2][0], float32(1.1))

	Equal(t, len(test.Float32Map), 1)
	Equal(t, len(test.Float32PtrMap), 1)

	v, ok := test.Float32Map[float32(1.1)]
	Equal(t, ok, true)
	Equal(t, v, float32(3.3))

	Equal(t, test.NoURLValue, float32(0.0))
}

func TestBool(t *testing.T) {

	type TestBool struct {
		Bool              bool
		BoolPtr           *bool
		BoolArray         []bool
		BoolPtrArray      []*bool
		BoolArrayArray    [][]bool
		BoolPtrArrayArray [][]*bool
		BoolMap           map[bool]bool
		BoolPtrMap        map[*bool]*bool
		NoURLValue        bool
	}

	values := url.Values{
		"Bool":                    []string{"true"},
		"BoolPtr":                 []string{"true"},
		"BoolArray":               []string{"off", "t", "on"},
		"BoolPtrArray[0]":         []string{"true"},
		"BoolPtrArray[2]":         []string{"T"},
		"BoolArrayArray[0][0]":    []string{"TRUE"},
		"BoolArrayArray[0][2]":    []string{"True"},
		"BoolArrayArray[2][0]":    []string{"true"},
		"BoolPtrArrayArray[0][0]": []string{"true"},
		"BoolPtrArrayArray[0][2]": []string{"t"},
		"BoolPtrArrayArray[2][0]": []string{"1"},
		"BoolMap[true]":           []string{"true"},
		"BoolPtrMap[t]":           []string{"true"},
	}

	var test TestBool

	test.BoolArray = make([]bool, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Bool, true)

	Equal(t, *test.BoolPtr, true)

	Equal(t, len(test.BoolArray), 4)
	Equal(t, test.BoolArray[0], false)
	Equal(t, test.BoolArray[1], true)
	Equal(t, test.BoolArray[2], true)
	Equal(t, test.BoolArray[3], false)

	Equal(t, len(test.BoolPtrArray), 3)
	Equal(t, *test.BoolPtrArray[0], true)
	Equal(t, test.BoolPtrArray[1], nil)
	Equal(t, *test.BoolPtrArray[2], true)

	Equal(t, len(test.BoolArrayArray), 3)
	Equal(t, len(test.BoolArrayArray[0]), 3)
	Equal(t, len(test.BoolArrayArray[1]), 0)
	Equal(t, len(test.BoolArrayArray[2]), 1)
	Equal(t, test.BoolArrayArray[0][0], true)
	Equal(t, test.BoolArrayArray[0][1], false)
	Equal(t, test.BoolArrayArray[0][2], true)
	Equal(t, test.BoolArrayArray[2][0], true)

	Equal(t, len(test.BoolPtrArrayArray), 3)
	Equal(t, len(test.BoolPtrArrayArray[0]), 3)
	Equal(t, len(test.BoolPtrArrayArray[1]), 0)
	Equal(t, len(test.BoolPtrArrayArray[2]), 1)
	Equal(t, *test.BoolPtrArrayArray[0][0], true)
	Equal(t, test.BoolPtrArrayArray[0][1], nil)
	Equal(t, *test.BoolPtrArrayArray[0][2], true)
	Equal(t, *test.BoolPtrArrayArray[2][0], true)

	Equal(t, len(test.BoolMap), 1)
	Equal(t, len(test.BoolPtrMap), 1)

	v, ok := test.BoolMap[true]
	Equal(t, ok, true)
	Equal(t, v, true)

	Equal(t, test.NoURLValue, false)
}

func TestStruct(t *testing.T) {

	type Phone struct {
		Number string
	}

	type TestStruct struct {
		Name      string `form:"name"`
		Phone     []Phone
		PhonePtr  []*Phone
		NonNilPtr *Phone
		Ignore    string `form:"-"`
		Anonymous struct {
			Value     string
			Ignore    string `form:"-"`
			unexposed string
		}
		Time                       time.Time
		TimePtr                    *time.Time
		unexposed                  string
		Invalid                    interface{}
		ExistingMap                map[string]string `form:"mp"`
		MapNoValue                 map[int]int
		NilArray                   []string
		TooSmallArray              []string
		TooSmallCapOKArray         []string
		ZeroLengthArray            []string
		TooSmallNumberedArray      []string
		TooSmallCapOKNumberedArray []string
		BigEnoughNumberedArray     []string
		IfaceNonNil                interface{}
		IfaceInvalid               interface{}
	}

	values := url.Values{
		"name":                          []string{"joeybloggs"},
		"Ignore":                        []string{"ignore"},
		"Phone[0].Number":               []string{"1(111)111-1111"},
		"Phone[1].Number":               []string{"9(999)999-9999"},
		"PhonePtr[0].Number":            []string{"1(111)111-1111"},
		"PhonePtr[1].Number":            []string{"9(999)999-9999"},
		"NonNilPtr.Number":              []string{"9(999)999-9999"},
		"Anonymous.Value":               []string{"Anon"},
		"Time":                          []string{"2016-01-02"},
		"TimePtr":                       []string{"2016-01-02"},
		"mp[key]":                       []string{"value"},
		"NilArray":                      []string{"1", "2"},
		"TooSmallArray":                 []string{"1", "2"},
		"TooSmallCapOKArray":            []string{"1", "2"},
		"ZeroLengthArray":               []string{},
		"TooSmallNumberedArray[2]":      []string{"2"},
		"TooSmallCapOKNumberedArray[2]": []string{"2"},
		"BigEnoughNumberedArray[2]":     []string{"1"},
	}

	var test TestStruct
	test.ExistingMap = map[string]string{"existingkey": "existingvalue"}
	test.NonNilPtr = new(Phone)
	test.IfaceNonNil = new(Phone)
	test.IfaceInvalid = nil
	test.TooSmallArray = []string{"0"}
	test.TooSmallCapOKArray = make([]string, 0, 10)
	test.TooSmallNumberedArray = []string{"0"}
	test.TooSmallCapOKNumberedArray = make([]string, 0, 10)
	test.BigEnoughNumberedArray = make([]string, 3, 10)

	decoder := NewDecoder()
	decoder.SetTagName("form")
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
	}, time.Time{})

	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Name, "joeybloggs")
	Equal(t, test.Ignore, "")
	Equal(t, len(test.Phone), 2)
	Equal(t, test.Phone[0].Number, "1(111)111-1111")
	Equal(t, test.Phone[1].Number, "9(999)999-9999")
	Equal(t, len(test.PhonePtr), 2)
	Equal(t, (*test.PhonePtr[0]).Number, "1(111)111-1111")
	Equal(t, (*test.PhonePtr[1]).Number, "9(999)999-9999")
	Equal(t, test.NonNilPtr.Number, "9(999)999-9999")
	Equal(t, test.Anonymous.Value, "Anon")
	Equal(t, len(test.ExistingMap), 2)
	Equal(t, test.ExistingMap["existingkey"], "existingvalue")
	Equal(t, test.ExistingMap["key"], "value")
	Equal(t, len(test.NilArray), 2)
	Equal(t, test.NilArray[0], "1")
	Equal(t, test.NilArray[1], "2")
	Equal(t, len(test.TooSmallArray), 2)
	Equal(t, test.TooSmallArray[0], "1")
	Equal(t, test.TooSmallArray[1], "2")
	Equal(t, len(test.ZeroLengthArray), 0)
	Equal(t, len(test.TooSmallNumberedArray), 3)
	Equal(t, test.TooSmallNumberedArray[0], "0")
	Equal(t, test.TooSmallNumberedArray[1], "")
	Equal(t, test.TooSmallNumberedArray[2], "2")
	Equal(t, len(test.BigEnoughNumberedArray), 3)
	Equal(t, cap(test.BigEnoughNumberedArray), 10)
	Equal(t, test.BigEnoughNumberedArray[0], "")
	Equal(t, test.BigEnoughNumberedArray[1], "")
	Equal(t, test.BigEnoughNumberedArray[2], "1")
	Equal(t, len(test.TooSmallCapOKArray), 2)
	Equal(t, cap(test.TooSmallCapOKArray), 10)
	Equal(t, test.TooSmallCapOKArray[0], "1")
	Equal(t, test.TooSmallCapOKArray[1], "2")
	Equal(t, len(test.TooSmallCapOKNumberedArray), 3)
	Equal(t, cap(test.TooSmallCapOKNumberedArray), 10)
	Equal(t, test.TooSmallCapOKNumberedArray[0], "")
	Equal(t, test.TooSmallCapOKNumberedArray[1], "")
	Equal(t, test.TooSmallCapOKNumberedArray[2], "2")

	tm, _ := time.Parse("2006-01-02", "2016-01-02")
	Equal(t, test.Time.Equal(tm), true)
	Equal(t, (*test.TimePtr).Equal(tm), true)

	s := struct {
		Value     string
		Ignore    string `form:"-"`
		unexposed string
	}{}

	errs = decoder.Decode(&s, values)
	Equal(t, errs, nil)
	Equal(t, s.Value, "")
	Equal(t, s.Ignore, "")
	Equal(t, s.unexposed, "")
}

func TestNativeTime(t *testing.T) {

	type TestError struct {
		Time        time.Time
		TimeNoValue time.Time
	}

	values := url.Values{
		"Time":        []string{"2006-01-02T15:04:05Z"},
		"TimeNoValue": []string{""},
	}

	var test TestError

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	Equal(t, test.Time.Equal(tm), true)
	Equal(t, test.TimeNoValue.Equal(tm), false)
}

func TestErrors(t *testing.T) {

	type TestError struct {
		Bool           bool `form:"bool"`
		Int            int
		Uint           uint
		Float32        float32
		String         string
		Time           time.Time
		MapBadIntKey   map[int]int
		MapBadUintKey  map[uint]uint
		MapBadFloatKey map[float32]float32
		MapBadBoolKey  map[bool]bool
		MapBadKeyType  map[complex64]int
		BadArrayValue  []int
	}

	values := url.Values{
		"bool":                  []string{"uh-huh"},
		"Int":                   []string{"bad"},
		"Uint":                  []string{"bad"},
		"Float32":               []string{"bad"},
		"String":                []string{"str bad return val"},
		"Time":                  []string{"bad"},
		"MapBadIntKey[key]":     []string{"1"},
		"MapBadUintKey[key]":    []string{"1"},
		"MapBadFloatKey[key]":   []string{"1.1"},
		"MapBadBoolKey[uh-huh]": []string{"true"},
		"MapBadKeyType[1.4]":    []string{"5"},
		"BadArrayValue[0]":      []string{"badintval"},
	}

	var test TestError

	decoder := NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return nil, errors.New("Bad Type Conversion")
	}, "")
	errs := decoder.Decode(&test, values)
	NotEqual(t, errs, nil)

	e := errs.Error()
	NotEqual(t, e, "")

	err := errs.(DecodeErrors)
	k := err["bool"]
	Equal(t, k.Error(), "Invalid Boolean Value 'uh-huh' Type 'bool' Namespace 'bool'")

	k = err["Int"]
	Equal(t, k.Error(), "Invalid Integer Value 'bad' Type 'int' Namespace 'Int'")

	k = err["Uint"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'bad' Type 'uint' Namespace 'Uint'")

	k = err["Float32"]
	Equal(t, k.Error(), "Invalid Float Value 'bad' Type 'float32' Namespace 'Float32'")

	k = err["String"]
	Equal(t, k.Error(), "Bad Type Conversion")

	k = err["Time"]
	Equal(t, k.Error(), "parsing time \"bad\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"bad\" as \"2006\"")

	k = err["MapBadIntKey"]
	Equal(t, k.Error(), "Invalid Integer Value 'key' Type 'int' Namespace 'MapBadIntKey'")

	k = err["MapBadUintKey"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'key' Type 'uint' Namespace 'MapBadUintKey'")

	k = err["MapBadFloatKey"]
	Equal(t, k.Error(), "Invalid Float Value 'key' Type 'float32' Namespace 'MapBadFloatKey'")

	k = err["MapBadBoolKey"]
	Equal(t, k.Error(), "Invalid Boolean Value 'uh-huh' Type 'bool' Namespace 'MapBadBoolKey'")

	k = err["MapBadKeyType"]
	Equal(t, k.Error(), "Unsupported Map Key '1.4', Type 'complex64' Namespace 'MapBadKeyType'")

	k = err["BadArrayValue[0]"]
	Equal(t, k.Error(), "Invalid Integer Value 'badintval' Type 'int' Namespace 'BadArrayValue[0]'")
}

func TestPanics(t *testing.T) {

	type Phone struct {
		Number string
	}

	type TestError struct {
		Phone []Phone
	}

	values := url.Values{
		"Phone[0.Number": []string{"1(111)111-1111"},
	}

	var test TestError

	decoder := NewDecoder()

	PanicMatches(t, func() { decoder.Decode(&test, values) }, "Invalid formatting for key 'Phone[0.Number' missing bracket")

	i := 1
	PanicMatches(t, func() { decoder.Decode(&i, values) }, "interface must be a pointer to a struct")
}

func TestMapKeys(t *testing.T) {

	type TestMapKeys struct {
		MapIfaceKey  map[interface{}]string
		MapFloatKey  map[float32]float32
		MapNestedInt map[int]map[int]int
	}

	values := url.Values{
		"MapIfaceKey[key]":   []string{"3"},
		"MapFloatKey[0.0]":   []string{"3.3"},
		"MapNestedInt[1][2]": []string{"3"},
	}

	var test TestMapKeys

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)
	Equal(t, test.MapIfaceKey["key"], "3")
	Equal(t, test.MapFloatKey[float32(0.0)], float32(3.3))

	Equal(t, len(test.MapNestedInt), 1)
	Equal(t, len(test.MapNestedInt[1]), 1)
	Equal(t, test.MapNestedInt[1][2], 3)
}

func TestStructRecursion(t *testing.T) {

	type Nested struct {
		Value  string
		Nested *Nested
	}

	type TestRecursive struct {
		Nested Nested
	}

	values := url.Values{
		"Value":        []string{"value"},
		"Nested.Value": []string{"value"},
	}

	var test TestRecursive

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)
	Equal(t, test.Nested.Value, "value")
	Equal(t, test.Nested.Nested, nil)
}

func TestFormDecode(t *testing.T) {

	type Struct2 struct {
		Foo string
		Bar string
	}

	type Struct2Wrapper struct {
		InnerSlice []Struct2
	}

	sliceValues := map[string][]string{
		"InnerSlice[0].Foo": {"foo-is-set"},
	}

	singleValues := map[string][]string{
		"Foo": {"foo-is-set"},
	}

	fd := NewDecoder()

	dst := Struct2Wrapper{}
	err := fd.Decode(&dst, sliceValues)
	Equal(t, err, nil)
	NotEqual(t, dst.InnerSlice, nil)
	Equal(t, dst.InnerSlice[0].Foo, "foo-is-set")

	dst2 := Struct2{}
	err = fd.Decode(&dst2, singleValues)
	Equal(t, err, nil)
	Equal(t, dst2.Foo, "foo-is-set")
}
