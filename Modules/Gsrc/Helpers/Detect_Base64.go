package Frizz_Helper

import "encoding/base64"

func VALB64(s string) bool {
	_, x := base64.StdEncoding.DecodeString(s)
	return x == nil
}
