package Frizz_Helper

import "encoding/base64"

func DECB64(s string) string {
	f, _ := base64.StdEncoding.DecodeString(s)
	return string(f)
}
