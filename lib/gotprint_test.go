package gotprint

import (
	"fmt"
	"testing"
)

type teststruct struct {
	FirstString      string
	SecondString     string
	FirstFloatString float64
	FirstInt         int
}

type nestedstruct struct {
	firstfloat   float64
	nestedstruct teststruct
	teststring   string
}

type nestedneststruct struct {
	firstfloat   float64
	nestedstruct nestedstruct
	teststring   string
}

type nestslicestruct struct {
	firstfloat   float64
	nestedstruct []teststruct
	teststring   string
}

func TestMatrixSet(t *testing.T) {
	sm := newStringMatrix("")
	sm.strings = [][]string{{"00-1", "01-1"}, {"10-1", "11-1"}}

	sm2 := newStringMatrix("")
	sm2.strings = [][]string{{"00-2", "01-2"}, {"10-2", "11-2"}}

	sm3 := newStringMatrix("")
	sm3.strings = [][]string{{"00-3", "01-3"}, {"10-3", "11-3"}}

	sm4 := newStringMatrix("")
	sm4.strings = [][]string{{"00-4", "01-4"}, {"10-4", "11-4"}}

	sm.set(sm2, 2, 2)
	sm.set(sm3, 1, 1)
	sm.set(sm4, 2, 0)

	s := sm.String()

	expected := `00-1 01-1          
10-1 00-3 01-3     
00-4 01-4 11-3 01-2
10-4 11-4 10-2 11-2`

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}

	fmt.Printf("Result:\n%s\n", s)
}

func TestTprintNestedStruct(t *testing.T) {
	expected := "5.1 [ asdf foobarXY 1000.456 5 ] foo"

	s := Sprint(nestedstruct{5.1, teststruct{"asdf", "foobarXY", 1000.456, 5}, "foo"})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintNestedNestStructSlice(t *testing.T) {
	expected := "100.5 [ 5.1 ( asdf foobarXY 1000.456 5 ) foo ] lastEntry"

	s := Sprint(nestedneststruct{100.5, nestedstruct{5.1, teststruct{"asdf", "foobarXY", 1000.456, 5}, "foo"}, "lastEntry"})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintFormattedNestedNestStructSlice(t *testing.T) {
	defer SetDefaultFormat()

	expected := "- 100.5 { 5.1 [ asdf foobarXY 1000.456 5 ] foo } lastEntry -"
	Format().SetStructFixes([]Fix{{"-", "-"}, {"{", "}"}, {"[", "]"}})
	s := Sprint(nestedneststruct{100.5, nestedstruct{5.1, teststruct{"asdf", "foobarXY", 1000.456, 5}, "foo"}, "lastEntry"})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}

	expected = "( 100.5 ( 5.1 ( asdf foobarXY 1000.456 5 ) foo ) lastEntry )"
	Format().SetStructFixes([]Fix{{"(", ")"}})
	s = Sprint(nestedneststruct{100.5, nestedstruct{5.1, teststruct{"asdf", "foobarXY", 1000.456, 5}, "foo"}, "lastEntry"})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintNestedStructSlice(t *testing.T) {
	expected := `5.1   [ asdf foobarXY 1000.456    5 ] foo
100.5 [ X    Y        10.44343456 3 ] bar`

	s := Sprint([]nestedstruct{
		{5.1, teststruct{"asdf", "foobarXY", 1000.456, 5}, "foo"},
		{100.50, teststruct{"X", "Y", 10.44343456, 3}, "bar"},
	})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintStructSlice(t *testing.T) {
	expected := `asdf foobarXY 1000.456    5
X    Y        10.44343456 3`

	s := Sprint([]teststruct{
		{"asdf", "foobarXY", 1000.456, 5},
		{"X", "Y", 10.44343456, 3},
	})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintSimpleSlice(t *testing.T) {
	expected := "asdf\nasdf\ntest"
	s := Sprint([]string{"asdf", "asdf", "test"})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintStruct(t *testing.T) {
	expected := "asdf foobar 30000.12 5"
	s := Sprint(teststruct{"asdf", "foobar", 30000.12, 5})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}

func TestTprintNestedSlices(t *testing.T) {
	expected := `123.45 [ s01 s02 0.5 2   ] I'm a string!
       [ s11 s12 0.5 2   ]              
       [ s21 s22 1.5 800 ]              `
	s := Sprint(nestslicestruct{123.45, []teststruct{
		{"s01", "s02", 0.5, 2}, {"s11", "s12", 0.5, 2}, {"s21", "s22", 1.5, 800},
	}, "I'm a string!"})

	fmt.Printf("Result:\n%s\n", s)

	if s == "" {
		t.Error("Struct error")
	} else if s != expected {
		t.Error("Output mismatch.")
		fmt.Printf("Expected:\n%s\n", expected)
	}
}
