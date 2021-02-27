package authentication

import (
	"net/http"
	"strings"
)

//AuthMiddleware verifies the api request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		addCorsHeaders(w)

		if r.Method != "OPTIONS" && r.URL.Path != "/api/authentication/login" {
			authToken := r.Header.Get("Authorization")
			if len(authToken) == 0 {
				addErrorHeaders(w, "Missing Authorization Header")
				return
			}
			authToken = strings.Replace(authToken, "Bearer ", "", 1)
			token, errParsing := ParseToken(authToken)

			if errParsing != nil {
				addErrorHeaders(w, "Authorization Token could not be parsed: "+errParsing.Error())
				return
			}

			accountID, accountIsFound := token.Get("accountID")
			if !accountIsFound {
				addErrorHeaders(w, "Missing accountID")
				return
			}

			sessionID, sesionIDIsFound := token.Get("sessionID")
			if !sesionIDIsFound {
				addErrorHeaders(w, "Missing sessionID")
				return
			}

			errValidation := ValidateToken(authToken, accountID.(string), sessionID.(string), "authToken")
			if errValidation != nil {
				addErrorHeaders(w, "Error verifying JWT token: "+errValidation.Error())
				return
			}
			next.ServeHTTP(w, r)
		} else if r.Method != "OPTIONS" && r.URL.Path == "/api/authentication/login" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusOK)
			return
		}

	})
}

func addErrorHeaders(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(err))
}

func addCorsHeaders(w http.ResponseWriter) {
	// OPTIONS requests https://stackoverflow.com/questions/22972066/how-to-handle-preflight-cors-requests-on-a-go-server
	w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
	w.Header().Add("Access-Control-Allow-Headers", "content-type, Authorization")
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
}
