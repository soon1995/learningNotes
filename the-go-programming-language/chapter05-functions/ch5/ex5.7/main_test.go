package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// But html.Parse parses pretty much anything, so this test is useless.
func TestPrettyOutputCanBeParsed(t *testing.T) {
	input := `<html>
<body>
<p class="something" id="short" ><span class="special">hi</span></p><br/>
</body>
</html>
`
	expect := `<html>
  <head/>
  <body>
    <p class='something' id='short'>
      <span class='special'>
        hi
      </span>
    </p>
    <br/>
  </body>
</html>
`
	out := bytes.NewBuffer(nil)
	p := NewPrettier(out)

	p.Pretty(strings.NewReader(input))
	fmt.Println(out.String())
	fmt.Println(expect)

	if out.String() != expect {
		t.Fatal("not equal")
	}
}
