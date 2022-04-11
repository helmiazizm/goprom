package model

import "time"

type Student struct {
	Id        int
	Name      string
	Age       int
	JoinDate  time.Time
	IdCard    string
	Senior    bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
