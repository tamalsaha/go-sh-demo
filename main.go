package main

import "github.com/codeskyblue/go-sh"

func main() {
	sh.Command("env", map[string]string{
		"demo": "yes",
	}).Command("grep", "demo").Run()
}