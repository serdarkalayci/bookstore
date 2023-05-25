package application

import "fmt"

// ErrorCannotFindBookStock is used when no Book stock data cannot be retrieved on the underlying data source
type ErrorCannotFindBookStock struct {
	ISBN string
}

func (e *ErrorCannotFindBookStock) Error() string {
	return fmt.Sprintf("Cannot load the book stock from the data source with ISBN:%s", e.ISBN)
}

// ErrorParsePayload is used when the payload is cannot be parsed by the communications package
type ErrorParsePayload struct{}

func (e *ErrorParsePayload) Error() string {
	return "Cannot parse payload"
}

// ErrorReadPayload is used when the payload is cannot be read by the communications package
type ErrorReadPayload struct{}

func (e *ErrorReadPayload) Error() string {
	return "Cannot read payload"
}

// ErrorPayloadMissing is used when the communication package expects a payload and there is none
type ErrorPayloadMissing struct{}

func (e *ErrorPayloadMissing) Error() string {
	return "Payload is missing"
}
