package infra

import (
	"github.com/sota0121/janusly/server/internal/domain/repository"
)

type DynamoDBURLSetRepository struct {
}

var _ repository.URLSetRepository = (*DynamoDBURLSetRepository)(nil)
