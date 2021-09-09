package log

import "io"

// ensure we always implement io.Writer (compile error otherwise)
var _ io.Writer = (*Writer)(nil)

type Writer struct {
	// config
	Filter []string // Prefixes of log messages to ignore see log.New() in the official Golang log package

	// internals
	log Logger
}

func (w *Writer) Write(p []byte) (n int, err error) {
	line := string(p)
	// skip messages of prefixes we don't want to print (such as "HTTP:")
	for _, filter := range w.Filter {
		if line[0:len(filter)] == filter {
			return 0, nil // this is a filtered prefix
		}
	}
	w.log.Debugf("Fwd log: %s", line)
	return len(line), nil
}
