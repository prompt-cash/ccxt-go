package log

import (
	"io"
	"sync"
)

// ensure we always implement io.WriteCloser (compile error otherwise)
var _ io.WriteCloser = (*Buffer)(nil)

// This is a log buffer of our most recent log lines.
// It can be used to serve the logs to an admin interface (via HTTP , gRPC,...)
type Buffer struct {
	// config
	MaxSize   int // The max number of stored log lines.
	CleanSize int // After reaching MaxSize, reduce the stored log lines to this value.

	// internals
	//logList   list.List // the cached lines. zero value already is an empty list
	logList   []BufferLogLine // the cached lines
	mu        sync.Mutex
	lineCount int64 // the line number of the latest log line
}

type BufferLogLine struct {
	Nr   int64  `json:"nr"`
	Line string `json:"line"`
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.lineCount++
	line := BufferLogLine{
		Nr:   b.lineCount,
		Line: string(p),
	}
	b.logList = append(b.logList, line)
	n = len(p) // we "wrote" all bytes

	// cleanup
	if len(b.logList) >= b.MaxSize {
		//b.logList = b.logList[b.CleanSize:] // for slice. but we still should unset them to free memory, so better create a new slice
		nextList := make([]BufferLogLine, 0, b.MaxSize)
		copy(nextList, b.logList[b.CleanSize:])
		b.logList = nextList
	}

	return
}

func (b *Buffer) Close() error {
	// nothing to do here
	return nil
}

func (b *Buffer) GetLines() []BufferLogLine {
	return b.logList
}
