package src

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"text/tabwriter"
)

// Trim prefix and suffix newlines from the http status text
func formatStatusText(text string) string {
	return strings.TrimSuffix(strings.TrimPrefix(text, "\n"), "\n")

}

// Always print the header
func printHeader(w *tabwriter.Writer) {
	defer w.Flush()

	fmt.Fprintf(w, "\nᗢ httpurr\n")
	fmt.Fprintf(w, "=========\n\n")
}

// Print all the status in a tabular format
func printStatusCodes(w *tabwriter.Writer) {
	defer w.Flush()

	fmt.Fprintf(w, "Status Codes\n")
	fmt.Fprintf(w, "------------\n\n")

	// Print group headers, e.g. '---- 1xx ----'
	for idx, code := range statusCodes {
		if idx == 0 && strings.HasPrefix(code, "-") {
			fmt.Fprintf(w, "%s%s\n", code, "\n")
			continue
		}

		// Print a newline before and after the group header
		if idx != 0 && strings.HasPrefix(code, "-") {
			fmt.Fprintf(w, "\n%s%s\n", code, "\n")
			continue
		}

		// Print status text
		codeInt, _ := strconv.Atoi(code)
		statusText := http.StatusText(codeInt)

		// Print '-' when the status text is not available
		if statusText == "" {
			fmt.Fprintf(w, "%s\t%s\n", code, "-")
			continue
		}
		fmt.Fprintf(w, "%s\t%s\n", code, statusText)
	}
}

// Print the status text for a given status code
func printStatusText(w *tabwriter.Writer, code string) {
	defer w.Flush()

	statusText := statusCodeMap[code]
	fmt.Fprintln(w, formatStatusText(statusText))
}

// Cli assembly
func Cli() {
	code := flag.String("code", "", "Print description of an HTTP status code")
	help := flag.Bool("help", false, "Print usage")
	list := flag.Bool("list", false, "List all HTTP status codes")

	// Create a tabwriter to print the output in a tabular format
	w := tabwriter.NewWriter(flag.CommandLine.Output(), 0, 4, 4, ' ', 0)
	flag.CommandLine.SetOutput(w)

	prevUsage := flag.Usage
	flag.Usage = func() {
		defer w.Flush()
		printHeader(w)
		prevUsage()
	}

	flag.Parse()

	// If no arguments are provided, print the help message
	if len(os.Args) < 2 || *help {
		flag.Usage()
		return
	}

	// If the list flag is provided, print all the status codes
	if *list {
		printHeader(w)
		printStatusCodes(w)
		return
	}

	// If the code flag is provided, print the status text
	if *code != "" {
		printHeader(w)
		printStatusText(w, *code)
		return
	}
}
