package user

import "context"

type IRepository interface {
	FindAll(ctx context.Context) (Users, error)
}
