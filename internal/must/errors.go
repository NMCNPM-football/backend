package must

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	//CommonErrors
	ErrInvalidCredentials  = &Error{Code: codes.Unauthenticated, Message: "Invalid credentials."}
	ErrInternalServerError = &Error{Code: codes.Internal, Message: "Internal server error"}
	ErrTokenExpired        = &Error{Code: codes.Unauthenticated, Message: "Token expired."}

	//Users
	ErrEmailExists     = &Error{Code: codes.AlreadyExists, Message: "Email already exists."}
	ErrInvalidEmail    = &Error{Code: codes.InvalidArgument, Message: "Must be business email."}
	ErrInvalidPassword = &Error{Code: codes.InvalidArgument, Message: "Invalid password."}
	ErrEmailNotExists  = &Error{Code: codes.NotFound, Message: "Email doesn't exist."}
)

type Error struct {
	Code    codes.Code `json:"code"`
	Message string     `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) GetCode() codes.Code {
	return e.Code
}

func HandlerError(err error, logger *zap.Logger) error {
	var detectErr *Error
	switch {
	case errors.As(err, &detectErr):
		e := err.(*Error)
		return status.Error(e.Code, e.Error())
	default:
		logger.Error("s.ErrorResponseHandle", zap.Error(err))
		return status.Error(ErrInternalServerError.Code, ErrInternalServerError.Error())
	}
}

func HttpErrorException(logger *zap.Logger, w http.ResponseWriter, err error) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		switch err.(type) {
		case *Error:
			errJson, _ := json.Marshal(err.(*Error))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(string(errJson)))
			return
		default:
			logger.Error("ErrorException", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			errJson, _ := json.Marshal(ErrInternalServerError)
			w.Write([]byte(string(errJson)))
			return
		}
	}
}

func HttpErrorAuthorize(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	errJson, _ := json.Marshal(ErrInvalidCredentials)
	w.Write([]byte(string(errJson)))
	return
}

func handleRoutingError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	fmt.Println("handleRoutingError", httpStatus)

	if httpStatus != http.StatusMethodNotAllowed {
		runtime.DefaultRoutingErrorHandler(ctx, mux, marshaler, w, r, httpStatus)
		return
	}

	// Use HTTPStatusError to customize the DefaultHTTPErrorHandler status code
	err := &runtime.HTTPStatusError{
		HTTPStatus: http.StatusInternalServerError,
		Err:        status.Errorf(http.StatusInternalServerError, ErrInternalServerError.Message),
	}

	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}
