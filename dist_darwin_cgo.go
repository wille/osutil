//+build darwin,cgo

package osutil

/*
#include <sys/sysctl.h>
*/
import "C"

import "unsafe"

func getKernelRelease() string {
	str := C.CString("                                                                                                          ")
	size := 256

	C.sysctlbyname(C.CString("kern.osrelease"), unsafe.Pointer(str), (*C.size_t)(unsafe.Pointer(&size)), nil, 0)

	return C.GoString(str)
}
