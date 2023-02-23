package repositories

import (
	"context"
	"fmt"
	"reflect"

	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

func structCheck(data any, ctx context.Context) (reflect.Type, reflect.Value) {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	for dataType.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
		dataType = dataValue.Type()
	}
	if dataType.Kind() != reflect.Struct {
		panic(errors.New(errors.UnexpectedStatus, errors.Report, translate("InternalServerError")))
	}
	return dataType, dataValue
}

func formatValue(input any) string {
	switch input.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return fmt.Sprint(input)
	default:
		return fmt.Sprintf("'%s'", input)
	}
}

func InsertInto(table string, data any, ctx context.Context) string {
	dataType, dataValue := structCheck(data, ctx)
	query := fmt.Sprintf("INSERT INTO %s ", table)
	keys := ""
	values := ""
	for _, f := range reflect.VisibleFields(dataType) {
		if f.IsExported() {
			name := f.Name
			value := dataValue.FieldByName(name).Interface()
			if value != nil {
				if keys == "" {
					keys += name
					values += formatValue(value)
					continue
				}
				keys += ", " + name
				values += ", " + formatValue(value)
			}
		}
	}
	return query + fmt.Sprintf("(%s) VALUES(%s);", keys, values)
}

func Select(table string, keyValues map[string]any, ctx context.Context) string {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	wheres := ""
	for key, value := range keyValues {
		if wheres == "" {
			wheres = fmt.Sprintf("%s = %s", key, formatValue(value))
			continue
		}
		wheres += " AND " + fmt.Sprintf("%s = %s", key, formatValue(value))
	}
	if wheres == "" {
		panic(errors.New(errors.ForbiddenStatus, errors.DoNothing, translate("RetrievingAllTableNotAllowed")))
	}
	return fmt.Sprintf("SELECT DISTINCT * FROM %s WHERE %s;", table, wheres)
}

func Delete(table string, keyValues map[string]any, ctx context.Context) string {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	wheres := ""
	for key, value := range keyValues {
		if wheres == "" {
			wheres = fmt.Sprintf("%s = %s", key, formatValue(value))
			continue
		}
		wheres += " AND " + fmt.Sprintf("%s = %s", key, formatValue(value))
	}
	if wheres == "" {
		panic(errors.New(errors.ForbiddenStatus, errors.DoNothing, translate("RemovingAllTableNotAllowed")))
	}
	return fmt.Sprintf("DELETE FROM %s WHERE %s;", table, wheres)
}
