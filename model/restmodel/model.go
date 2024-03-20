package restmodel

type RegisterUser struct {
	Username string `json:"username" gorm:"column:username;unique;not null" validate:"required"`
	Password string `json:"password" gorm:"column:password;not null" validate:"required"`
	Email    string `json:"email" gorm:"column:email" validate:"required,email"`
	Name     string `json:"name" gorm:"column:name" validate:"required"`
	Role     string `json:"role" gorm:"column:role" validate:"required"`
}

type RegisterUserResponse struct {
	Username string `json:"username"`
	Stored   bool   `json:"stored"`
}

type Login struct {
	Username string `json:"username" gorm:"column:username;unique;not null" validate:"required"`
	Password string `json:"password" gorm:"column:password;not null" validate:"required"`
}

type LoginUserResponse struct {
	LoggedIn bool   `json:"logged_in"`
	Success  string `json:"success"`
	Data     Data   `json:"data"`
}

type Data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
