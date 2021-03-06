// Copyright 2015 Geofrey Ernest a.k.a gernest, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package controllers

type LoraController struct {
	MainController
}

func (c *LoraController) Pricing() {
	c.ActivateView("lora/pricing")
	c.Data["Title"] = "Pricing"
}
func (c *LoraController) Services() {
	c.ActivateView("lora/services")
	c.Data["Title"] = "Services"
}
func (c *LoraController) Contacts() {
	c.ActivateView("lora/contacts")
	c.Data["Title"] = "Contacts"

}
func (c *LoraController) Legal() {
	c.ActivateContent("lora/legal")
	c.Data["Title"] = "Legal"

}
func (c *LoraController) Terms() {
	c.ActivateContent("lora/terms")
	c.Data["Title"] = "Terms of servie"

}
func (c *LoraController) Help() {
	c.ActivateContent("lora/help")
	c.Data["Title"] = "Help"

}
func (c *LoraController) Companies() {
	c.ActivateView("lora/companies")
	c.Data["Title"] = "Companies"

}
