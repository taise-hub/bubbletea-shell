package main

import (
	"github.com/creack/pty"
	"golang.org/x/term"
	"io"
	"os"
	"os/exec"
)

type Terminal struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (t *Terminal) Run() error {
	c := exec.Command("bash")
	ptmx, err := pty.Start(c)
	defer func() { _ = ptmx.Close() }()
	if err != nil {
		return err
	}
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()
	go func() { io.Copy(ptmx, t.Stdin) }()
	io.Copy(t.Stdout, ptmx)
	return nil
}

func (t *Terminal) SetStdin(r io.Reader) {
	if t.Stdin == nil {
		t.Stdin = r
	}
}

func (t *Terminal) SetStdout(w io.Writer) {
	if t.Stdout == nil {
		t.Stdout = w
	}
}

func (t *Terminal) SetStderr(w io.Writer) {
	if t.Stderr == nil {
		t.Stderr = w
	}
}
