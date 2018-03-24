package main

import (
    "flag"
    "time"
    "text/template"
    "strings"
    "bufio"
    "os"
)


var template_string =
`/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*
*           Author:  {{.AuthorName}}
*           Project: {{.ProjectName}}
*           (c) {{.ProjectYear}}
*
*
*           {{.WittyText}}
*
*
*           File created on {{.CreatedOn}}
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * /
`


// Context that gets passed to the template
type Context struct {
    ProjectName string
    AuthorName  string
    ProjectYear int
    WittyText   string
    CreatedOn   string
}


func main() {
    projectName := flag.String(
        "project",
        "Super Secret Project",
        "Name of the project you are working on.")
    authorName  := flag.String(
        "author",
        "Johnny English",
        "Name of the person who created the file.")
    projectYear := flag.Int(
        "year",
        time.Now().Year(),
        "Calendar yearear relevant to the project.")
    wittyText   := flag.String(
        "witty-text",
        "Simple, but not simpler. - Albert Einstein",
        "Witty text :)")

    flag.Parse()

    con := &Context{
        ProjectName: *projectName,
        AuthorName:  *authorName,
        ProjectYear: *projectYear,
        WittyText:   *wittyText,
        CreatedOn:   time.Now().Format("Mon 2.1.2006 at 15:4"),
    }

    filenames := flag.Args()
    for _, filename := range filenames {
        if ! strings.HasSuffix(strings.ToLower(filename), ".go") {
            filename = filename + ".go"
        }
        file, _ := os.Create(filename)
        w := bufio.NewWriter(file) //use "os.Stdout" instead of "file" for quick debugging
        t, _ := template.New("templ").Parse(template_string)
        t.ExecuteTemplate(w, "templ", con)
        w.Flush()
        file.Close()
    }
}
