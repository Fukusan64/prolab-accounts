package userstore

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/ProgrammingLab/prolab-accounts/dao"
	"github.com/ProgrammingLab/prolab-accounts/model"
	"github.com/ProgrammingLab/prolab-accounts/store"
)

type userStoreImpl struct {
	ctx context.Context
	db  *sql.DB
}

// NewUserStore returns new user store
func NewUserStore(ctx context.Context, db *sql.DB) store.UserStore {
	return &userStoreImpl{
		ctx: ctx,
		db:  db,
	}
}

func (s *userStoreImpl) GetUser(userID model.UserID) (*dao.User, error) {
	u, err := dao.Users(qm.Where("id = ?", userID)).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "")
	}

	return u, nil
}

func (s *userStoreImpl) FindUserByEmailOrName(name string) (*dao.User, error) {
	u, err := dao.Users(qm.Where("email = ? or name = ?", name, name)).One(s.ctx, s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, errors.Wrap(err, "")
	}

	return u, nil
}
