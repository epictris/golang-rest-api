package routes

import "database/sql"

type getUserArgs struct {
	UserId string `json:"userId" validate:"required"`
}

type getUserResponse struct {
	Email string `json:"email" validate:"required"`
}

func GetUser(session *sql.DB, args getUserArgs) (getUserResponse, error) {
	getUserSQL := `
		SELECT email FROM users WHERE id = ?
	`
	var email string

	err := session.QueryRow(getUserSQL, args.UserId).Scan(&email)
	if err != nil {
		return getUserResponse{}, err
	}

	return getUserResponse{Email: email}, nil
}
