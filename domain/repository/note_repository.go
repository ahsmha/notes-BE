//go:generate mockgen -source=$GOFILE -package=mocks -destination=mocks/note_repository.go
package repository

import "ahsmha/notes/domain/model"

type NoteRepository interface {
	GetById(id int) (*model.Note, error)
	Create(note *model.Note) (int64, error)
	Update(note *model.Note) error
	Delete(id int) error
}
