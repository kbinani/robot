package robot

/*
#cgo LDFLAGS: -framework IOKit
#include <IOKit/pwr_mgt/IOPMLib.h>

static void releaseCFString(CFStringRef o) {
	CFRelease(o);
}
*/
import "C"

import (
	"os/exec"
)

func pw(op PwOp) {
	switch op {
	case MonitorOn:
		// Same effect to `caffeinate -u -t 2`.
		// However, do this manually because the command is very slow.
		assertionName := cfstr("UserIsActive")
		name := cfstr("name")
		details := cfstr("details")
		reason := cfstr("reason")
		bundlePath := cfstr("/System/Library/CoreServices/powerd.bundle")
		ioPMAssertionTimeoutActionRelease := cfstr("TimeoutActionRelease")
		defer func() {
			C.releaseCFString(assertionName)
			C.releaseCFString(name)
			C.releaseCFString(details)
			C.releaseCFString(reason)
			C.releaseCFString(bundlePath)
			C.releaseCFString(ioPMAssertionTimeoutActionRelease)
		}()

		timeout := 2
		var assertionID C.IOPMAssertionID
		C.IOPMAssertionCreateWithDescription(assertionName,
			name,
			details,
			reason,
			bundlePath,
			C.CFTimeInterval(timeout),
			ioPMAssertionTimeoutActionRelease,
			&assertionID)
	case MonitorOff:
		exec.Command("/usr/bin/pmset", "displaysleepnow").Run()
	}
}
