package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type Scratch struct {
	things map[string]string
}

func (s *Scratch) Get(name string) string {
	return s.things[name]
}

func (s *Scratch) Set(name string, thing string) string {
	s.things[name] = thing
	return ""
}

var scr = &Scratch{
	things: map[string]string{},
}

func main() {
	jsonfile := os.Args[1]
	raw, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	splitSpecNameRegex := regexp.MustCompile(`^\[(\d+)\]([^[]+) \[(\d+\.\d+)\]([^[]+) \[(\d+\.\d+\.\d+)](.+)$`)
	t, err := template.New(filepath.Base(os.Args[2])).Funcs(template.FuncMap{
		"splitSpecName": func(name string) []string {
			found := splitSpecNameRegex.FindStringSubmatch(name)
			return []string{found[1] + found[2], found[3] + found[4], found[5] + found[6]}
		},
		"scratch": func() *Scratch { return scr },
	}).ParseFiles(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out := os.Stdout
	if len(os.Args) == 4 {
		out, err = os.Create(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer out.Close()
	}

	if err := t.Execute(out, data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
