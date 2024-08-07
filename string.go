package decompressor

// TODO: Adicionar exemplos e melhorar coverage

// ToString is a convenient function that converts a given value into a string representation.
// It uses the ToStringWithErr function with the specified type to perform the conversion. If an error occurred, it panics,
// otherwise a string is returned.
// This is a useful function when you are confident that the input is in correct format.
//
// Parameters:
//   - t: The type of decompression to be used. The type is defined by Type values (TypeGzip, TypeGzipBase64, TypeDeflate, TypeDeflateBase64).
//   - a: The interface value to be converted to a string.
//
// Returns:
//   - string: The string representation of the given interface value. If there is no error during decompression and conversion.
//
// Panic:
//   - Throws a panic if any error occurs during the conversion process.
//
// Example:
//
//	var val = "compressed data as string"
//
//	// Using TypeGzip
//	fmt.Println(ToString(TypeGzip, val)) // "Decompressed data (gzip)"
//
//	var valBase64 = "compressed data as base64 string"
//
//	// Using TypeGzipBase64
//	fmt.Println(ToString(TypeGzipBase64, valBase64)) // "Decompressed data (gzip base64)"
func ToString(t Type, a any) string {
	str, err := ToStringWithErr(t, a)
	if err != nil {
		panic(err)
	}
	return str
}

// ToStringWithErr converts a given value into a string using the specified
// decompression type. The function relies on the ToBytesWithErr function to
// perform the operation, returning any errors that occur during the conversion
// process.
//
// Parameters:
//   - t: The type of decompression to use, defined by Type values (TypeGzip,
//     TypeGzipBase64, TypeDeflate, TypeDeflateBase64).
//   - a: The value to be converted into a string.
//
// Returns:
//   - string: The string representation of the given value, if no errors
//     occurred during the decompression and conversion process.
//   - error: An error message detailing any issues that occurred during the
//     operation.
//
// Example usages:
//
//	var compressedString = "compressed data as string"
//	var compressedBase64String = "compressed data as base64 string"
//
//	// Using TypeGzip
//	str, err := ToStringWithErr(TypeGzip, compressedString)
//	if err != nil {
//		fmt.Println("Error while decompressing (gzip):", err)
//	} else {
//		fmt.Println("Decompressed and converted data (gzip):", str)
//	}
//
//	// Using TypeGzipBase64
//	str, err = ToStringWithErr(TypeGzipBase64, compressedBase64String)
//	if err != nil {
//		fmt.Println("Error while decompressing (gzip base64):", err)
//	} else {
//		fmt.Println("Decompressed and converted data (gzip base64):", str)
//	}
func ToStringWithErr(t Type, a any) (string, error) {
	bs, err := ToBytesWithErr(t, a)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
