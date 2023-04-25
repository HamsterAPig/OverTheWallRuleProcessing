package ContentParse

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// GetUrlRespAndDecode 接受一个url，获取其内容，然后返回其相应
func GetUrlResp(url string) (string, error) {
	// 使用 http.Get 来获取url的内容
	resp, err := http.Get(url)
	if err != nil {
		// 如果获取失败，返回错误
		return "", err
	}
	// 延迟关闭响应体
	defer resp.Body.Close()
	// 读取响应体为字节
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// 如果读取失败，返回错误
		return "", err
	}
	return string(body), nil
}

// SplitAndDecode 接受一个文本，按照换行符分割，然后对每一部分进行URL解码
func SplitAndDecode(text string) ([]string, error) {
	// 按照换行符分割文本
	parts := strings.Split(text, "\n")
	// 创建一个切片来存储解码后的部分
	decoded := make([]string, len(parts))
	// 遍历每一部分，对其进行URL解码
	for i, part := range parts {
		// 使用 url.QueryUnescape 来解码部分
		unescaped, err := url.QueryUnescape(part)
		if err != nil {
			// 如果解码失败，返回错误
			return nil, err
		}
		// 把解码后的部分存储在切片中
		decoded[i] = unescaped
	}
	// 返回解码后的切片
	return decoded, nil
}

// ParseB64String base64解码
func ParseB64String(b64String string) ([]byte, error) {
	missingPadding := len(b64String) % 4
	if missingPadding != 0 {
		b64String = b64String + strings.Repeat("=", missingPadding)
	}
	decodedBytes, err := base64.RawURLEncoding.DecodeString(b64String)
	if err != nil {
		decodedBytes, err = base64.URLEncoding.DecodeString(b64String)
		if err != nil {
			decodedBytes, err = base64.StdEncoding.DecodeString(b64String)
			if err != nil {
				log.Println("decode base64 fail:", err)
				return []byte{}, err
			}
		}
	}
	return decodedBytes, nil
}
