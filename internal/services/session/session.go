package session_service

import (
	"encoding/hex"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
)

type sessionService struct{}

var (
	instance                = &sessionService{}
	db       *goqu.Database = nil
)

func (obj *sessionService) get_value(key string) (*models.Session, error) {
	sessionObj := &models.Session{}
	ok, _ := db.From(models.SessionName).Limit(1).Where(
		goqu.Ex{"key": key},
	).Executor().ScanStruct(sessionObj)
	if !ok {
		return nil, errors.New(errors.UnauthorizedStatus, errors.Resend, g.Trans().TranslateEN("InvalidSessionID"))
	}

	return sessionObj, nil
}

// Get gets the value for the given key.
// It returns ErrNotFound if the storage does not contain the key.
func (obj *sessionService) Get(key string) ([]byte, error) {
	session, err := obj.get_value(key)
	if err != nil {
		return nil, err
	}

	if session.ExpireDate.Unix() < time.Now().Unix() {
		obj.Delete(key)
		return nil, errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("ExpiredSessionID"))
	}

	res, err := hex.DecodeString(session.Value)
	if err != nil {
		return nil, errors.New(errors.UnexpectedStatus, errors.ReSingIn, g.Trans().TranslateEN("DecodeFailure"))
	}
	return res, nil
}

// Set stores the given value for the given key along with a
// time-to-live expiration value, 0 means live for ever
// Empty key or value will be ignored without an error.
func (obj *sessionService) Set(key string, val []byte, ttl time.Duration) error {
	v, _ := obj.Get(key)
	if v != nil {
		return errors.New(errors.InvalidStatus, errors.Resend, g.Trans().TranslateEN("DuplicateSession"))
	}
	if key == "" || val == nil {
		return nil
	}

	t := time.Now().Add(ttl)
	sessions := []models.Session{
		{
			Key:        key,
			Value:      hex.EncodeToString(val),
			ExpireDate: t,
		},
	}

	_, err := db.Insert(models.SessionName).Rows(sessions).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, errors.Report, g.Trans().TranslateEN("CreationSessionFailed"))
	}

	return nil
}

// Delete deletes the value for the given key.
// It returns no error if the storage does not contain the key,
func (obj *sessionService) Delete(key string) error {
	_, err := db.Delete(models.SessionName).Where(goqu.Ex{
		"key": key,
	}).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, errors.Report, g.Trans().TranslateEN("DeletionSessionFailed"))
	}

	return nil
}

// Reset resets the storage and delete all keys.
func (obj *sessionService) Reset() error {
	_, err := db.Delete(models.SessionName).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, errors.Report, g.Trans().TranslateEN("ResetSessionFailed"))
	}

	return nil
}

// Close closes the storage and will stop any running garbage
// collectors and open connections.
func (obj *sessionService) Close() error {
	return nil
}

func SetDB(entryDB ...*goqu.Database) {
	if entryDB != nil {
		db = entryDB[0]
	} else {
		db = g.DB
	}
}

func New() fiber.Storage {
	return instance
}
