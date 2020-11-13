package model

import (
	"time"
)

// Role :nodoc:
type Role int

// user roles
const (
	RoleAdmin   = Role(1)
	RoleStudent = Role(2)
)

// ToString :nodoc:
func (u Role) ToString() string {
	switch u {
	case RoleAdmin:
		return "ADMIN"
	case RoleStudent:
		return "STUDENT"
	default:
		return ""
	}
}

// ParseRole :nodoc:
func ParseRole(s string) Role {
	switch s {
	case "ADMIN":
		return RoleAdmin
	case "STUDENT":
		return RoleStudent
	default:
		return RoleStudent
	}
}

// User :nodoc:
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}