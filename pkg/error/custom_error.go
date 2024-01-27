package custom_error

import (
	"net/http"
	"runtime"

	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
)

type CustomErrors struct {
	ErrorContext *ErrorContext
}

type ErrorContext struct {
	Code     int
	IsIgnore bool
	Message  string
	HTTPCode int
	Func     string
}

func (c *CustomErrors) Error() string {
	if c.ErrorContext.HTTPCode == 0 {
		c.ErrorContext.HTTPCode = http.StatusInternalServerError
	}

	return constant.ErrorMessageMap[constant.DefaultUnhandledError]
}

func (c *CustomErrors) IsIgnore() bool {
	return c.ErrorContext.IsIgnore
}

func SetCustomError(contextError *ErrorContext) *CustomErrors {
	contextError.Func = getCallerFunctionName()
	return &CustomErrors{
		ErrorContext: contextError,
	}
}

func getCallerFunctionName() string {
	// Skip GetCallerFunctionName and the function to get the caller of
	//nolint:gomnd
	return getFrame(2).Function
}

//nolint:gomnd
func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
