package guitars

import (
	"context"
)

type Service interface {
	ListGuitars(ctx context.Context) (error)
}

type svc struct {

}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListGuitars(ctx context.Context) error {
	return nil
}
