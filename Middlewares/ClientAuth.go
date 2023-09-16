package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserType string

const (
	ClientType  UserType = "client"
	CompanyType UserType = "company"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		secretKey := []byte(os.Getenv("JWT_SECRET"))

		return secretKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
		return
	}

	sub, subOK := claims["sub"].(float64)
	if !subOK || sub <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		c.Abort()
		return
	}

	userTypeClaim, userTypeClaimOK := claims["user_type"].(string)
	if !userTypeClaimOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User type not specified"})
		c.Abort()
		return
	}

	var userType UserType
	if userTypeClaim == string(ClientType) {
		fmt.Println(ClientType, " in client type ################################")

		userType = ClientType
	} else if userTypeClaim == string(CompanyType) {
		userType = CompanyType
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user type"})
		c.Abort()
		return
	}

	expirationTime := int64(claims["exp"].(float64))
	currentTime := time.Now().Unix()
	if currentTime > expirationTime {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
		c.Abort()
		return
	}
	fmt.Println(userType, " before setting $$$$$$$$$$")

	c.Set("UserID", int(sub))
	c.Set("UserType", string(userType))
	c.Next()
}

func ClientAuthMiddleware(c *gin.Context) {
	userType := c.GetString("UserType")
	fmt.Println(userType, "##################################")

	if userType != string(ClientType) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Client authentication required"})
		c.Abort()
		return
	}
	c.Next()
}

func CompanyAuthMiddleware(c *gin.Context) {
	userType := c.GetString("UserType")

	if userType != string(CompanyType) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Company authentication required"})
		c.Abort()
		return
	}

	c.Next()
}
