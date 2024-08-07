package decompressor

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
)

func decodeBase64(a any) ([]byte, error) {
	bs, err := toBytes(a)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(string(bs))
}

func decompressWithGzipBase64(a any) ([]byte, error) {
	bs, err := decodeBase64(a)
	if err != nil {
		return nil, err
	}

	return decompressWithGzip(bs)
}

func decompressWithDeflateBase64(a any) ([]byte, error) {
	bs, err := decodeBase64(a)
	if err != nil {
		return nil, err
	}

	return decompressWithDeflate(bs)
}

func decompressWithGzip(a any) ([]byte, error) {
	bs, err := toBytes(a)
	if err != nil {
		return nil, err
	}

	reader, err := gzip.NewReader(bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	result, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = reader.Close()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func decompressWithDeflate(a any) ([]byte, error) {
	bs, err := toBytes(a)
	if err != nil {
		return nil, err
	}

	reader := flate.NewReader(bytes.NewReader(bs))

	result, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = reader.Close()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func toBytes(a any) ([]byte, error) {
	reflectValue := reflect.ValueOf(a)

	switch reflectValue.Kind() {
	case reflect.String:
		return []byte(reflectValue.String()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(reflectValue.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return []byte(strconv.FormatUint(reflectValue.Uint(), 10)), nil
	case reflect.Float32, reflect.Float64:
		return []byte(strconv.FormatFloat(reflectValue.Float(), 'g', -1, 64)), nil
	case reflect.Complex64, reflect.Complex128:
		return []byte(strconv.FormatComplex(reflectValue.Complex(), 'g', -1, 64)), nil
	case reflect.Bool:
		return []byte(strconv.FormatBool(reflectValue.Bool())), nil
	case reflect.Array, reflect.Slice:
		if reflectValue.Type().Elem().Kind() == reflect.Uint8 {
			return reflectValue.Bytes(), nil
		}
		marshal, _ := json.Marshal(reflectValue.Interface())
		return marshal, nil
	case reflect.Map, reflect.Struct:
		marshal, _ := json.Marshal(reflectValue.Interface())
		return marshal, nil
	case reflect.Ptr, reflect.Interface:
		if reflectValue.IsNil() {
			return nil, errors.New("error convert to bytes, it is null")
		}
		return toBytes(reflectValue.Elem().Interface())
	default:
		return nil, fmt.Errorf("error convert to bytes, unsupported type %s", reflectValue.Kind().String())
	}
}
