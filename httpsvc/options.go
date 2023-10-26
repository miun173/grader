package httpsvc

import (
	"github.com/fahmifan/autograd/model"
	"gorm.io/gorm"
)

// Option ..
type Option func(*Server)

// WithUserUsecase ..
func WithUserUsecase(u model.UserUsecase) Option {
	return func(s *Server) {
		s.userUsecase = u
	}
}

func WithGormDB(db *gorm.DB) Option {
	return func(s *Server) {
		s.gormDB = db
	}
}

// WithAssignmentUsecase ..
func WithAssignmentUsecase(a model.AssignmentUsecase) Option {
	return func(s *Server) {
		s.assignmentUsecase = a
	}
}

// WithSubmissionUsecase ..
func WithSubmissionUsecase(sub model.SubmissionUsecase) Option {
	return func(s *Server) {
		s.submissionUsecase = sub
	}
}

// WithMediaUsecase ..
func WithMediaUsecase(med model.MediaUsecase) Option {
	return func(s *Server) {
		s.mediaUsecase = med
	}
}
