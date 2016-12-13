package http_handlers

type error_t struct {
	Message		string		`json:"message"`
}

type user_create_form struct {
	Name		string		`json:"name" binding:"required,alphanum"`
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"password" binding:"required"`
}

type user_update_form struct {
	Name		string		`json:"name" binding:"required,alphanum"`
	Email		string		`json:"email" binding:"required,email"`
	Password	string		`json:"password" binding:"required"`
	NewPassword	string		`json:"new_password"`
}

type user_login_form struct {
	Name		string		`json:"name" binding:"required,alphanum"`
	Password	string		`json:"password" binding:"required"`
}
