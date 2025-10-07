// 代码生成时间: 2025-10-08 01:45:30
package controllers

import (
    "encoding/json"
    "errors"
    "net/http"
    "revel"
    "revel/sessions"
    "strings"
    "time"
)

const CSRF_TOKEN_NAME = "CSRF-TOKEN"

// CSRFToken represents the CSRF token structure.
type CSRFToken struct {
    Token string    "json:"token"
    Expiry time.Time "json:"expiry"
}

// VerifyCSRFToken checks if the provided CSRF token is valid.
func VerifyCSRFToken(token string) error {
    // Retrieve the session.
    session := sessions.SessionStartedCtx.revel.RevelSession
    if session == nil {
        return errors.New("Session not started")
    }

    // Retrieve the stored CSRF token from the session.
    storedToken, ok := session[CSRF_TOKEN_NAME].(CSRFToken)
    if !ok {
        return errors.New("No CSRF token found in session")
    }

    // Check if the token has expired.
    if time.Now().After(storedToken.Expiry) {
        return errors.New("CSRF token has expired")
    }

    // Compare the provided token with the stored token.
    if token != storedToken.Token {
        return errors.New("CSRF token mismatch")
    }

    return nil
}

// GenerateCSRFToken generates a new CSRF token and stores it in the session.
func GenerateCSRFToken() (string, error) {
    session := sessions.SessionStartedCtx.revel.RevelSession
    if session == nil {
        return "", errors.New("Session not started")
    }

    // Generate a new random token.
    token, err := revel.RandomToken(32)
    if err != nil {
        return "", err
    }

    // Set the expiry time for the token.
    expiry := time.Now().Add(24 * time.Hour)

    // Store the token in the session.
    session[CSRF_TOKEN_NAME] = CSRFToken{Token: token, Expiry: expiry}

    return token, nil
}

// CSRFController is the base controller that handles CSRF protection.
type CSRFController struct {
    revel.Controller
}

// Before is called before each action to check for CSRF token validity.
func (c CSRFController) Before() revel.Result {
    // Get the CSRF token from the form or URL parameter.
    token := c.Params.Form.Get("csrf_token")

    // Check if the token is provided.
    if token == "" {
        return c.RenderError(http.StatusBadRequest, "CSRF token missing")
    }

    // Verify the CSRF token.
    if err := VerifyCSRFToken(token); err != nil {
        return c.RenderError(http.StatusForbidden, err.Error())
    }

    return nil
}

// RenderError renders an error response with a JSON message.
func (c CSRFController) RenderError(statusCode int, message string) revel.Result {
    c.Response.Status = statusCode
    c.Response.ContentType = "application/json"
    return c.RenderJson(map[string]string{"error": message})
}
