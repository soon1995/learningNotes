// mostly copied from torbiak/gopl/ex12.7
package main

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Decoder struct {
	lex *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	s := scanner.Scanner{}
	s.Init(r)
	return &Decoder{
		lex: &lexer{
			scan: s,
		},
	}
}

func (d *Decoder) Decode(out interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Pos(), x)
		}
	}()
	if d.lex.token == 0 {
		d.lex.next()
	}
	d.read(reflect.ValueOf(out).Elem())
	return nil
}

func (d *Decoder) read(v reflect.Value) {
	switch d.lex.token {
	case scanner.Ident:
		switch d.lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			d.lex.next()
			return
		case "true":
			v.SetBool(true)
			d.lex.next()
			return
		case "false":
			v.SetBool(false)
			d.lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(d.lex.text())
		v.SetString(s)
		d.lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(d.lex.text())
		v.SetInt(int64(i))
		d.lex.next()
		return
	case '(':
		d.lex.next()
		d.readList(v)
		d.lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q, %v, at %s", d.lex.text(), scanner.TokenString(d.lex.token), d.lex.scan.Pos()))
}

func (d *Decoder) readList(v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !d.isEnd(); i++ {
			d.read(v.Index(i))
		}

	case reflect.Slice:
		for !d.isEnd() {
			item := reflect.New(v.Type().Elem()).Elem()
			d.read(item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct:
		for !d.isEnd() {
			d.lex.consume('(')
			if d.lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", d.lex.text()))
			}
			name := d.lex.text()
			d.lex.next()
			d.read(v.FieldByName(name))
			d.lex.consume(')')
		}

	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !d.isEnd() {
			d.lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			d.read(key)
			value := reflect.New(v.Type().Elem()).Elem()
			d.read(value)
			v.SetMapIndex(key, value)
			d.lex.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func (d *Decoder) isEnd() bool {
	switch d.lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q at pos %s", lex.text(), want, lex.scan.Pos()))
	}
	lex.next()
}
