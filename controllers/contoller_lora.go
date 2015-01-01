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
	c.ActivateView("lora/contacts")
}
func (c *LoraController) Legal() {
	c.ActivateContent("lora/legal")
}
func (c *LoraController) Terms() {
	c.ActivateContent("lora/terms")
}
func (c *LoraController) Help() {
	c.ActivateContent("lora/help")
}
