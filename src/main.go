package src

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/terassyi/cyberrange-manager/src/handler"
	"net/http"
)

var (
	db *sql.DB
)

const (
	DRIVER           = "mysql"
	DATA_SOURCE_NAME = "root:toor@tcp([localhost]:3306)/panel?parseTime=true" // TODO decide database name
)

func init() {
	var initErr error
	db, initErr = sql.Open(DRIVER, DATA_SOURCE_NAME)
	if initErr != nil {
		panic(initErr)
	}
}
func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("secret", store))

	router.LoadHTMLGlob("./templates/**/*.html")
	//router.LoadHTMLGlob("./templates/admin/*.html")

	h := handler.New(router, db)

	// handler function
	router.GET("/", h.Index)
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	router.GET("/admin/login", h.AdminLogin)

	router.GET("/signup", h.SignUp)
	router.GET("/login", h.UserLogin)

	router.GET("/home", h.UserHome)
	router.GET("/logout", h.UserLogout)

	router.GET("/admin/dashboard", h.DashBoard)
	router.GET("/admin/logout", h.AdminLogout)

	router.Run()
}
