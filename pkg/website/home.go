package website
import (
	//"net/http"
	"html/template"
	"os"
	"strings"
	"path/filepath"
)
func home(){
	
	
	// generate the template
	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
	}

	absPath, err := filepath.Abs(".")
	if err != nil {panic(err)}
	println("the absPath inside website.go is:" + absPath)

	var tmplPath2 = "C:\\Users\\simon\\GitHub\\Four\\pkg\\website\\home.tmpl"
	var tmplFile2 = "home.tmpl"
	tmpl, err := template.New(tmplFile2).Funcs(funcMap).ParseFiles(tmplPath2)
		if err != nil {panic(err)}
	var f *os.File
	f, err = os.Create("C:\\Users\\simon\\GitHub\\Four\\pkg\\website\\home.html")
		if err != nil {panic(err)}
	err = tmpl.Execute(f, os.DevNull)
		if err != nil {panic(err)}
	err = f.Close()
		if err != nil {panic(err)}	
}