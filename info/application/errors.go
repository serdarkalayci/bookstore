package application

import "fmt"

// ErrorCannotFindBooks is used when no Book data cannot be retrieved on the underlying data source
type ErrorCannotFindBooks struct {}

func (e *ErrorCannotFindBooks) Error() string {
	return fmt.Sprintf("Cannot load the books from the data source")
}

// ErrorCannotFindBook is used when the Book with the given ISBN cannot be found on the underlying data source
type ErrorCannotFindBook struct {
}

func (e *ErrorCannotFindBook) Error() string {
	return fmt.Sprintf("Cannot find the book with filter provided")
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
