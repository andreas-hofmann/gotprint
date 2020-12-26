package gotprint

import (
	"reflect"
	"sort"
	"strings"
)

type StringMatrix struct {
	strings [][]string
	maxLen  []int
	format  FormatSettings
}

func (sm *StringMatrix) Format() *FormatSettings {
	return &sm.format
}

func ToStringMatrix(s interface{}) (result StringMatrix) {
	return toStringMatrix(s, 0)
}

func (sm *StringMatrix) String() string {
	padLen := strings.Count(sm.format.pad, "") - 1

	sm.updateMaxLen()

	s := ""
	for r := 0; r < sm.rows(); r++ {
		if r > 0 {
			s += "\n"
		}

		for c := 0; c < sm.cols(); c++ {
			if sm.format.fixes.Pre != "" {
				s += sm.format.fixes.Pre
			}

			myStr := sm.strings[r][c]
			s += myStr

			if padLen > 0 {
				for i := strings.Count(myStr, "") - 1; i < sm.maxLen[c]; i += padLen {
					s += sm.format.pad
				}
			}

			if sm.format.fixes.Post != "" {
				s += sm.format.fixes.Post
			}

			if c < sm.cols()-1 {
				s += sm.format.separator
			}
		}
	}

	return s
}

func toStringMatrix(s interface{}, level int) (result StringMatrix) {
	sm, ok := s.(StringMatrix)
	if ok {
		return sm
	}

	v, ok := s.(reflect.Value)

	if !ok {
		v = reflect.ValueOf(s)
	}

	result = newStringMatrix("")

	// Check for supported data structures.
	switch v.Kind() {
	// Unsupported types
	case reflect.Chan:
		return

	// Supported types
	case reflect.Interface:
		// Fetch the actual interface value and retry.
		return toStringMatrix(v.Interface(), level)

	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			r := result.rows()
			if i == 0 {
				r = 0
			}
			result.set(toStringMatrix(v.Index(i), level), r, 0)
		}
		return

	case reflect.Map:
		fallthrough
	case reflect.Struct:
		l := level
		if l > 0 && l >= len(result.format.structfixes) {
			l = len(result.format.structfixes) - 1
		}

		skipFix := true

		if result.format.structfixes[l].Pre != "" {
			result.set(newStringMatrix(result.format.structfixes[l].Pre), 0, 0)
			skipFix = false
		}

		r := result.rows() - 1

		switch v.Kind() {
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				column := result.cols(r)
				if i == 0 && skipFix {
					column = 0
				}
				result.set(toStringMatrix(v.Field(i), level+1), 0, column)
			}

		case reflect.Map:
			i := 0

			var content stringmap

			iter := v.MapRange()
			for iter.Next() {
				key := iter.Key()
				value := iter.Value()

				content = append(content, stringmapentry{key: key.String(), value: value})
			}

			content.Sort()

			for _, value := range content {
				column := result.cols(r)
				if i == 0 && skipFix {
					column = 0
				}

				result.set(toStringMatrix(value.value, level+1), 0, column)
				i += 1
			}
		}

		if result.format.structfixes[l].Post != "" {
			result.set(newStringMatrix(result.format.structfixes[l].Post), 0, result.cols(r))
		}

		return

	// Simple, generic types. Just hand them to the generic string function.
	default:
		return newStringMatrix(genericString(v))
	}
}

func newStringMatrix(s ...string) StringMatrix {
	return StringMatrix{
		strings: [][]string{
			s,
		},
		format: currentformat,
	}
}

func (sm StringMatrix) rows() int {
	return len(sm.strings)
}

func (sm StringMatrix) cols(row ...int) int {
	r := 0

	if len(sm.strings) <= r {
		return 0
	}

	if len(row) > 0 {
		r = row[0]
	}

	return len(sm.strings[r])
}

func (sm *StringMatrix) updateMaxLen() {
	sm.maxLen = append(sm.maxLen, make([]int, sm.cols()-len(sm.maxLen))...)

	for r := 0; r < sm.rows(); r++ {
		for c := 0; c < sm.cols(r); c++ {
			l := strings.Count(sm.strings[r][c], "") - 1
			if sm.maxLen[c] < l {
				sm.maxLen[c] = l
			}
		}
	}
}

func (sm *StringMatrix) grow(row, col int) {
	// Resize the rows that new data can fit in.
	sm.strings = append(sm.strings, make([][]string, row-sm.rows())...)

	// Resize the columns, that they can hold all new data.
	for r := 0; r < sm.rows(); r++ {
		if sm.cols(r) < col {
			sm.strings[r] = append(sm.strings[r], make([]string, col-sm.cols(r))...)
		}
	}
}

func (sm *StringMatrix) set(v StringMatrix, row, col int) {
	maxcol := sm.cols()
	maxrow := sm.rows()

	if col+v.cols() >= maxcol {
		maxcol = col + v.cols()
	}

	if row+v.rows() >= maxrow {
		maxrow = row + v.rows()
	}

	sm.grow(maxrow, maxcol)

	// Insert new data.
	for r := 0; r < v.rows(); r++ {
		for c := 0; c < v.cols(r); c++ {
			sm.strings[row+r][col+c] = v.strings[r][c]
		}
	}
}

type stringmapentry struct {
	key   string
	value interface{}
}

type stringmap []stringmapentry

func (sm stringmap) Less(i, j int) bool {
	return sm[i].key < sm[j].key
}

func (sm stringmap) Len() int {
	return len(sm)
}

func (sm stringmap) Swap(i, j int) {
	tmp := sm[i]
	sm[i] = sm[j]
	sm[j] = tmp
}

func (sm stringmap) Sort() {
	sort.Sort(sm)
}
