package models

type UserSignupValidator struct {
	User           UserParams
	UserRepository UserRepository
}
type ValidationFunc func() (bool, string)

type UserParams struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmation_password"`
}

func (usv *UserSignupValidator) ValidateUserParams() (bool, string) {
	validations := []ValidationFunc{
		usv.validBlankAttributes,
		usv.validPasswordConfirmation,
		usv.userNotExists,
	}

	for _, validate := range validations {
		valid, errMsg := validate()
		if !valid {
			return false, errMsg
		}
	}

	return true, ""
}

func (usv *UserSignupValidator) validBlankAttributes() (bool, string) {
	fieldsToValidate := map[string]string{
		"Name":                 usv.User.Name,
		"Email":                usv.User.Email,
		"Password":             usv.User.Password,
		"ConfirmationPassword": usv.User.ConfirmationPassword,
	}

	for field, value := range fieldsToValidate {
		if len(value) == 0 {
			return false, "Missing Attribute: " + field + " is blank"
		}
	}
	return true, ""
}

func (usv *UserSignupValidator) validPasswordConfirmation() (bool, string) {
	if usv.User.ConfirmationPassword == usv.User.Password {
		return true, ""
	} else {
		return false, "Confirmation Password Error: password and confirmation password are not the same"
	}
}
func (usv *UserSignupValidator) userNotExists() (bool, string) {
	user := usv.UserRepository.FindByEmail(usv.User.Email)
	if len(user) >= 1 {
		return false, "Email Error: Already exists an user with this email: " + usv.User.Email
	}
	return true, ""
}
