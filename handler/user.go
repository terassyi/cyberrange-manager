package handler

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func (h *Handler) UserLogin(c *gin.Context) {
	resource := template.Must(template.ParseFiles("templates/base.html",
		"templates/components/header.html",
		"templates/components/footer.html",
		"templates/login.html"))
	h.SetHTMLTemplate(resource)
	c.HTML(http.StatusOK, "base.html", gin.H{"message": "Login to User panel"})
}

func (h *Handler) UserLogout(c *gin.Context) {
	// TODO delete session
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func (h *Handler) SignUp(c *gin.Context) {
	resource := template.Must(template.ParseFiles("templates/base.html",
		"templates/components/header.html",
		"templates/components/footer.html",
		"templates/signup.html"))
	h.SetHTMLTemplate(resource)
	c.HTML(http.StatusOK, "base.html", nil)
}

func (h *Handler) UserHome(c *gin.Context) {
	// TODO check session

	resource := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/user/user_header.html",
		"templates/components/footer.html",
		"templates/user/home.html",
	))
	h.SetHTMLTemplate(resource)
	c.HTML(http.StatusOK, "base.html", nil)
}
