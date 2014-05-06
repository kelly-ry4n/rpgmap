package server

import (
	db "./mapdb"
	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
)

type User struct {
	User string `form:"username" binding:"required"`
	Pass string `form:"password" binding:"required"`
}

type NewUser struct {
	Username        string
	Password        string
	PasswordConfirm string
}

func hello() string {
	return "Hello World!"

}

func createUser(user NewUser) string {
	err := db.AddUser(user.Username, user.Password)
	if err != nil {
		return fmt.Sprintf("Error Adding User: %S", err)
	}
	return fmt.Sprintf("Added User %s", user)
}

func createUserView(r render.Render) {
	r.HTML(200, "new_user", nil)
}

func RunServer() {

	con := db.GetDbConn()
	defer con.Close()

	m := martini.Classic()
	m.Use(auth.Basic("kelly", "kelly"))
	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
		Charset:   "UTF-8",
	}))

	m.Get("/", hello)
	m.Get("/create_user", createUserView)
	m.Post("/create_user", binding.Bind(NewUser{}), createUser)
	m.Run()
}
