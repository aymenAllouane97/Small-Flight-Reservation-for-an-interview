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

func LoginCompany(c *gin.Context) {
	var body struct {
		Name     string
		Password string
		Email    string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
		return
	}
	fmt.Println(body)
	var company models.Company
	initializers.DB.First(&company, "email = ?", body.Email)
	if company.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "company not found"})
		return
	}
	res := bcrypt.CompareHashAndPassword([]byte(company.Password), []byte(body.Password))
	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "credentials are wrong"})
		return
	}

	token := generateCompanyJWT(company)

	c.JSON(200, gin.H{
		"company": company,
		"token":   token,
	})
}

func generateCompanyJWT(company models.Company) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       company.ID,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
		"user_type": "company",
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_secret")))
	if err != nil {
		fmt.Println("Error while signing JWT token:", err)
		return ""
	}

	return tokenString
}

func CreateCompany(c *gin.Context) {
	var body struct {
		Name     string
		Password string
		Email    string
	}

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't parse request body"})
		return
	}
	fmt.Println(body, "$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't hash password"})
		return
	}
	company := models.Company{Name: body.Name, Password: string(hash), Email: body.Email}

	result := initializers.DB.Create(&company)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"company": company})
}

func GetAllCompanies(c *gin.Context) {
	var companies []models.Company

	result := initializers.DB.Find(&companies)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

func GetCompany(c *gin.Context) {
	companyID := c.Param("id")

	var company models.Company
	result := initializers.DB.First(&company, companyID)

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})
}

func UpdateCompany(c *gin.Context) {
	companyID := c.Param("id")
	var body struct {
		Name     string
		Password string
		Email    string
	}
	err := c.Bind(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var company models.Company
	result := initializers.DB.First(&company, companyID)
	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if body.Email != "" && body.Email != company.Email {
		var existingCompany models.Company
		err := initializers.DB.Where("email = ?", body.Email).First(&existingCompany).Error
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}
	}

	if body.Name != "" && body.Name != company.Name {
		var existingCompany models.Company
		err := initializers.DB.Where("name = ?", body.Name).First(&existingCompany).Error
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name already exists"})
			return
		}
	}

	if body.Name != "" {
		company.Name = body.Name
	}
	if body.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't hash password"})
			return
		}
		company.Password = string(hash)
	}
	if body.Email != "" {
		company.Email = body.Email
	}

	result = initializers.DB.Save(&company)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})
}

func DeleteCompany(c *gin.Context) {
	companyID := c.Param("id")

	var company models.Company
	result := initializers.DB.First(&company, companyID)

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	result = initializers.DB.Delete(&company)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
func GetPassengersByDate(c *gin.Context) {
	airportID := c.Param("id")
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	fmt.Println(date, " the dateee@@@@@@@@@@@@@@@@@@@2")
	var passengers []string

	result := initializers.DB.Raw("SELECT passengers.name FROM passengers JOIN reservations ON passengers.id = reservations.passenger_id JOIN flights ON reservations.flight_id = flights.id WHERE flights.departure_date = ? AND flights.departure_airport_id = ?", date, airportID).Pluck("name", &passengers)

	if result.Error != nil {
		c.Status(404)
		return
	}
	fmt.Println(result, "lisssssssssssssssst")

	c.JSON(200, gin.H{"passengers": passengers})
}
