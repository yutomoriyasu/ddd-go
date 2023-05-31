package user

import "context"

type IRepository interface {
	FindAll(ctx context.Context) (*User, error)
}
