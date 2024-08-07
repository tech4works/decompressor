package decompressor

import "errors"

// TODO: Adicionar exemplos e melhorar coverage

// ToBytes converts a given value into a byte slice using the specified type.
// It uses the ToBytesWithErr function to perform the type conversion and decompression
// but unlike ToBytesWithErr, this function will panic if an error occurred during the process.
//
// Parameters:
//   - t: The type of decompression to use, defined by a Type value (TypeGzip, TypeGzipBase64, TypeDeflate, TypeDeflateBase64).
//   - a: The interface value to be converted to a byte slice.
//
// Returns:
//   - []byte: A byte slice which is the result of decompressing the input value.
//
// Example:
//
//	var data = "hello world"
//
//	// Using TypeGzip
//	bs :=  ToBytes(TypeGzip, data)
//	fmt.Println("Decompressed data (gzip): ", string(bs))
//
//	// Using TypeGzipBase64
//	bs = ToBytes(TypeGzipBase64, data)
//	fmt.Println("Decompressed data (gzip base64): ", string(bs))
//
// Panic:
// The function will panic if any error occurred during the operation.
// For instance, if used with an unsupported Type value, it will cause a panic.
// Therefore, it's recommended to use it within a try catch if the types are not certain.
//
//	// This will cause a panic
//	defer func() {
//	if r := recover(); r != nil {
//		fmt.Println("Recovered from ", r)
//	}
//	}()
//	bs = ToBytes(Type("Unsupported"), data)
func ToBytes(t Type, a any) []byte {
	bs, err := ToBytesWithErr(t, a)
	if err != nil {
		panic(err)
	}
	return bs
}

// ToBytesWithErr converts a given value into a byte slice using the specified
// type. The function returns an error if the compression type is unsupported
// or if an error occurred during decompression.
//
// Parameters:
//   - t: The type of decompression to use, defined by a Type values (TypeGzip, TypeGzipBase64, TypeDeflate, TypeDeflateBase64).
//   - a: The interface value to be converted to a byte slice.
//
// Returns:
//   - []byte: A byte slice which is the result of decompressing the input value (if no error occurred)
//   - error: An error message detailing any error that occurred during the operation
//
// Example usages:
//
//	// Using TypeGzip
//	bs, err :=  ToBytesWithErr(TypeGzip, "compressed data as string")
//	if err != nil {
//	   fmt.Println("Error while decompressing (gzip): ", err)
//	} else {
//	   fmt.Println("Decompressed data (gzip): ", string(bs))
//	}
//
//	// Using TypeGzipBase64
//	bs, err  = ToBytesWithErr(TypeGzipBase64, "compressed data as base64 string")
//	if err != nil {
//	   fmt.Println("Error while decompressing (gzip base64): ", err)
//	} else {
//	   fmt.Println("Decompressed data (gzip base64): ", string(bs))
//	}
func ToBytesWithErr(t Type, a any) ([]byte, error) {
	switch t {
	case TypeGzip:
		return decompressWithGzip(a)
	case TypeGzipBase64:
		return decompressWithGzipBase64(a)
	case TypeDeflate:
		return decompressWithDeflate(a)
	case TypeDeflateBase64:
		return decompressWithDeflateBase64(a)
	default:
		return nil, errors.New("unsupported compress type")
	}
}
