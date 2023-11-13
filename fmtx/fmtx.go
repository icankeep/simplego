package fmtx

import (
	"bytes"
	"fmt"
	_ "github.com/golangci/gofmt/gofmt"
	"go/format"
	"golang.org/x/sync/semaphore"
	"io"
	"io/fs"
	_ "unsafe"
)

type StringWriter struct {
	buf []byte
	str string
}

func (s *StringWriter) Write(p []byte) (n int, err error) {
	s.buf = p
	s.str = string(p)
	return len(p), nil
}

// ---------- Copy from gofmt ---------
type reporterState struct {
	out, err io.Writer
	exitCode int
}

type reporter struct {
	prev  <-chan *reporterState
	state *reporterState
}

type sequencer struct {
	maxWeight int64
	sem       *semaphore.Weighted   // weighted by input bytes (an approximate proxy for memory overhead)
	prev      <-chan *reporterState // 1-buffered
}

// --------------------------
// https://medium.com/@yardenlaif/accessing-private-functions-methods-types-and-variables-in-go-951acccc05a6
//
//go:linkname newSequencer github.com/golangci/gofmt/gofmt.newSequencer
func newSequencer(maxWeight int64, out, err io.Writer) *sequencer

//go:linkname Add github.com/golangci/gofmt/gofmt.(*sequencer).Add
func Add(s *sequencer, weight int64, f func(*reporter) error)

//go:linkname processFile github.com/golangci/gofmt/gofmt.processFile
func processFile(filename string, info fs.FileInfo, in io.Reader, r *reporter) error

// Deprecated
func formatGoCode(code string) (string, error) {
	maxWeight := int64(1)
	out := new(StringWriter)
	err := new(StringWriter)
	s := newSequencer(maxWeight, out, err)
	Add(s, 1, func(r *reporter) error {
		return processFile("<standard input>", nil, bytes.NewReader([]byte(code)), r)
	})
	<-s.prev

	if len(err.str) != 0 {
		return "", fmt.Errorf(err.str)
	}
	return out.str, nil
}

func FormatGoCode(code string) (string, error) {
	fmtCode, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	return string(fmtCode), nil
}
