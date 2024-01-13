package lib

import "strconv"

const (
	gxMajor = 0
	gxMinor = 0
	gxPatch = 1
)

func Version() string {
	maj := strconv.Itoa(gxMajor)
	min := strconv.Itoa(gxMinor)
	patch := strconv.Itoa(gxPatch)
	return maj + "." + min + "." + patch
}

const (
	protoMajor = 0
	protoMinor = 14
	protoPatch = 0
)

func Protocol() string {
	maj := strconv.Itoa(protoMajor)
	min := strconv.Itoa(protoMinor)
	patch := strconv.Itoa(protoPatch)
	return maj + "." + min + "." + patch
}
