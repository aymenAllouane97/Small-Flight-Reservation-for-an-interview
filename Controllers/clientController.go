package controllers

import (
	initializers "awesomeProject/Initializers"
	models "awesomeProject/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func LoginClient(c *gin.Context) {
	var body struct {
		Name     string
		Password string
		Email    string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't "})
		return
	}
	fmt.Println(body)
	var client models.Client
	initializers.DB.First(&client, "email = ?", body.Email)
	if client.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "client not found"})
		return
	}
	res := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(body.Password))
	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "credentials are wrong"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       client.ID,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
		"user_type": "client",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		fmt.Println(os.Getenv("jwt_secret"))
		fmt.Println("Error while signing JWT token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(200, gin.H{
		"token": tokenString,
	})

}

func CreateClient(c *gin.Context) {
	var body struct {
		Name     string
		Password string
		Email    string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't "})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't hash password"})
		return
	}
	client := models.Client{Name: body.Name, Password: string(hash), Email: body.Email}

	result := initializers.DB.Create(&client)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(201, gin.H{"client": client})
}

func GetAllClients(c *gin.Context) {
	var clients []models.Client

	result := initializers.DB.Find(&clients)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"clients": clients})
}

func GetClient(c *gin.Context) {
	clientID := c.Param("id")

	var client models.Client
	result := initializers.DB.First(&client, clientID)

	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{"client": client})
}
func UpdateClient(c *gin.Context) {
	clientID := c.Param("id")
	var body struct {
		Name     string
		Password string
		Email    string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(400)
		return
	}

	var client models.Client
	result := initializers.DB.First(&client, clientID)
	if result.Error != nil {
		c.Status(404)
		return
	}

	// Check if the provided email already exists
	if body.Email != "" && body.Email != client.Email {
		var existingClient models.Client
		err := initializers.DB.Where("email = ?", body.Email).First(&existingClient).Error
		if err == nil {
			// Email already exists, return an error
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}
	}

	// Check if the provided name already exists
	if body.Name != "" && body.Name != client.Name {
		var existingClient models.Client
		err := initializers.DB.Where("name = ?", body.Name).First(&existingClient).Error
		if err == nil {
			// Name already exists, return an error
			c.JSON(400, gin.H{"error": "Name already exists"})
			return
		}
	}

	// Update fields only if they are provided in the request
	if body.Name != "" {
		client.Name = body.Name
	}
	if body.Password != "" {
		client.Password = body.Password
	}
	if body.Email != "" {
		client.Email = body.Email
	}

	result = initializers.DB.Save(&client)
	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{"client": client})
}

func DeleteClient(c *gin.Context) {
	clientID := c.Param("id")

	var client models.Client
	result := initializers.DB.First(&client, clientID)

	if result.Error != nil {
		c.Status(404)
		return
	}

	result = initializers.DB.Delete(&client)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.Status(204)
}
