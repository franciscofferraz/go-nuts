package app

import (
	"fmt"
	"net/http"

	"github.com/franciscofferraz/go-struct/internal/utils"
	"go.uber.org/zap"
)

func LogError(r *http.Request, logger *zap.SugaredLogger, err error) {
	logger.Error("An error occurred",
		zap.Error(err),
		zap.String("request_method", r.Method),
		zap.String("request_url", r.URL.String()),
	)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, status int, message interface{}) {
	env := utils.Envelope{"error": message}

	err := utils.WriteJSON(w, status, env, nil, logger)
	if err != nil {
		LogError(r, logger, err)
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, err error) {
	logger.Error("The server encountered a problem and could not process the request", zap.Error(err))

	ErrorResponse(w, r, logger, http.StatusInternalServerError, "The server encountered a problem and could not process your request")
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "The requested resource could not be found"
	ErrorResponse(w, r, logger, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	ErrorResponse(w, r, logger, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, err error) {
	ErrorResponse(w, r, logger, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger, errors map[string]string) {
	ErrorResponse(w, r, logger, http.StatusUnprocessableEntity, errors)
}

func EditConflictResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "unable to update the record due to an edit conflict, please try again"
	ErrorResponse(w, r, logger, http.StatusConflict, message)
}

func RateLimitExceededResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "rate limit exceeded"
	ErrorResponse(w, r, logger, http.StatusTooManyRequests, message)
}

func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "invalid authentication credentials"
	ErrorResponse(w, r, logger, http.StatusUnauthorized, message)
}

func InvalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	ErrorResponse(w, r, logger, http.StatusUnauthorized, message)
}

func AuthenticationRequiredResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "you must be authenticated to access this resource"
	ErrorResponse(w, r, logger, http.StatusUnauthorized, message)
}

func InactiveAccountResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "your user account must be activated to access this resource"
	ErrorResponse(w, r, logger, http.StatusForbidden, message)
}

func NotPermittedResponse(w http.ResponseWriter, r *http.Request, logger *zap.SugaredLogger) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	ErrorResponse(w, r, logger, http.StatusForbidden, message)
}
