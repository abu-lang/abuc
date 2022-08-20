package preprocessor

// TrivialStream is a minimal interface for chunk-based text streams.
type TrivialStream interface {
	// Size returns the number of the chunks that compose the stream.
	Size() int
	// GetText returns the section of the content delimited by the
	// specified chunk indexes (e.g. 0,Size()-1).
	GetText(int, int) string
}
