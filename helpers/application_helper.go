package helpers

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func ParseUUID(idStr string) uuid.UUID {
	id, err := uuid.Parse(idStr)

	if err != nil {
		return id
	} else {
		logrus.WithFields(logrus.Fields{"id": idStr}).Warn("Cannot Convert uuid")
		return uuid.Nil
	}
}
