package src

import (
	"bytes"
	"flag"
	"os"
	"strings"
	"testing"

	"text/tabwriter"
)

// ================== Cli helper tests start ==================

func TestFormatStatusText(t *testing.T) {

	// Test trimming newlines from start
	in := "\nHello"
	want := "Hello"

	out := formatStatusText(in)
	if out != want {
		t.Errorf("formatStatusText(%q) = %q, want %q", in, out, want)
	}

	// Test trimming newlines from end
	in = "Hello\n"
	want = "Hello"

	out = formatStatusText(in)
	if out != want {
		t.Errorf("formatStatusText(%q) = %q, want %q", in, out, want)
	}

	// Test trimming newlines from both
	in = "\nHello\n"
	want = "Hello"

	out = formatStatusText(in)
	if out != want {
		t.Errorf("formatStatusText(%q) = %q, want %q", in, out, want)
	}

	// Test with no surrounding whitespace
	in = "Hello"
	want = "Hello"

	out = formatStatusText(in)
	if out != want {
		t.Errorf("formatStatusText(%q) = %q, want %q", in, out, want)
	}
}

func TestPrintHeader(t *testing.T) {
	// Create a tabwriter
	w := tabwriter.Writer{}

	// Capture output
	var buf bytes.Buffer
	w.Init(&buf, 0, 4, 0, ' ', 0)

	// Call function
	printHeader(&w)

	// Check output
	got := buf.String()
	want := "\nᗢ httpurr\n==========\n\n"

	if got != want {
		t.Errorf("printHeader() = %q, want %q", got, want)
	}
}

func TestPrintStatusCodes(t *testing.T) {
	// Create a tabwriter
	w := tabwriter.Writer{}

	// Capture output
	var buf bytes.Buffer
	w.Init(&buf, 0, 4, 0, ' ', 0)

	// Call function
	printStatusCodes(&w)

	// Check output
	got := buf.String()

	// Check for the first line
	want := "Status Codes\n------------\n\n"

	// Check want is in got
	if strings.Contains(got, want) == false {
		t.Errorf("printStatusCodes() = %q, want %q", got, want)
	}

	// Spot check a few lines
	wantLines := []string{
		"100Continue",
		"101Switching Protocols",
		"102Processing",
		"103Early Hints",
		"200OK",
		"201Created",
		"202Accepted",
		"203Non-Authoritative Information",
		"204No Content",
		"205Reset Content",
		"206Partial Content",
		"207Multi-Status",
		"208Already Reported",
		"226IM Used",
		"300Multiple Choices",
		"301Moved Permanently",
		"302Found",
		"303See Other",
		"304Not Modified",
		"305Use Proxy",
		"-",
		"307Temporary Redirect",
		"308Permanent Redirect",
		"400Bad Request",
		"401Unauthorized",
		"402Payment Required",
		"403Forbidden",
		"404Not Found",
		"405Method Not Allowed",
		"406Not Acceptable",
		"407Proxy Authentication Required",
		"408Request Timeout",
		"409Conflict",
		"410Gone",
		"411Length Required",
		"412Precondition Failed",
		"413Request Entity Too Large",
		"414Request URI Too Long",
		"415Unsupported Media Type",
		"416Requested Range Not Satisfiable",
		"417Expectation Failed",
		"418I'm a teapot",
		"421Misdirected Request",
		"422Unprocessable Entity",
		"423Locked",
		"424Failed Dependency",
		"425Too Early",
		"426Upgrade Required",
		"428Precondition Required",
		"429Too Many Requests",
		"431Request Header Fields Too Large",
		"451Unavailable For Legal Reasons",
		"500Internal Server Error",
		"501Not Implemented",
		"502Bad Gateway",
		"503Service Unavailable",
		"504Gateway Timeout",
		"505HTTP Version Not Supported",
		"506Variant Also Negotiates",
		"507Insufficient Storage",
		"508Loop Detected",
		"510Not Extended",
		"511Network Authentication Required",
	}

	for _, want := range wantLines {

		t.Run(want, func(t *testing.T) {
			if !strings.Contains(got, want) {
				t.Errorf("printStatusCodes() = %q, want %q", got, want)
			}
		})
	}

}

func TestPrintStatusText(t *testing.T) {

	w := new(tabwriter.Writer)
	var buf bytes.Buffer
	w.Init(&buf, 0, 8, 1, '\t', 0)

	code := "100"
	printStatusText(w, code)

	wantLines := []string{
		"Description",
		"-----------",
		"The HTTP 100 Continue informational status response",
		"https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100",
	}

	got := buf.String()

	for _, want := range wantLines {
		if !strings.Contains(got, want) {
			t.Errorf("printStatusText(%q) = %q, want %q", code, got, want)
		}
	}

}

// ================== Cli helper tests end ==================

// ================== Cli tests start ==================

func TestCliHelp(t *testing.T) {
	// Must reset flag.CommandLine to avoid 'flag redefined' error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cli", "-help"}

	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)

	Cli("v1.0")

	if !strings.Contains(buf.String(), "Usage") {
		t.Errorf("Expected help text to be printed")
	}
}

func TestCliVersion(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	os.Args = []string{"cli", "-version"}

	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)

	Cli("v1.0")

	if !strings.Contains(buf.String(), "") {
		t.Errorf("Expected version to be printed")
	}
}

func TestCliList(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cli", "-list"}

	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)

	Cli("v1.0")

	if !strings.Contains(buf.String(), "418") {
		t.Errorf("Expected status codes to be printed")
	}
}

func TestCliCode(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	os.Args = []string{"cli", "-code", "404"}

	var buf bytes.Buffer
	flag.CommandLine.SetOutput(&buf)

	Cli("v1.0")

	if !strings.Contains(buf.String(), "404 Not Found") {
		t.Errorf("Expected 404 status text to be printed")
	}
}

// ================== Cli tests end ==================
