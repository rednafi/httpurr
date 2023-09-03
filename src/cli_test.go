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

	// Test trimming newlines from both ends
	in = "\nHello\n"
	want = "Hello"

	out = formatStatusText(in)
	if out != want {
		t.Errorf("formatStatusText(%q) = %q, want %q", in, out, want)
	}

	// Test with no surrounding whitespaces
	in = "Hello"
	want = "Hello"

	out = formatStatusText(in)
	if out != want {
		t.Errorf("formatStatusText(%q) = %q, want %q", in, out, want)
	}
}

func TestPrintHeader(t *testing.T) {
	// Create a tabwriter with a buffer
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)

	// Call function
	printHeader(w)

	// Check output
	got := buf.String()
	want := "\nᗢ httpurr\n==========\n\n"

	if got != want {
		t.Errorf("printHeader() = %q, want %q", got, want)
	}
}

func TestPrintStatusCodes(t *testing.T) {
	// Create a tabwriter with a buffer
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)

	// Call function
	_ = printStatusCodes(w, "")

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
		"100    Continue",
		"101    Switching Protocols",
		"102    Processing",
		"103    Early Hints",
		"200    OK",
		"201    Created",
		"202    Accepted",
		"203    Non-Authoritative Information",
		"204    No Content",
		"205    Reset Content",
		"206    Partial Content",
		"207    Multi-Status",
		"208    Already Reported",
		"226    IM Used",
		"300    Multiple Choices",
		"301    Moved Permanently",
		"302    Found",
		"303    See Other",
		"304    Not Modified",
		"305    Use Proxy",
		"-",
		"307    Temporary Redirect",
		"308    Permanent Redirect",
		"400    Bad Request",
		"401    Unauthorized",
		"402    Payment Required",
		"403    Forbidden",
		"404    Not Found",
		"405    Method Not Allowed",
		"406    Not Acceptable",
		"407    Proxy Authentication Required",
		"408    Request Timeout",
		"409    Conflict",
		"410    Gone",
		"411    Length Required",
		"412    Precondition Failed",
		"413    Request Entity Too Large",
		"414    Request URI Too Long",
		"415    Unsupported Media Type",
		"416    Requested Range Not Satisfiable",
		"417    Expectation Failed",
		"418    I'm a teapot",
		"421    Misdirected Request",
		"422    Unprocessable Entity",
		"423    Locked",
		"424    Failed Dependency",
		"425    Too Early",
		"426    Upgrade Required",
		"428    Precondition Required",
		"429    Too Many Requests",
		"431    Request Header Fields Too Large",
		"451    Unavailable For Legal Reasons",
		"500    Internal Server Error",
		"501    Not Implemented",
		"502    Bad Gateway",
		"503    Service Unavailable",
		"504    Gateway Timeout",
		"505    HTTP Version Not Supported",
		"506    Variant Also Negotiates",
		"507    Insufficient Storage",
		"508    Loop Detected",
		"510    Not Extended",
		"511    Network Authentication Required",
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
	// Create a tabwriter with a buffer
	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)

	code := "100"
	_ = printStatusText(w, code)

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
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	// Test --
	os.Args = []string{"cli", "--help"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "Usage") {
		t.Errorf("Expected help text to be printed")
	}

	// Test -
	buf.Reset()
	os.Args = []string{"cli", "-h"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "Usage") {
		t.Errorf("Expected help text to be printed")
	}
}

func TestCliVersion(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	// Test --
	os.Args = []string{"cli", "--version"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "v1.0") {
		t.Errorf("Expected version to be printed")
	}

	// Test -
	buf.Reset()
	os.Args = []string{"cli", "-v"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "v1.0") {
		t.Errorf("Expected version to be printed")
	}
}

func TestCliList(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	// Test --
	os.Args = []string{"cli", "--list"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "418") {
		t.Errorf("Expected status codes to be printed")
	}

	// Test -
	buf.Reset()
	os.Args = []string{"cli", "-l"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "418") {
		t.Errorf("Expected status codes to be printed")
	}

}

func TestCliCode(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	// Test --
	os.Args = []string{"cli", "-code", "404"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "404 Not Found") {
		t.Errorf("Expected 404 status text to be printed")
	}

	// Test -
	buf.Reset()
	os.Args = []string{"cli", "-c", "404"}

	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "404 Not Found") {
		t.Errorf("Expected 404 status text to be printed")
	}
}

func TestCliCat(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	os.Args = []string{"cli", "-list", "-cat", "1"}

	Cli(w, "v1.0", func(int) {})

	wantLines := []string{
		"100    Continue",
		"101    Switching Protocols",
		"102    Processing",
		"103    Early Hints",
	}

	for _, want := range wantLines {
		if !strings.Contains(buf.String(), want) {
			t.Errorf("expected 1xx status codes to be printed; want %q", want)
		}
	}
}

func TestCliError(t *testing.T) {
	// Must reset flag.CommandLine to avoid "flag redefined" error
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	// Invalid flag
	os.Args = []string{"cli", "--invalid"}
	Cli(w, "v1.0", func(int) {})
	if !strings.Contains(buf.String(), "flag provided but not defined: -invalid") {
		t.Errorf("Expected error message to be printed, got %s\n", buf.String())
	}

	// Missing code
	os.Args = []string{"cli", "--code"}
	Cli(w, "v1.0", func(int) {})
	if !strings.Contains(buf.String(), "flag needs an argument: -code") {
		t.Errorf("Expected error message to be printed, got %s\n", buf.String())
	}
	buf.Reset()

	// Invalid code
	os.Args = []string{"cli", "--code", "999"}
	Cli(w, "v1.0", func(int) {})
	if !strings.Contains(buf.String(), "error: invalid status code 999") {
		t.Errorf("Expected error message to be printed, got %s\n", buf.String())
	}
	buf.Reset()

	// Using --cat without --list
	os.Args = []string{"cli", "--cat", "6"}
	Cli(w, "v1.0", func(int) {})
	if !strings.Contains(buf.String(), "error: cannot use --cat without --list") {
		t.Errorf("Expected error message to be printed, got %s\n", buf.String())
	}
	buf.Reset()

	// Using --cat with --list but invalid category
	os.Args = []string{"cli", "--list", "--cat", "6"}
	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(),
		"error: invalid category 6; allowed categories are 1, 2, 3, 4, 5") {
		t.Errorf("Expected error message to be printed, got %s\n", buf.String())
	}
	buf.Reset()

	// Using --list after --cat but without category value
	os.Args = []string{"cli", "--cat", "-list"}
	Cli(w, "v1.0", func(int) {})

	if !strings.Contains(buf.String(), "error: cannot use --cat without --list") {
		t.Errorf("Expected error message to be printed, got %s\n", buf.String())
	}
	buf.Reset()
}

// ================== Cli tests end ==================
