package shorturl

import "github.com/spaolacci/murmur3"

// ShortUrlWithSeed 指定种子生成短网址
func ShortUrlWithSeed(url string, seed uint32) (string, error) {
	n, err := GetUrlIntValue(url, seed)
	if err != nil {
		return "", err
	}
	return uint32To62(n), nil
}

// ShortUrl 生成短网址
func ShortUrl(url string) (string, error) {
	return ShortUrlWithSeed(url, 0)
}

// getUrlIntValue 获取原始网址的整数值
func GetUrlIntValue(url string, seed uint32) (n uint32, err error) {
	h := murmur3.New32WithSeed(seed)
	_, err = h.Write([]byte(url))
	if err != nil {
		return 0, err
	}
	n = h.Sum32()
	return
}

// uint32To62 将一个 uint32 数转换成六十二进制的文本表示
func uint32To62(n uint32) string {
	dict := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	buf := make([]byte, 0, 6)
	for n > 0 {
		idx := n % 62
		buf = append(buf, dict[idx])
		n /= 62
	}
	reverse(buf)
	return string(buf)
}

// reverse 反转切片
func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
