package controllers

type LoraController struct {
	MainController
}

func (c *LoraController) Pricing() {
	c.ActivateView("lora/pricing")
}
func (c *LoraController) Services() {
	c.ActivateView("lora/services")
}
func (c *LoraController) Contacts() {
	c.ActivateView("lnotyet")
}
func (c *LoraController) Legal() {
	c.ActivateContent("notyet")
}
func (c *LoraController) Terms() {
	c.ActivateContent("notyet")
}
func (c *LoraController) Help() {
	c.ActivateContent("notyet")
}
