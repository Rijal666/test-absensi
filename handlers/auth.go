package handlers

import (
	"fmt"
	"log"
	"net/http"
	authdto "test-absensi/dto/auth"
	resultdto "test-absensi/dto/result"
	userdto "test-absensi/dto/user"
	"test-absensi/models"
	jwtToken "test-absensi/pkg/jwt"
	"test-absensi/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var path_file = "http://localhost:5002/uploads/"

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c *gin.Context) {
	// c.Header("content-Type", "application/json")

	dataFile := c.MustGet("dataFile").(string)
	fmt.Println("this is data file: ", dataFile)

	request := authdto.AuthRequset{
		Fullname: c.Request.FormValue("fullname"),
		Email: c.Request.FormValue("email"),
		Password: c.Request.FormValue("password"),
		Image: dataFile,
	}

	fmt.Println("value dari form",request)

	user := models.User{
		Fullname: request.Fullname,
		Email: request.Email,
		Password: request.Password,
		Image: request.Image,
	}
	fmt.Println("disimpan di ", user)
	
	data, err := h.AuthRepository.Register(user)
	fmt.Println("melakukan registrasi",data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "Your registration is successful", Data: data})
}

func (h *handlerAuth) GetUser(c *gin.Context) {
	userLogin := c.MustGet("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	newcurrentTime := time.Now()

	// Format waktu sesuai dengan format yang diinginkan
	newformattedTime := newcurrentTime.Format("02 January 2006 15:04:05")

	request := userdto.UpdateUserRequest {
		UserID: int(userId),
		ClockOut: newformattedTime,
	}

	 c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "User data successfully logout", Data: request})
}


func(h *handlerAuth) Login(c *gin.Context) {
	
	dataFile := c.MustGet("dataFile").(string)
	fmt.Println("this is data file: ", dataFile)
	
	request := authdto.LoginRequest{
		Email: c.Request.FormValue("email"),
		Password: c.Request.FormValue("password"),
		Image: dataFile,
	}

	
	user := models.User{
		Email: request.Email,
		Password: request.Password,
		Image: request.Image,
	}
	
	
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil{
		c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	fmt.Println(user.ClockIn)

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 24hours expired
	
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil{
		log.Println(errGenerateToken)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error" : "Unauthorized"})
		return
	}

	newcurrentTime := time.Now()

	// Format waktu sesuai dengan format yang diinginkan
	newformattedTime := newcurrentTime.Format("02 January 2006 15:04:05")

	loginResponse := authdto.LoginResponse{
		ID: user.ID,
		Email: user.Email,
		Image: user.Image,
		Token: token,
		ClockIn: newformattedTime,
	}
	c.JSON(http.StatusOK, resultdto.SuccessResult{Status: http.StatusOK, Message: "You have successfully completed attendance today", Data: loginResponse})
	
	
}
