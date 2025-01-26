package contract

import "task-tracker-cli/entity"

type UserWriteStore interface {
	Save(u entity.User)
}

type UserReadStore interface {
	Load(serialiazatinMode string) []entity.User
}