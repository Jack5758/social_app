package request

type PostSignUpForm struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}
