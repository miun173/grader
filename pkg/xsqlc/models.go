// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package xsqlc

import (
	"database/sql"
	"time"
)

type ActivationToken struct {
	ID        string
	Token     string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Assignment struct {
	ID               string
	AssignedBy       string
	Name             string
	Description      string
	CaseOutputFileID string
	CaseInputFileID  string
	CreatedAt        time.Time
	DeadlineAt       time.Time
	UpdatedAt        time.Time
	DeletedAt        sql.NullTime
	Template         string
}

type Course struct {
	ID          string
	Name        string
	Description string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type File struct {
	ID        string
	Name      string
	Type      string
	Ext       string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type OutboxItem struct {
	ID            string
	IdempotentKey string
	Status        string
	JobType       string
	Payload       string
	Version       int32
}

type RelAssignmentToCourse struct {
	CourseID     string
	AssignmentID string
}

type RelCourseUser struct {
	CourseID string
	UserID   string
	UserType string
}

type RelUserToActivationToken struct {
	UserID            string
	ActivationTokenID string
	DeletedAt         sql.NullTime
}

type Submission struct {
	ID           string
	AssignmentID string
	SubmittedBy  string
	IsGraded     int32
	Grade        int32
	Feedback     string
	FileID       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      string
	Active    int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
