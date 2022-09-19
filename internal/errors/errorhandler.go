package errors

import (
	"fmt"
	"net/http"
	"os"
)

func HandleError(w http.ResponseWriter, err error) {
	var status int
	errorType := GetType(err)
	switch errorType {
	case BadRequest:
		status = http.StatusBadRequest
	case NotFound:
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)

	/*if errorType == NoType {
		log.Errorf(err)
	}*/
	//fmt.Fprintf(os.Stderr, "Stack trace: %+v\n", errors.WithStack(errors.Cause(err)))
	fmt.Fprintf(os.Stderr, "Stack trace: %+v\n", err)
	errorContext := GetContext(err)
	if errorContext != nil {
		fmt.Fprintf(os.Stderr, "Context %s", errorContext)
		//msg, _ := json.Marshal(errorContext)
		if errorContext["field"] == ContextField {
			w.Write([]byte(errorContext["message"]))
		}
	}
}
