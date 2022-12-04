package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// content type represents the HTML content tto add into the template
type content struct {
	Title string
	Body  template.HTML
}

const (
	header = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
		<title>Markdown Preview Tool</title>
	</head>
<body>
			`
	footer = `
	</body>
</html>
`
	defaultTemplate = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">
		<title>{{ .Title }}</title>
	</head>
	<body>
{{ .Body }}
	</body>
</html>
`
)

func main() {
	// parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	// if user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string, tFname string, out io.Writer, skipPreview bool) error {
	// read all the data from the input file and check for errors
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	htmlData, err := parseContent(input, tFname)
	if err != nil {
		return err
	}

	// create temporary file and check for errors
	tmp, err := ioutil.TempFile("", "mdp.html")
	if err != nil {
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	outName := tmp.Name()
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}

	defer os.Remove(outName)

	return preview(outName)
}

func parseContent(input []byte, tFname string) ([]byte, error) {
	// parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	// parse the content of the defaultTemplate const into a new Template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// if user provided alternate template file, replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}

	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(body),
	}

	// create a buffer of bytes to writeto file
	var buffer bytes.Buffer

	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func saveHTML(filename string, data []byte) error {
	// write the bytes to the file
	return ioutil.WriteFile(filename, data, 0644)
}

func preview(fname string) error {
	cName := ""
	cParams := []string{}

	// define executable based on OS
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}

	// append filename to parameters slice
	cParams = append(cParams, fname)

	// locate executable in PATH
	cPath, err := exec.LookPath(cName)

	if err != nil {
		return nil
	}

	// open the file using default program
	exec.Command(cPath, cParams...).Run()

	// give the browser some time to open the file before deleting it
	time.Sleep(2 * time.Second)
	return err
}
