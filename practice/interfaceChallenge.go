package practice

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	wrt io.Writer
}

func PrintViaWriter() {
	c := &Capper{os.Stdout}
	/*
		func Fprintln(w io.Writer, a ...any) (n int, err error)
	*/
	fmt.Fprintln(c, "Hello There")
	name := "Kim"
	age := 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")
	fmt.Fprintln(&Capper{os.Stdout}, name, "is", age, "years old.")
}

func (c *Capper) Write(p []byte) (int, error) {
	// chars have numeric value (byte) under the hood(asci)
	// so we can operate with them as with numbers
	// very similler to C way
	diff := byte('a' - 'A')
	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	/*
		io.Writer is the interface that wraps the basic Write method.
		Write writes len(p) bytes from p to the underlying data stream. It returns the number of bytes written from p (0 <= n <= len(p)) and any error encountered that caused the write to stop early
		So the Capper wrt is io.Writer it recieves array of byte
	*/
	return c.wrt.Write(out)
}
