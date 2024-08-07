package decompressor

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/base64"
)

type baseCase struct {
	name         string
	compressType Type
	arg          any
	want         []byte
	wantErr      bool
}

func initTestCases() []baseCase {
	strArg, _ := toBytes("Hello world")
	intArg, _ := toBytes(23)
	floatArg, _ := toBytes(23.3)
	mapArg, _ := toBytes(map[string]any{"test": 23, "test2": "de novo"})

	return []baseCase{
		{
			"Gzip String",
			TypeGzip,
			compressWithGzip(strArg),
			strArg,
			false,
		},
		{
			"Gzip Uint",
			TypeGzip,
			compressWithGzip(strArg),
			strArg,
			false,
		},
		{
			"Gzip Map",
			TypeGzip,
			compressWithGzip(mapArg),
			mapArg,
			false,
		},
		{
			"Gzip Int",
			TypeGzip,
			compressWithGzip(intArg),
			intArg,
			false,
		},
		{
			"Gzip Float",
			TypeGzip,
			compressWithGzip(floatArg),
			floatArg,
			false,
		},
		{
			"Gzip Base64 String",
			TypeGzipBase64,
			compressWithGzipBase64(strArg),
			strArg,
			false,
		},
		{
			"Gzip Base64 Map",
			TypeGzipBase64,
			compressWithGzipBase64(mapArg),
			mapArg,
			false,
		},
		{
			"Gzip Base64 Int",
			TypeGzipBase64,
			compressWithGzipBase64(intArg),
			intArg,
			false,
		},
		{
			"Gzip Base64 Float",
			TypeGzipBase64,
			compressWithGzipBase64(floatArg),
			floatArg,
			false,
		},
		{
			"Deflate String",
			TypeDeflate,
			compressWithDeflate(strArg),
			strArg,
			false,
		},
		{
			"Deflate Map",
			TypeDeflate,
			compressWithDeflate(mapArg),
			mapArg,
			false,
		},
		{
			"Deflate Int",
			TypeDeflate,
			compressWithDeflate(intArg),
			intArg,
			false,
		},
		{
			"Deflate Float",
			TypeDeflate,
			compressWithDeflate(floatArg),
			floatArg,
			false,
		},
		{
			"Deflate Base64 String",
			TypeDeflateBase64,
			compressWithDeflateBase64(strArg),
			strArg,
			false,
		},
		{
			"Deflate Base64 Map",
			TypeDeflateBase64,
			compressWithDeflateBase64(mapArg),
			mapArg,
			false,
		},
		{
			"Deflate Base64 Int",
			TypeDeflateBase64,
			compressWithDeflateBase64(intArg),
			intArg,
			false,
		},
		{
			"Deflate Base64 Float",
			TypeDeflateBase64,
			compressWithDeflateBase64(floatArg),
			floatArg,
			false,
		},
		{
			"Unsupported Type",
			"Unsupported",
			nil,
			nil,
			true,
		},
		{
			"Invalid Arg",
			"Unsupported",
			"qualquer coisa",
			nil,
			true,
		},
	}
}

func compressWithGzip(a any) []byte {
	bs, err := toBytes(a)
	if err != nil {
		panic(err)
	}

	var gzipBuffer bytes.Buffer
	gz := gzip.NewWriter(&gzipBuffer)

	_, err = gz.Write(bs)
	if err != nil {
		panic(err)
	}

	err = gz.Close()
	if err != nil {
		panic(err)
	}

	return gzipBuffer.Bytes()
}

func compressWithGzipBase64(a any) string {
	bs := compressWithGzip(a)
	return base64.StdEncoding.EncodeToString(bs)
}

func compressWithDeflate(a any) []byte {
	bs, err := toBytes(a)
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	writer, err := flate.NewWriter(&buffer, 9)
	if err != nil {
		panic(err)
	}

	_, err = writer.Write(bs)
	if err != nil {
		panic(err)
	}

	err = writer.Close()
	if err != nil {
		panic(err)
	}

	return buffer.Bytes()
}

func compressWithDeflateBase64(a any) string {
	bs := compressWithDeflate(a)
	return base64.StdEncoding.EncodeToString(bs)
}
