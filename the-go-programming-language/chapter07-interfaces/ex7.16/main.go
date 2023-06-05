// Write a web-based calculator program.
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var templ *template.Template

type TemplateBody struct {
	Expr   string
	Result float64
}

func main() {
	var err error
	templ, err = setTemplate("template.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/calculate", CalculateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	exprStr := r.URL.Query().Get("expr")
	envStr := r.URL.Query().Get("env")
	expr, err := Parse(exprStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad expression: %s", err)
		return
	}
	env, err := parseEnv(envStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad assignment: %s", err)
		return
	}
	result := expr.Eval(env)
	templ.Execute(w, TemplateBody{exprStr, result})
}

func parseEnv(envStr string) (Env, error) {
	env := Env{}
	assignments := strings.Fields(envStr)
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			return nil, fmt.Errorf("bad assignment: %s\n", a)
		}
		if len(fields) == 1 {
			fields = append(fields, "0")
		}
		ident, valStr := fields[0], fields[1]
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("bad value for %s, using zero: %s\n", ident, err)
		}
		env[Var(ident)] = val
	}
	return env, nil
}

func setTemplate(filename string) (*template.Template, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	templ, err := template.New(filename).Parse(buf.String())
	if err != nil {
		return nil, err
	}
	return templ, nil
}
