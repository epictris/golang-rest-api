package routes

import "database/sql"

type createUserArgs struct {
	Email string `json:"email" validate:"required"`
}

type createUserResponse struct {
	UserId int `json:"userId" validate:"required"`
}

func CreateUser(session *sql.DB, args createUserArgs) (createUserResponse, error) {
	createUserSQL := `
		INSERT INTO users (email)
		VALUES (?)
		ON CONFLICT DO NOTHING;
	`
	_, err := session.Exec(createUserSQL, args.Email)
	if err != nil {
		return createUserResponse{}, err
	}

	getUserSQL := `
		SELECT id FROM users WHERE email = ?
	`
	var userId int

	err = session.QueryRow(getUserSQL, args.Email).Scan(&userId)
	if err != nil {
		return createUserResponse{}, err
	}

	return createUserResponse{UserId: userId}, nil
}
