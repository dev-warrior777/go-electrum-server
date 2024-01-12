package lib

import "strconv"

const (
	gxMajorVersion = 0
	gxMinorVersion = 0
	gxPatch        = 1
)

func Version() string {
	maj := strconv.Itoa(gxMajorVersion)
	min := strconv.Itoa(gxMinorVersion)
	patch := strconv.Itoa(gxPatch)
	return maj + "." + min + "." + patch
}
