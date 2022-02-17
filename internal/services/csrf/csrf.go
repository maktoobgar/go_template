package csrf_service

import (
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

var obj *csrfService = &csrfService{}

func (c *csrfService) get_value(key string) (*models.CSRF, error) {
	scanner, err := g.DB.From(models.CSRFName).Where(
		goqu.Ex{"key": key},
	).Executor().Scanner()
	if err != nil {
		return nil, err
	}

	obj := &models.CSRF{}
	if scanner.Next() {
		err = scanner.ScanStruct(obj)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	scanner.Close()

	return obj, nil
}

// Get gets the value for the given key.
// It returns ErrNotFound if the storage does not contain the key.
func (c *csrfService) Get(key string) ([]byte, error) {
	obj, err := c.get_value(key)
	if err != nil || obj == nil {
		return nil, errDataNotFound
	}

	if obj.ExpireDate.Unix() < time.Now().Unix() {
		c.Delete(key)
		return nil, errDataNotFound
	}

	return []byte(obj.Value), nil
}

// Set stores the given value for the given key along with a
// time-to-live expiration value, 0 means live for ever
// Empty key or value will be ignored without an error.
func (c *csrfService) Set(key string, val []byte, ttl time.Duration) error {
	v, _ := c.Get(key)
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
func (c *csrfService) Delete(key string) error {
	_, err := g.DB.Delete(models.CSRFName).Where(goqu.Ex{
		"key": key,
	}).Executor().Exec()
	if err != nil {
		return err
	}

	return nil
}

// Reset resets the storage and delete all keys.
func (c *csrfService) Reset() error {
	_, err := g.DB.Delete(models.CSRFName).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, err.Error())
	}

	return nil
}

// Close closes the storage and will stop any running garbage
// collectors and open connections.
func (c *csrfService) Close() error {
	return nil
}

func New() fiber.Storage {
	return obj
}
