package handlers

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/epictris/go/gen/proto/v1"
)

func (s *Server) CreateUser(
	ctx context.Context,
	req *connect.Request[apiv1.CreateUserRequest],
) (*connect.Response[apiv1.CreateUserResponse], error) {
	createUserSQL := `
		INSERT INTO users (email)
		VALUES (?)
		ON CONFLICT DO NOTHING;
	`
	_, err := s.DB.Exec(createUserSQL, req.Msg.Email)
	if err != nil {
		return nil, err
	}

	getUserSQL := `
		SELECT id FROM users WHERE email = ?
	`
	var userId int64

	err = s.DB.QueryRow(getUserSQL, req.Msg.Email).Scan(&userId)
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&apiv1.CreateUserResponse{UserId: userId}), nil
}
