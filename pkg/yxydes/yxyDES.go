package yxydes

/*
#include "bridge.h"
*/
import "C"

var key = string("")

func Decrypt(s string, k string) string {
	//ip := C.Decrypt(C.CString(s), C.CString(k))
	//fmt.Print(C.GoString(ip))
	return C.GoString(C.Decrypt(C.CString(s), C.CString(k)))
}
func Encrypt(s string, k string) string {
	return C.GoString(C.Encrypt(C.CString(s), C.CString(k)))
}
func GetKey() string {
	return key
}

func SetKey(s string) string {
	key = s
	return key
}