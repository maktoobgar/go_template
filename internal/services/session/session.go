package session_service

import (
	"encoding/hex"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/errors/messages"
)

type sessionService struct{}

var errDataNotFound = errors.New(errors.NotFoundStatus, messages.ErrDataNotFound)
var errDataDuplication = errors.New(errors.NotFoundStatus, messages.ErrDataDuplication)

var instance = &sessionService{}

func (obj *sessionService) get_value(key string) (*models.Session, error) {
	sessionObj := &models.Session{}
	ok, _ := g.DB.From(models.SessionName).Limit(1).Where(
		goqu.Ex{"key": key},
	).Executor().ScanStruct(sessionObj)
	if !ok {
		return nil, errDataNotFound
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
		return nil, errDataNotFound
	}

	res, err := hex.DecodeString(session.Value)
	if err != nil {
		return nil, errors.New(errors.UnexpectedStatus, err.Error())
	}
	return res, nil
}

// Set stores the given value for the given key along with a
// time-to-live expiration value, 0 means live for ever
// Empty key or value will be ignored without an error.
func (obj *sessionService) Set(key string, val []byte, ttl time.Duration) error {
	v, _ := obj.Get(key)
	if v != nil {
		return errDataDuplication
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

	_, err := g.DB.Insert(models.SessionName).Rows(sessions).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, err.Error())
	}

	return nil
}

// Delete deletes the value for the given key.
// It returns no error if the storage does not contain the key,
func (obj *sessionService) Delete(key string) error {
	_, err := g.DB.Delete(models.SessionName).Where(goqu.Ex{
		"key": key,
	}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// Reset resets the storage and delete all keys.
func (obj *sessionService) Reset() error {
	_, err := g.DB.Delete(models.SessionName).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, err.Error())
	}

	return nil
}

// Close closes the storage and will stop any running garbage
// collectors and open connections.
func (obj *sessionService) Close() error {
	return nil
}

func New() fiber.Storage {
	return instance
}
