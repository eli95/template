package application

import (
	logger "github.com/eli95/template/pkg/logger/domain/ports"
	mdl "github.com/eli95/template/pkg/middleware/domain/ports"
	validator "github.com/eli95/template/pkg/validator/domain/ports"
)

type FiberMiddlewares struct {
	repo      mdl.MiddlewareRepository
	validator validator.ValidatorApplication
	logger    logger.LoggerApplication
}

var _ mdl.MiddlewareApplication = (*FiberMiddlewares)(nil)

func NewFiberMiddlewares(repo mdl.MiddlewareRepository, val validator.ValidatorApplication, logger logger.LoggerApplication) *FiberMiddlewares {
	return &FiberMiddlewares{
		repo:      repo,
		validator: val,
		logger:    logger,
	}
}

func (f *FiberMiddlewares) Authenticate(token string) (code int, err error) {
	//TODO implement me
	panic("implement me")
}
