package models

type (
	// RegistrationForm has info strings for verification during registration process
	RegistrationForm struct {
		UserName string `valid:"Required"`
		Company  string `valid:"Required"`
		Email    string `valid:"Email"`
		Password string `valid:"MinSize(6)"`
		Confirm  string `valid:"Required"`
	}

	UserForm struct {
		Company string `valid:"Required"`
	}
	UserProfileForm struct {
		Phone string `valic:"Required`
	}
)
