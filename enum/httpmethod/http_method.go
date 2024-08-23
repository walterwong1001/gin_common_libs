package httpmethod

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	OPTIONS HTTPMethod = "OPTIONS"
	HEAD    HTTPMethod = "HEAD"
)

// IsValid 检查 HTTP 方法是否有效
func (m HTTPMethod) IsValid() bool {
	switch m {
	case GET, POST, PUT, DELETE, OPTIONS, HEAD:
		return true
	default:
		return false
	}
}
