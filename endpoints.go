package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userCredentials struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// UserLoginResponse holds id and role the user is authenticated for
type UserLoginResponse struct {
	ID   string `json:"user_id"`
	Role string `json:"role"`
}

type userSignupResponse struct {
	ID   string `json:"user_id"`
	Role string `json:"role"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type tokenMessage struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func checkTokenAgeMaturity(issuedAt int64) bool {
	return time.Now().Unix() < time.Unix(issuedAt, 0).Add(minTokenAge).Unix()
}

func loginHandl(context *gin.Context) {
	var creds userCredentials

	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, err := json.Marshal(creds)

	resp, err := http.Post(userManagementServer+"/login", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		currentError := errorResponse{}
		if jsonErr := json.Unmarshal(respBody, &currentError); jsonErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": jsonErr.Error()})
			return
		}

		context.JSON(resp.StatusCode, currentError)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	currentUser := UserLoginResponse{}
	jsonErr := json.Unmarshal(respBody, &currentUser)
	if jsonErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := generateNewToken(currentUser.ID, currentUser.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: token,
		Role:  currentUser.Role,
	}
	context.JSON(http.StatusOK, tokenResp)
}

func signupHandl(context *gin.Context) {
	var creds userCredentials

	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, err := json.Marshal(creds)

	resp, err := http.Post(userManagementServer+"/signup", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		currentError := errorResponse{}
		if jsonErr := json.Unmarshal(respBody, &currentError); jsonErr != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": jsonErr.Error()})
			return
		}

		context.JSON(resp.StatusCode, currentError)
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	currentUser := userSignupResponse{}
	jsonErr := json.Unmarshal(respBody, &currentUser)
	if jsonErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := generateNewToken(currentUser.ID, currentUser.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: token,
		Role:  currentUser.Role,
	}
	context.JSON(http.StatusCreated, tokenResp)
}

func validateTokenHandl(context *gin.Context) {
	token := context.MustGet("encodedToken").(string)

	// Parse and validate token
	parsedToken, ok, oldKey, err := validateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": err.Error()})
		return
	}
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": "wrong signiture"})
		return
	}

	// TODO: what to do when token is generated by old key
	if oldKey {
		log.Println(oldKey)
	}

	context.JSON(http.StatusOK, gin.H{"token": parsedToken})
}

func renewTokenHandl(context *gin.Context) {
	token := context.MustGet("encodedToken").(string)

	// Parse and validate token
	parsedToken, ok, _, err := validateToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": err.Error()})
		return
	}
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid", "reason": "wrong signiture"})
		return
	}

	// Check for too frequent requests:
	if checkTokenAgeMaturity(parsedToken.StandardClaims.IssuedAt) {
		context.JSON(http.StatusTeapot, gin.H{"error": "can't renew token so often"})
		return
	}

	// Generate new token:
	newToken, err := generateNewToken(parsedToken.UserID, parsedToken.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send response
	tokenResp := tokenMessage{
		Token: newToken,
	}
	context.JSON(http.StatusOK, tokenResp)
}
