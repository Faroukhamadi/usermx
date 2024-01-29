package utils

type User struct {
	Email    string
	password string
}

type Data struct {
	Users []User
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

type Page struct {
	Data Data
	Form FormData
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func NewUser(email, password string) User {
	return User{
		Email:    email,
		password: password,
	}
}

func NewData() Data {
	return Data{
		Users: []User{
			NewUser("test@test.com", "test"),
			NewUser("aze@aze.com", "aze"),
		},
	}
}

func NewPage() Page {
	return Page{
		Data: NewData(),
		Form: NewFormData(),
	}
}

// check if user exists using email
func (d Data) UserExists(email string) bool {
	for _, user := range d.Users {
		if user.Email == email {
			return true
		}
	}
	return false
}
