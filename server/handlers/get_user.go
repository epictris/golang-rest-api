package handlers

import (
	"context"

	"connectrpc.com/connect"

	apiv1 "github.com/epictris/go/gen/proto/v1"
)

func (s *Server) GetUser(
	ctx context.Context,
	req *connect.Request[apiv1.GetUserRequest],
) (*connect.Response[apiv1.GetUserResponse], error) {
	getUserSQL := `
		SELECT email FROM users WHERE id = ?
	`
	var email string

	err := s.DB.QueryRow(getUserSQL, req.Msg.UserId).Scan(&email)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&apiv1.GetUserResponse{Email: email}), nil
}
