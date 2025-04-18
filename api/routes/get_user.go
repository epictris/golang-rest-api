package routes

type args struct {
	UserId int `validate:"required"`
}

type response struct {
	Name string `validate:"required"`
}

func GetUser(args args) (response, error) {
	response := response{}
	response.Name = "John Doe"
	return response, nil
}
