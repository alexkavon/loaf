package controllers

type Controller struct{}

type ControllerInterface interface {
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
	Redirect()
}

func (c *Controller) Prepare() {
}

func (c *Controller) Finish() {
}

func (c *Controller) Get() {
}

func (c *Controller) Post() {
}

func (c *Controller) Delete() {
}

func (c *Controller) Put() {
}

func (c *Controller) Head() {
}

func (c *Controller) Patch() {
}

func (c *Controller) Options() {
}

func (c *Controller) Redirect() {
}
