package controller

type Controller struct {
	Auth *Auth
}

func NewController() *Controller {

	ctlr := new(Controller)
	ctlr.Auth = NewAuth()

	return ctlr
}
