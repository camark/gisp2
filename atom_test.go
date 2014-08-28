package gisp

import (
	p "github.com/Dwarfartisan/goparsec"
	"reflect"
	"testing"
)

func TestAtomParse0(t *testing.T) {
	data := "x"
	state := p.MemoryParseState(data)
	a, err := AtomParser(state)
	if err == nil {
		test := Atom{"x", ANY}
		if !reflect.DeepEqual(test, a) {
			t.Fatalf("except Atom{\"x\", ATOM} but %v", a)
		}
	} else {
		t.Fatalf("except Atom{\"x\", ATOM} but %v", err)
	}
}

func TestAtomParse1(t *testing.T) {
	data := "x::atom"
	state := p.MemoryParseState(data)
	a, err := AtomParser(state)
	if err == nil {
		test := Atom{"x", ATOM}
		if !reflect.DeepEqual(test, a) {
			t.Fatalf("except Atom{\"x\", ATOM} but %v", a)
		}
	} else {
		t.Fatalf("except Atom{\"x\", ATOM} but %v", err)
	}
}

func TestAtomParse2(t *testing.T) {
	data := "x::any"
	state := p.MemoryParseState(data)
	a, err := AtomParser(state)
	if err == nil {
		test := Atom{"x", ANY}
		if !reflect.DeepEqual(test, a) {
			t.Fatalf("except Atom{\"x\", ANY} but %v", a)
		}
	} else {
		t.Fatalf("except Atom{\"x\", ANY} but %v", err)
	}
}

func TestAtomParse3(t *testing.T) {
	data := "x::int"
	state := p.MemoryParseState(data)
	a, err := AtomParser(state)
	if err == nil {
		test := Atom{"x", INT}
		if !reflect.DeepEqual(test, a) {
			t.Fatalf("except Atom{\"x\", INT} but %v", a)
		}
	} else {
		t.Fatalf("except Atom{\"x\", INT} but %v", err)
	}
}
