package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type malformedRequest string

func (m malformedRequest) Error() string {
	return string(m)
}

// decodeJSONBody decodes request body (expect JSON) to v
// idea from https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func decodeJSONBody(r *http.Request, v interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(v)
	if err != nil {
		var syntaxErr *json.SyntaxError
		var unmarshalTypeErr *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxErr):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)",
				syntaxErr.Offset)
			return malformedRequest(msg)

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return malformedRequest(msg)

		case errors.As(err, &unmarshalTypeErr):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)",
				unmarshalTypeErr.Field,
				unmarshalTypeErr.Offset)
			return malformedRequest(msg)

		// Catch the error caused by extra unexpected fields in the request body.
		// We extract the field name from the error message and interpolate it
		// in our custom error message.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			field := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", field)
			return malformedRequest(msg)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return malformedRequest(msg)

		default:
			return err
		}
	}
	return nil
}
