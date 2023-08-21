package errors

import (
	"fmt"
	"net/http"

	"github.com/franciscofferraz/go-struct/internal/logger"
	"github.com/franciscofferraz/go-struct/internal/utils"
	"go.uber.org/zap"
)

func LogError(r *http.Request, err error) {
	logger.Logger.Error("An error occurred",
		zap.Error(err),
		zap.String("request_method", r.Method),
		zap.String("request_url", r.URL.String()),
	)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := utils.Envelope{"error": message}

	err := utils.WriteJSON(w, status, env, nil, logger.Logger)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logger.Logger.Error("The server encountered a problem and could not process the request", zap.Error(err))
	ErrorResponse(w, r, http.StatusInternalServerError, "the server encountered a problem and could not process your request")
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusNotFound, "the requested resource could not be found")
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusMethodNotAllowed, fmt.Sprintf("the %s method is not supported for this resource", r.Method))
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusConflict, "unable to update the record due to an edit conflict, please try again")
}

func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusTooManyRequests, "rate limit exceeded")
}

func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusUnauthorized, "invalid authentication credentials")
}

func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	ErrorResponse(w, r, http.StatusUnauthorized, "invalid or missing authentication token")
}

func AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusUnauthorized, "you must be authenticated to access this resource")
}

func NotPermittedResponse(w http.ResponseWriter, r *http.Request) {
	ErrorResponse(w, r, http.StatusForbidden, "your user account doesn't have the necessary permissions to access this resource")
}
