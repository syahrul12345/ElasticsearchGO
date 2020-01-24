package auth

import (
	"backend/models"
	"backend/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

//JwtAuthentication authenticates the received JWT token
var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestPath := request.URL.Path
		//auth is the list of paths that requires authentication
		auth := []string{
			"/dashboard",
			"/api/v1/search",
		}
		// If the current path is not in the list of auth routes, we can serve the http.
		requireAuth := false
		for _, value := range auth {
			if value == requestPath {
				requireAuth = true
				break
			}
		}
		if !requireAuth {
			// if requestPath == "/api/v1/login" {
			// 	// Set headers everytime login api is hit
			// 	// rw := &responsewriter{w: writer}
			// 	// next.ServeHTTP(rw, request)
			// 	return
			// }
			next.ServeHTTP(writer, request)
			return
		}
		//other wise it requires authentication
		response := make(map[string]interface{})
		tokenHeader := request.Header.Get("Authorization")

		if tokenHeader == "" {
			response = utils.Message(false, "Missing auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}
		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}
		tokenPart := splitted[1] // the information that we're interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		//malformed token, return 403
		if err != nil {
			response = utils.Message(false, "Malformed auth token")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}
		//token is invalid
		if !token.Valid {
			response = utils.Message(false, "Token is invalid")
			writer.WriteHeader(http.StatusForbidden)
			writer.Header().Add("Content-Type", "application/json")
			utils.Respond(writer, response)
			return
		}

		//everything went well
		fmt.Sprintf("User ", tk.UserName)
		// Set the cookie

		ctx := context.WithValue(request.Context(), "user", tk.UserID)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}

type responsewriter struct {
	w    http.ResponseWriter
	buf  bytes.Buffer
	code int
}

func (rw *responsewriter) Header() http.Header {
	return rw.w.Header()
}

func (rw *responsewriter) WriteHeader(statusCode int) {
	rw.code = statusCode
}

func (rw *responsewriter) Write(data []byte) (int, error) {
	return rw.buf.Write(data)
}

func (rw *responsewriter) Done() (int64, error) {
	if rw.code > 0 {
		rw.w.WriteHeader(rw.code)
	}
	return io.Copy(rw.w, &rw.buf)
}
