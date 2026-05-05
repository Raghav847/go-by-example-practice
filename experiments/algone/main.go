package main

import (
	"bytes"
	"fmt"
	"io"
)

var data = []struct {
	input  []byte
	output []byte
}{
	{[]byte("abc"), []byte("abc")},
	{[]byte("elvis"), []byte("Elvis")},
	{[]byte("aElvis"), []byte("aElvis")},
	{[]byte("abcelvis"), []byte("abcElvis")},
	{[]byte("eelvis"), []byte("eElvis")},
	{[]byte("aelvis"), []byte("aElvis")},
	{[]byte("aabeeeelvis"), []byte("aabeeeElvis")},
	{[]byte("e l v i s"), []byte("e l v i s")},
	{[]byte("aa bb e l v i saa"), []byte("aa bb e l v i saa")},
	{[]byte(" elvi s"), []byte(" elvi s")},
	{[]byte("elvielvis"), []byte("elviElvis")},
	{[]byte("elvielvielviselvi1"), []byte("elvielviElviselvi1")},
	{[]byte("elvielviselvis"), []byte("elviElvisElvis")},
}

// assembleInputStream combines all the input into a
// single stream for processing.
func assembleInputStream() []byte {
	var in []byte
	for _, d := range data {
		in = append(in, d.input...)
	}
	return in
}

// assembleOutputStream combines all the output into a
// single stream for comparing.
func assembleOutputStream() []byte {
	var out []byte
	for _, d := range data {
		out = append(out, d.output...)
	}
	return out
}

func main() {
	var output bytes.Buffer
	in := assembleInputStream()
	out := assembleOutputStream()

	find := []byte("elvis")
	repl := []byte("Elvis")

	fmt.Println("=======================================\nRunning Algorithm One")
	output.Reset()
	algOne(in, find, repl, &output)
	matched := bytes.Compare(out, output.Bytes())
	fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())

}

func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	input := bytes.NewBuffer(data)
	size := len(find)

	buf := make([]byte, size)
	end := size - 1

	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {
		if _, err := io.ReadFull(input, buf[end:]); err != nil {
			output.Write(buf[:end])
			return
		}

		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}
			continue
		}
		// Write the front byte since it has been compared.
		output.WriteByte(buf[0])

		// Slice that front byte out.
		copy(buf, buf[1:])
	}

}
