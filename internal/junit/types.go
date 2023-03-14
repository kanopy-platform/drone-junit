package junit

import "encoding/xml"

// Testsuites is a collection of JUnit testsuites.
type Testsuites struct {
	XMLName xml.Name `xml:"testsuites" json:"-"`

	Name     string `xml:"name,attr,omitempty" json:"name,omitempty"`
	Time     string `xml:"time,attr,omitempty" json:"time,omitempty"` // total duration in seconds
	Tests    int    `xml:"tests,attr,omitempty" json:"tests"`
	Errors   int    `xml:"errors,attr,omitempty" json:"errors"`
	Failures int    `xml:"failures,attr,omitempty" json:"failures"`
	Skipped  int    `xml:"skipped,attr,omitempty" json:"skipped"`
	Disabled int    `xml:"disabled,attr,omitempty" json:"disabled"`

	Suites []Testsuite `xml:"testsuite,omitempty" json:"suites"`
}

type Testsuite struct {
	// required attributes
	Name     string `xml:"name,attr" json:"name"`
	Tests    int    `xml:"tests,attr" json:"tests"`
	Failures int    `xml:"failures,attr" json:"failures"`
	Errors   int    `xml:"errors,attr" json:"errors"`
	ID       int    `xml:"id,attr" json:"id"`

	// optional attributes
	Disabled  int    `xml:"disabled,attr,omitempty" json:"disabled"`
	Hostname  string `xml:"hostname,attr,omitempty" json:"hostname"`
	Package   string `xml:"package,attr,omitempty" json:"package"`
	Skipped   int    `xml:"skipped,attr,omitempty" json:"skipped"`
	Time      string `xml:"time,attr" json:"time"`                     // duration in seconds
	Timestamp string `xml:"timestamp,attr,omitempty" json:"timestamp"` // date and time in ISO8601
	File      string `xml:"file,attr,omitempty" json:"file"`

	Properties *[]Property `xml:"properties>property,omitempty" json:"properties,omitempty"`
	TestCases  []Testcase  `xml:"testcase,omitempty" json:"testCases,omitempty"`
	SystemOut  *Output     `xml:"system-out,omitempty" json:"systemOut,omitempty"`
	SystemErr  *Output     `xml:"system-err,omitempty" json:"systemErr,omitempty"`
}

// Testcase represents a single test with its results.
type Testcase struct {
	// required attributes
	Name      string `xml:"name,attr" json:"name"`
	ClassName string `xml:"classname,attr" json:"className"`

	// optional attributes
	Time   string `xml:"time,attr,omitempty" json:"time"` // duration in seconds
	Status string `xml:"status,attr,omitempty" json:"status"`

	Skipped   *Result `xml:"skipped,omitempty" json:"skipped,omitempty"`
	Error     *Result `xml:"error,omitempty" json:"error,omitempty"`
	Failure   *Result `xml:"failure,omitempty" json:"failure,omitempty"`
	SystemOut *Output `xml:"system-out,omitempty" json:"systemOut,omitempty"`
	SystemErr *Output `xml:"system-err,omitempty" json:"systemErr,omitempty"`
}

// Property represents a key/value pair.
type Property struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

// Result represents the result of a single test.
type Result struct {
	Message string `xml:"message,attr" json:"message"`
	Type    string `xml:"type,attr,omitempty" json:"type"`
	Data    string `xml:",cdata" json:"data"`
}

// Output represents output written to stdout or sderr.
type Output struct {
	Data string `xml:",cdata" json:"data"`
}
