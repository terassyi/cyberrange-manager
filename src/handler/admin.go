package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/terassyi/cyberrange-manager/model/account"
	"github.com/terassyi/cyberrange-manager/model/setting"
	"html/template"
	"net/http"
)

// GET /admin/login
func (h *Handler) AdminLogin(c *gin.Context) {
	resource := template.Must(template.ParseFiles("templates/base.html",
		"templates/components/header.html",
		"templates/components/footer.html",
		"templates/login.html"))
	h.SetHTMLTemplate(resource)
	c.HTML(http.StatusOK, "base.html", gin.H{"message": "Login to Admin panel"})
}

func (h *Handler) AdminLoginPost(c *gin.Context) {

	c.Redirect(http.StatusMovedPermanently, "/admin/dashboard")
}

func (h *Handler) AdminLogout(c *gin.Context) {
	// TODO delete session
	c.Redirect(http.StatusMovedPermanently, "/admin/login")
}

func (h *Handler) DashBoard(c *gin.Context) {
	accounts, err := account.GetAccounts(h.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error.")
	}
	settings := setting.Setting{
		Address:      "hogehoge",
		HostName:     "cyberrange.terassyi.net",
		LocalAddress: "192.168.100.1",
		Password:     "test",
		Hub:          "default",
	}
	resource := template.Must(template.ParseFiles(
		"templates/base.html",
		"templates/admin/admin_header.html",
		"templates/admin/settings.html",
		"templates/admin/user_list.html",
		"templates/admin/dashboard.html",
		"templates/components/footer.html",
	))
	h.SetHTMLTemplate(resource)
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Settings": settings,
		"Accounts": accounts,
	})
}
