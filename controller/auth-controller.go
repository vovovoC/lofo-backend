package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vovovoC/lofo-backend/dto"
	"github.com/vovovoC/lofo-backend/entity"
	"github.com/vovovoC/lofo-backend/res"
	"github.com/vovovoC/lofo-backend/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

var resp, err = http.Get("https://golangcode.com")

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := res.BuildErrorResponse(resp.StatusCode, "Failed to process request", errDTO.Error(), res.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := res.BuildDataResponse(http.StatusOK, "OK", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := res.BuildErrorResponse(http.StatusUnauthorized, "Please check again your credential", "Invalid Credential", res.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := res.BuildErrorResponse(http.StatusBadRequest, "Failed to process request", errDTO.Error(), res.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := res.BuildErrorResponse(http.StatusConflict, "Failed to process request", "Duplicate email", res.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := res.BuildDataResponse(http.StatusCreated, "OK", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
