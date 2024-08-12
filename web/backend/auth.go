package backend

import (
	"net/http"

	"taskflow/lib"
	"taskflow/models"
	"taskflow/web/templates"
	"taskflow/web/templates/components"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type body struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func SignUp(c *gin.Context) {
	// Get the JSON body and decode into variables
	var body body
	if err := c.Bind(&body); err != nil {
		lib.Render(c, components.Error("Can not sign up"))
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		lib.Render(c, components.Error("Can not sign up"))
		return
	}

	// Create the user
	user := models.User{
		Username: body.Username,
		Password: string(hash),
		UserID:   lib.GenerateRandomUserID(),
	}
	if err := lib.DB.Create(&user).Error; err != nil {
		lib.Render(c, components.Error("Can not sign up"))
		return
	}

	setToken(c, user)
	// Redirect to the home page if the user is signed up
	c.Data(200, "text/html", []byte(`<script>window.location.href = "/signin";</script>`))
}

func SignIn(c *gin.Context) {
	// Get the JSON body and decode into variables
	var body body

	if err := c.Bind(&body); err != nil {
		lib.Render(c, components.Error("Can not sign up"))
		return
	}

	// Find the user
	var user models.User
	lib.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		lib.Render(c, components.Error("User not found"))
		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		lib.Render(c, components.Error("Invalid password"))
		return
	}

	setToken(c, user)
	// Redirect to the home page if the user is signed in
	c.Data(200, "text/html", []byte(`<script>window.location.href = "/";</script>`))
}

func setToken(c *gin.Context, user models.User) {
	tokenString, err := lib.CreateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to sign token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 60*60*24, "/", "", false, true)
}

func GetUserID(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return userID.(string)
}

func SignOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/signin")
}

func SignInPageHandler(ctx *gin.Context) {
	lib.RenderWithLayout(ctx, templates.SignIn())
}

func SignUpPageHandler(ctx *gin.Context) {
	lib.RenderWithLayout(ctx, templates.SignUp())
}
