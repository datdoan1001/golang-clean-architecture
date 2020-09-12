package controller

// AppController struct
type AppController struct {
	User interface{ UserController }
}
