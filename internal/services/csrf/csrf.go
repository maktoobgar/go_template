package csrf_service

import (
	"strings"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
)

type csrfService struct{}

var errDataNotFound = errors.New(errors.NotFoundStatus, "requested data does not exist")
var errDataDuplication = errors.New(errors.NotFoundStatus, "the same data exist")

var instance *csrfService = &csrfService{}

func (obj *csrfService) get_value(key string) (*models.CSRF, error) {
	csrfObj := &models.CSRF{}
	ok, _ := g.DB.From(models.CSRFName).Limit(1).Where(
		goqu.Ex{"key": key},
	).Executor().ScanStruct(csrfObj)
	if !ok {
		return nil, errDataNotFound
	}

	return csrfObj, nil
}

// Get gets the value for the given key.
// It returns ErrNotFound if the storage does not contain the key.
func (obj *csrfService) Get(key string) ([]byte, error) {
	csrfObj, err := obj.get_value(key)
	if err != nil {
		return nil, errDataNotFound
	}

	if csrfObj.ExpireDate.Unix() < time.Now().Unix() {
		obj.Delete(key)
		return nil, errDataNotFound
	}

	return []byte(csrfObj.Value), nil
}

// Set stores the given value for the given key along with a
// time-to-live expiration value, 0 means live for ever
// Empty key or value will be ignored without an error.
func (obj *csrfService) Set(key string, val []byte, ttl time.Duration) error {
	v, _ := obj.Get(key)
	if v != nil {
		return errDataDuplication
	}
	if key == "" || val == nil {
		return nil
	}

	t := time.Now().Add(ttl)
	csrfs := []models.CSRF{
		{
			Key:        key,
			Value:      string(val),
			ExpireDate: t,
		},
	}
	_, err := g.DB.Insert(models.CSRFName).Rows(csrfs).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, err.Error())
	}

	return nil
}

// Delete deletes the value for the given key.
// It returns no error if the storage does not contain the key,
func (obj *csrfService) Delete(key string) error {
	_, err := g.DB.Delete(models.CSRFName).Where(goqu.Ex{
		"key": key,
	}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// Reset resets the storage and delete all keys.
func (obj *csrfService) Reset() error {
	_, err := g.DB.Delete(models.CSRFName).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, err.Error())
	}

	return nil
}

// Close closes the storage and will stop any running garbage
// collectors and open connections.
func (obj *csrfService) Close() error {
	return nil
}

func Next(c *fiber.Ctx) bool {
	return strings.Contains(c.Path(), "/api/")
}

func New() fiber.Storage {
	return instance
}
