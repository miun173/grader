package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/miun173/autograd/model"
	"github.com/miun173/autograd/utils"
	"github.com/sirupsen/logrus"
)

// UserRepository :nodoc:
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepository :nodoc:
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// Create :nodoc:
func (u *userRepo) Create(ctx context.Context, user *model.User) (err error) {
	err = u.db.Create(user).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":  utils.Dump(ctx),
			"user": utils.Dump(user),
		}).Error(err)
	}

	return err
}

// FindByEmail find user by username. Upon not found will return nil, nil
func (u *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("email = ?", email).Take(user).Error
	switch err {
	case nil: // ignore
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		logrus.WithFields(logrus.Fields{
			"ctx":   utils.Dump(ctx),
			"email": email,
		}).Error(err)
		return nil, err
	}

	return user, nil
}
