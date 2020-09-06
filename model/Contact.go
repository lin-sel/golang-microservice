package model

import (
	"errors"

	"bitbucket.com/lin-sel/golang-microservice/util"
	uuid "github.com/satori/go.uuid"
)

// Contact Contain Name, Number and ID
type Contact struct {
	ID      uuid.UUID
	Name    string
	Contact uint64
}

// MakeValidContactOrError Make Contact Valid Or Return Error
func (contact *Contact) MakeValidContactOrError() error {
	if !util.IsValidUUID(contact.ID) {
		contact.ID = util.GenerateUUID()
	}

	if util.IsEmpty(contact.Name) {
		return errors.New("Contact Name must be specified")
	}

	if contact.Contact <= 0 || contact.Contact > 9999999999 {
		return errors.New("Please Enter Valid Contact")
	}
	return nil
}
