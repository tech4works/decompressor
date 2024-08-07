package decompressor

type Type string

const (
	TypeGzip          Type = "GZIP"
	TypeGzipBase64    Type = "GZIP_BASE64"
	TypeDeflate       Type = "DEFLATE"
	TypeDeflateBase64 Type = "DEFLATE_BASE64"
)
