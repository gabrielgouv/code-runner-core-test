package main

import (
	"fmt"

	"github.com/gabrielgouv/code_runner_poc/cmd"
)

// Config object for tests
type Config struct {
	commands []string
}

func main() {

	args := cmd.ParseArgs()

	fmt.Println(args)

	// dirPtr := flag.String("dir", "", "Tells the directory where the files are located. All files inside will be compiled.")
	// execPtr := flag.String("exec", "", "Tells which is the main file to run.")

	// flag.Parse()

	// // config := Config{
	// // 	commands: []string{"javac " + *execPtr},
	// // }

	// fmt.Println("dir=" + *dirPtr)
	// fmt.Println("exec=" + *execPtr)

	// cmd := exec.Command("java", *execPtr)
	// cmd.Dir = *dirPtr

	// var buf bytes.Buffer
	// cmd.Stdout = &buf

	// var errBuf bytes.Buffer
	// cmd.Stderr = &errBuf

	// start := time.Now()

	// cmd.Start()
	// done := make(chan error)

	// go func() {
	// 	done <- cmd.Wait()
	// }()

	// select {
	// case err := <-done:

	// 	elapsed := time.Since(start)

	// 	fmt.Println("Elapsed:", int64(elapsed/time.Millisecond), "ms")

	// 	if errBuf.String() != "" {
	// 		fmt.Println("Error:", errBuf.String())
	// 	} else {
	// 		fmt.Println("Output:", buf.String())
	// 	}

	// 	if err != nil {
	// 		fmt.Println("Non-zero exit code:", err)
	// 	}
	// }

}
