package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/walterwong1001/gin_common_libs/access_log"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLog(logger accesslog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		var requestBody []byte
		if ctx.Request.Body != nil {
			// 读取请求体
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			// 重新设置请求体，以便后续处理
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		ctx.Next()

		// 计算响应时间
		latency := time.Since(startTime).Milliseconds()
		userId, _ := ctx.Get("CURRENT_USER_ID")

		// 生成日志记录
		metrics := map[string]any{
			"timestamp":            time.Now().UnixMilli(),
			"remote_addr":          ctx.ClientIP(),
			"remote_user":          userId,
			"request_method":       ctx.Request.Method,
			"request_uri":          ctx.Request.RequestURI,
			"server_protocol":      ctx.Request.Proto,
			"status":               ctx.Writer.Status(),
			"body_bytes_sent":      ctx.Writer.Size(),
			"http_referer":         ctx.Request.Referer(),
			"http_user_agent":      ctx.Request.UserAgent(),
			"http_x_forwarded_for": ctx.GetHeader("X-Forwarded-For"),
			"request_time":         latency,
		}

		if len(ctx.Errors) > 0 {
			metrics["error"] = ctx.Errors.Last().Err.Error()
		}

		// 提取文件元数据（如果有）
		var fileMetadata string
		if isFileUpload(ctx) {
			f, err := ctx.FormFile("file")
			if err == nil {
				fileMetadata = fmt.Sprintf(`{"file_name":"%s","file_size":%d,"file_type":"%s"}`, f.Filename, f.Size, f.Header.Get("Content-Type"))
			}
			metrics["file_metadata"] = fileMetadata
		} else {
			metrics["request_body"] = compressBody(requestBody)
		}

		_ = logger.Log(ctx, metrics)
	}
}

// 判断请求是否为文件上传
func isFileUpload(ctx *gin.Context) bool {
	return ctx.Request.Header.Get("Content-Type") == "multipart/form-data"
}

// compressBody attempts to compress the body by trying to compact it if it's valid JSON.
// If it's not valid JSON, it returns the original body as a string.
func compressBody(body []byte) string {
	// 尝试将请求体解析为JSON并压缩
	var buf bytes.Buffer
	if err := json.Compact(&buf, body); err == nil {
		return buf.String() // 如果是JSON且压缩成功，返回压缩后的JSON
	}
	return string(body) // 如果不是JSON或压缩失败，返回原始的请求体
}
