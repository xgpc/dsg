package exce

import (
	"google.golang.org/grpc/status"
)

func ParseErr(err error) (int32, string) {
	if err != nil {
		fromError, ok := status.FromError(err)
		if ok {
			msg := fromError.Proto().GetMessage()
			code := fromError.Proto().GetCode()
			if msg == "" {
				msg = ErrString[int(code)]
			}

			return code, msg
		}
	}
	return 0, ""
}
