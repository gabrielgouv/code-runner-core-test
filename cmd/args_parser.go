package cmd

import "flag"

const defaultTimeout int = 5000
const defaultNoOutput bool = false
const defaultStringValue string = ""

type Args struct {
	lang string
	exec string
	dir  string
	out  string
	// tests       []string
	langVersion string
	timeout     int
	noOutput    bool
}

func ParseArgs() (args Args) {
	langPtr := flag.String("lang", defaultStringValue, "Tells to compile/run for that language. If the language has not been defined, Coderunner will try to identify by file extension.")
	execPtr := flag.String("exec", defaultStringValue, "Tells which is the main file to run.")
	dirPtr := flag.String("dir", defaultStringValue, "Tells the directory where the files are located. All files inside will be compiled.")
	outPtr := flag.String("out", defaultStringValue, "File where the data results of the execution will be extracted.")
	// testsPtr := flag.String("tests", defaultStringValue, "Compares the execution result with expected results.")
	langVersionPtr := flag.String("lang-version", defaultStringValue, "Especifies which version of the language to execute must be used.")
	timeoutPtr := flag.Int("timeout", defaultTimeout, "Sets the timeout for execution. Default is 5 seconds.")
	noOutputPtr := flag.Bool("no-output", defaultNoOutput, "Disable output field on result of execution. Default is false.")

	flag.Parse()

	args = Args{*langPtr, *execPtr, *dirPtr, *outPtr /*testsPtr*/, *langVersionPtr, *timeoutPtr, *noOutputPtr}

	return
}
