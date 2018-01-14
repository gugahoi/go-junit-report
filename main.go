package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func usage() {
	fmt.Fprint(os.Stderr, `
      _             _ _                                  _                                      _            
     (_)_   _ _ __ (_) |_      _ __ ___ _ __   ___  _ __| |_       ___ ___  _ ____   _____ _ __| |_ ___ _ __ 
     | | | | | '_ \| | __|____| '__/ _ \ '_ \ / _ \| '__| __|____ / __/ _ \| '_ \ \ / / _ \ '__| __/ _ \ '__|
     | | |_| | | | | | ||_____| | |  __/ |_) | (_) | |  | ||_____| (_| (_) | | | \ V /  __/ |  | ||  __/ |   
    _/ |\__,_|_| |_|_|\__|    |_|  \___| .__/ \___/|_|   \__|     \___\___/|_| |_|\_/ \___|_|   \__\___|_|   
   |__/                                |_|                                                                   

	junit-report-converter report.xml	
	junit-report-converter report1.xml report2.xml	
	`)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		usage()
	}
	for _, f := range flag.Args() {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed reading file '%s': %v\n", f, err)
			continue
		}
		xmlToTemplate(os.Stdout, defaultTemplate, data)
	}
}

func xmlToTemplate(output io.Writer, template string, data []byte) {
	var report TestSuites
	xml.Unmarshal(data, &report)
	for _, s := range report.TestSuite {
		for _, c := range s.TestCases {
			if c.Failure != "" {
				fmt.Fprintf(output, template, c.Name, c.ClassName, c.Failure)
			}
		}
	}
}

// defaultTemplate is what is used if another template is not provided.
var defaultTemplate = `<details>
	<summary>
		<code>
		%s in %s
		</code>
	</summary>
	<code>
		<pre>
		%s
		</pre>
	</code>
</details>
`

// TestSuites are a collection of junit test suites
type TestSuites struct {
	Name      string      `xml:"name,attr"`
	TestSuite []TestSuite `xml:"testsuite"`
}

// TestSuite is a junit test suite
type TestSuite struct {
	Name      string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Errors    int        `xml:"errors,attr"`
	Failures  int        `xml:"failures,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Time      float32    `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}

// TestCase is a junit test case
type TestCase struct {
	ClassName string  `xml:"classname,attr"`
	Name      string  `xml:"name,attr"`
	Time      float32 `xml:"time,attr"`
	Failure   string  `xml:"failure"`
}
