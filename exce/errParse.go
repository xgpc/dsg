package exce

import "google.golang.org/grpc/status"

// ParseErr 解析gPRC错误
func ParseErr(err error) {
	if err != nil {
		fromError, ok := status.FromError(err)
		if ok {
			msg := fromError.Proto().GetMessage()
			code := fromError.Proto().GetCode()
			if msg == "" {
				msg = ErrString[DsgError(code)]
			}
			ThrowSys(DsgError(code), msg)
		}
		ThrowSys(err)
	}
	return
}
