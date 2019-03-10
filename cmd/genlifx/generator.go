package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/codemodus/kace"
)

// Generator is used to accumulate code lines and will write to a file
type Generator struct {
	Name string
	Kind string

	buf  bytes.Buffer
	tmpl *template.Template
}

// NewGenerator returns a Generator initialized with a header
func NewGenerator(name, kind, tmpl string) *Generator {
	g := Generator{Name: name, Kind: kind}

	funcMap := template.FuncMap{
		"pascal":     kace.Pascal,
		"formatchar": formatChar,
	}
	var err error
	g.tmpl, err = template.New("").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		log.Fatalf("Unable to create %s template: %s", name, err)
	}

	g.Printf("// Code generated by \"genlifx\"; DO NOT EDIT.\n")
	g.Printf("package lan\n")
	g.Printf("\n")

	return &g
}

// Printf adds lines to the buffer
func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.buf, format, args...)
}

// WriteFile will output a file for this generator
func (g Generator) WriteFile(data interface{}) error {
	g.tmpl.Execute(&g.buf, data)

	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		// Attempt to write unformatted.go to examine for debugging
		ioutil.WriteFile("unformatted.go", g.buf.Bytes(), 0644)
		return fmt.Errorf("unable to format code: %s", err)
	}

	filename := fmt.Sprintf("%s.go", g.Name)
	if g.Kind != "" {
		filename = fmt.Sprintf("%s_%s.go", g.Name, g.Kind)
	}

	if err := ioutil.WriteFile(filename, src, 0644); err != nil {
		return fmt.Errorf("unable to write file %s: %s", filename, err)
	}

	log.Printf("Wrote %s", filename)
	return nil
}

var (
	enumRegex  = regexp.MustCompile(`^(\[[1-9][0-9]*\])?<([A-Za-z0-9]+)>$`)
	intRegex   = regexp.MustCompile(`^u?int(8|16|32|64)$`)
	floatRegex = regexp.MustCompile(`^float(32|64)$`)
	bytesRegex = regexp.MustCompile(`^\[[0-9]+\]byte$`)
)

func formatChar(s string) string {
	if s == "bool" {
		return "%t"
	}
	if intRegex.MatchString(s) {
		return "%d"
	}
	if floatRegex.MatchString(s) {
		return "%f"
	}
	if bytesRegex.MatchString(s) {
		return "%s"
	}
	return "%s"
}
