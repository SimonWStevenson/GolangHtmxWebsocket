package website
import (
	//"net/http"
	"html/template"
	"os"
	"strings"
	"path/filepath"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
}
func Main(){
	// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go
	dogs := []Pet{
		{
			Name:   "Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
		},
		{
			Name:   "Bruce Wayne",
			Sex:    "Male",
			Intact: false,
			Age:    "3 years, 8 months",
			Breed:  "Chihuahua",
		},
	}
	
	// generate the template
	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
	}

	absPath, err := filepath.Abs(".")
	if err != nil {panic(err)}
	println("the absPath inside website.go is:" + absPath)

	var tmplPath = absPath + "\\pkg\\website\\home.tmpl"
	var tmplFile = "home.tmpl"
	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplPath)
		if err != nil {panic(err)}
	var f *os.File
	f, err = os.Create(absPath + "\\pkg\\website\\home.html")
		if err != nil {panic(err)}
	err = tmpl.Execute(f, dogs)
		if err != nil {panic(err)}
	err = f.Close()
		if err != nil {panic(err)}
}