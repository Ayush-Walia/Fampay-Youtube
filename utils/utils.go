package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/gookit/slog"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// CopyProperties copies all the fields from src struct to dst struct
func CopyProperties(src interface{}, dst interface{}) error {
	srcValue := reflect.ValueOf(src)
	if srcValue.Kind() != reflect.Struct {
		return fmt.Errorf("src must be a struct")
	}

	typeDst := reflect.TypeOf(dst)
	if typeDst.Kind() != reflect.Ptr {
		return fmt.Errorf("dst is not a pointer")
	}
	valDst := reflect.ValueOf(dst).Elem()
	valSrc := reflect.ValueOf(src)
	typeSrc := reflect.TypeOf(src)
	for i := 0; i < valSrc.NumField(); i++ {
		typeSrcField := typeSrc.Field(i)
		dstField := valDst.FieldByName(typeSrcField.Name)
		if !dstField.IsValid() {
			continue
		}
		if dstField.Kind() != valSrc.Field(i).Kind() {
			continue
		}
		dstField.Set(valSrc.Field(i))
	}
	return nil
}

func RespondWithJSON(w http.ResponseWriter, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		slog.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		slog.Error(err)
	}
}

func RespondWithString(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_, err := fmt.Fprint(w, msg)
	if err != nil {
		slog.Error(err)
	}
}
