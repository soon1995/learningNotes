# Exploring `fmt` package

Q: Which expressions below have the best performance when printing an error?

```go
fmt.Sprintf("%v", err) // 1
fmt.Sprintf("%v", err.Error()) // 2
fmt.Sprintf("%s", err) // 3
fmt.Sprintf("%s", err.Error()) // 4
```

The answer is the 4th expression. Let's delve into the source code:

First, the args types will be determined:

```go
func (p *pp) printArg(arg any, verb rune) {
// ...
	switch f := arg.(type) {
	case bool:
		p.fmtBool(f, verb)
	case float32:
		p.fmtFloat(float64(f), 32, verb)
	case float64:
		p.fmtFloat(f, 64, verb)
	case complex64:
		p.fmtComplex(complex128(f), 64, verb)
	case complex128:
		p.fmtComplex(f, 128, verb)
	case int:
		p.fmtInteger(uint64(f), signed, verb)
	case int8:
		p.fmtInteger(uint64(f), signed, verb)
	case int16:
		p.fmtInteger(uint64(f), signed, verb)
	case int32:
		p.fmtInteger(uint64(f), signed, verb)
	case int64:
		p.fmtInteger(uint64(f), signed, verb)
	case uint:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint8:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint16:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint32:
		p.fmtInteger(uint64(f), unsigned, verb)
	case uint64:
		p.fmtInteger(f, unsigned, verb)
	case uintptr:
		p.fmtInteger(uint64(f), unsigned, verb)
	case string:
		p.fmtString(f, verb)
	case []byte:
		p.fmtBytes(f, verb, "[]byte")
	case reflect.Value:
		// Handle extractable values with special methods
		// since printValue does not handle them at depth 0.
		if f.IsValid() && f.CanInterface() {
			p.arg = f.Interface()
			if p.handleMethods(verb) {
				return
			}
		}
		p.printValue(f, verb, 0)
	default:
		// If the type is not simple, it might have methods.
		if !p.handleMethods(verb) {
			// Need to use reflection, since the type had no
			// interface methods that could be used for formatting.
			p.printValue(reflect.ValueOf(f), verb, 0)
		}
	}
    // ...
}
```

If we pass `err.Error()` (a string) as argument, then the `p.fmtString(f, verb)` method
is called, this method do lesser compared to `p.handleMethod(verb)`

Below are the detailed implementation of these mentioned methods:

```go
func (p *pp) fmtString(v string, verb rune) {
	switch verb {
	case 'v':
		if p.fmt.sharpV {
			p.fmt.fmtQ(v)
		} else {
			p.fmt.fmtS(v)
		}
	case 's':
		p.fmt.fmtS(v)
	case 'x':
		p.fmt.fmtSx(v, ldigits)
	case 'X':
		p.fmt.fmtSx(v, udigits)
	case 'q':
		p.fmt.fmtQ(v)
	default:
		p.badVerb(verb)
	}
}

func (p *pp) handleMethods(verb rune) (handled bool) {
    // ... Skipped some operations

    // If a string is acceptable according to the format, see if
    // the value satisfies one of the string-valued interfaces.
    // Println etc. set verb to %v, which is "stringable".
    switch verb {
        case 'v', 's', 'x', 'X', 'q':
            // Is it an error or Stringer?
            // The duplication in the bodies is necessary:
            // setting handled and deferring catchPanic
            // must happen before calling the method.
            switch v := p.arg.(type) {
                case error:
                    handled = true
                        defer p.catchPanic(p.arg, verb, "Error")
                        p.fmtString(v.Error(), verb)
                        return

                case Stringer:
                        handled = true
                            defer p.catchPanic(p.arg, verb, "String")
                            p.fmtString(v.String(), verb)
                            return
            }
    }
    // ... Skipped some operations
}
```

In conclusion, below listed the expressions from highest to lowest performance:

```go
fmt.Sprintf("%s", err.Error())
fmt.Sprintf("%v", err.Error())
fmt.Sprintf("%s", err)
fmt.Sprintf("%v", err)
```
