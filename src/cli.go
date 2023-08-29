package src

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"slices"
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
	fmt.Fprintf(w, "==========\n\n")
}

// Print all the status codes in a tabular format
//
// w: the tabwriter to print to
//
// category: the status code category to print, empty prints all
//
// returns: any error encountered
func printStatusCodes(w *tabwriter.Writer, category string) error {
	defer w.Flush()

	fmt.Fprintf(w, "Status Codes\n")
	fmt.Fprintf(w, "------------\n\n")

	statusCodesSubset := statusCodes

	// Category header index
	if category != "" {
		catStart := slices.Index(
			statusCodes,
			fmt.Sprintf("------------------ %sxx ------------------", category),
		)

		catInt, _ := strconv.Atoi(category)
		catStr := strconv.Itoa(catInt + 1)

		catEnd := slices.Index(
			statusCodes,
			fmt.Sprintf("------------------ %sxx ------------------", catStr),
		)

		if catStart == -1 {
			return fmt.Errorf(
				"error: invalid category %s; allowed categories are 1, 2, 3, 4, 5",
				category,
			)
		}
		if catEnd == -1 {
			catEnd = len(statusCodes)
		}

		statusCodesSubset = statusCodes[catStart:catEnd]
	}

	// Print category headers, e.g. '---- 1xx ----'
	for idx, code := range statusCodesSubset {
		if idx == 0 && strings.HasPrefix(code, "-") {
			fmt.Fprintf(w, "%s%s\n", code, "\n")
			continue
		}

		// Print a newline before and after the category header
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
	return nil
}

// Print the status text for a given status code
//
// w: the tabwriter to print to
//
// code: the status code to lookup
//
// returns: any error encountered
func printStatusText(w *tabwriter.Writer, code string) error {
	defer w.Flush()

	statusText, ok := statusCodeMap[code]
	if !ok {
		return fmt.Errorf("error: invalid status code %s", code)
	}
	fmt.Fprintln(w, formatStatusText(statusText))
	return nil
}

// Assemble
//
// w: the tabwriter to print to
//
// version: the version string
//
// exitFunc: the function to call to exit
func Cli(w *tabwriter.Writer, version string, exitFunc func(int)) {
	// Flush the writer at the end of the function
	defer w.Flush()

	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.CommandLine = fs

	// Set the default output to the passed tabwriter
	fs.SetOutput(w)

	code := flag.String("code", "", "Print the description of an HTTP status code")
	help := flag.Bool("help", false, "Print usage")
	vers := flag.Bool("version", false, "Print version")
	list := flag.Bool("list", false, "Print HTTP status codes")
	cat := flag.String("cat", "",
		"Print HTTP status codes by category with -list; \n"+
			"allowed categories are 1, 2, 3, 4, 5",
	)

	// Print the header
	printHeader(w)

	// Override the default usage to flush the tabwriter
	flagUsageOld := fs.Usage
	fs.Usage = func() {
		flagUsageOld()
		w.Flush()
	}

	err := fs.Parse(os.Args[1:])
	if err != nil {
		exitFunc(2)
	}

	// If no arguments are provided, print the help message
	if len(os.Args) < 2 || *help {
		flag.Usage()
		return
	}

	// If the version flag is provided, print the version
	if *vers {
		fmt.Fprintln(w, version)
		return
	}

	if *cat != "" && !*list {
		fmt.Fprintln(w, "error: cannot use -cat without -list")
		exitFunc(2)
	}

	// Cat must be provided with the list flag but cat is optional
	if *list {
		err := printStatusCodes(w, *cat)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
			exitFunc(1)
		}
		return
	}

	// If the code flag is provided, print the status text
	if *code != "" {
		err := printStatusText(w, *code)
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
			exitFunc(1)
		}
		return
	}
}
