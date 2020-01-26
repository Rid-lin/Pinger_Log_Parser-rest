package checkmethods

import "runtime"

//Ping ...
func Ping(ip string) string {
	os := runtime.GOOS
	switch os {
	case "windows":
		return pingWin(ip)
	case "linux":
		return pingUnix(ip)
	case "darwin":
		return pingMac(ip)
	}
	return pingWin(ip)
}