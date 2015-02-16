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
func (c *LoraController) Companies() {
	c.ActivateView("lora/companies")
}
