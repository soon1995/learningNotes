// copied torbiak/gopl/ex13.4/bzip.go
package main

import (
	"io"
	"os/exec"
)

type Writer struct {
	cmd   exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(w io.Writer) (io.WriteCloser, error) {
	cmd := exec.Cmd{Path: "/bin/bzip2", Stdout: w}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &Writer{cmd, stdin}, nil
}

func (w *Writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

func (w *Writer) Close() error {
	pipeErr := w.stdin.Close()
	cmdErr := w.cmd.Wait()
	if pipeErr != nil {
		return pipeErr
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}

// func main() {
//   file, err := os.Create("abc")
//   if err != nil {
//     log.Fatal(err)
//   }
//   defer func() {
//     file.Close()
//   }()
//
// 	cmd := &exec.Cmd{
// 		Path:   "/bin/bzip2",
//     Args: []string{"bzip2", "test.tx"},
// 		Stdin:  os.Stdin,
// 		Stdout: os.Stdout,
// 		Stderr: os.Stderr,
// 	}
// 	err = cmd.Run()
// 	if err != nil {
// 		log.Fatalf("bzip2: %v\n", err)
// 	}
// }
