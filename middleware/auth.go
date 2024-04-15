package middleware

import (
 "fmt"
 "net/http"
 "strings"

 "github.com/gin-gonic/gin"
 "github.com/golang-jwt/jwt/v5"
)

func RequireAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
            c.Abort()
            return
        }

        // The token should be prefixed with "Bearer "
        tokenParts := strings.Split(tokenString, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
            c.Abort()
            return
        }

        tokenString = tokenParts[1]

        _, err := VerifyToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
            c.Abort()
            return
        }

        c.Next()
    }
}

var secretKey = []byte("secretpassword")
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method

        return secretKey, nil
    })

    // Check for errors
    if err != nil {
        return nil, err
    }

    // Validate the token
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("Invalid token")
}
