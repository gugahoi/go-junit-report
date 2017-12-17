package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	outputFilePath := flag.String("o", "", "file path to store the output (if blank, outputs to stdout)")
	flag.Parse()

	output, close := SetOutput(*outputFilePath)
	defer close()

	files := flag.Args()
	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed reading file '%s': %v\n", f, err)
		}

		var report TestSuites
		xml.Unmarshal([]byte(data), &report)

		for _, s := range report.TestSuite {
			for _, c := range s.TestCases {
				if c.Failure != "" {
					WriteToBuffer(output, c.Name, c.ClassName, c.Failure)
				}
			}
		}
	}
}

// WriteToBuffer writes a block with detailed failure output into a given writer
func WriteToBuffer(buf io.Writer, name string, class string, body string) error {
	tmpl := `
	<details>
		<summary>
			<code>%s in %s</code>
		</summary>
		<code>
			<pre>%s</pre>
		</code>
	</details>
`
	_, err := fmt.Fprintf(buf, tmpl, name, class, body)
	if err != nil {
		return err
	}
	return nil
}

// SetOutput returns a writer to a given file path, if none is given it returns os.Stdout
func SetOutput(filePath string) (output io.Writer, close func() error) {
	if filePath == "" {
		output = os.Stdout
		return
	}

	f, err := os.Create(filePath)
	close = f.Close
	if err != nil {
		close()
		fmt.Fprintf(os.Stderr, "Unable to open output file '%s': %v", filePath, err)
		os.Exit(1)
	}

	output = f
	return
}

// TestSuites are JUnit test suites
type TestSuites struct {
	Name      string      `xml:"name,attr"`
	TestSuite []TestSuite `xml:"testsuite"`
}

// TestSuite is a JUnit test suite
type TestSuite struct {
	Name      string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Errors    int        `xml:"errors,attr"`
	Failures  int        `xml:"failures,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Time      float32    `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}

// TestCase is a JUnit test case
type TestCase struct {
	ClassName string  `xml:"classname,attr"`
	Name      string  `xml:"name,attr"`
	Time      float32 `xml:"time,attr"`
	Failure   string  `xml:"failure"`
}
