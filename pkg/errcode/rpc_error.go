package errcode

import (
	"blog_service_grpc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(&pb.Error{Code: int32(err.code), Message: err.Msg()})
	return s.Err()
}

type Status struct {
	*status.Status
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func ToRPCStatus(code int, msg string) *Status {
	s, _ := status.New(ToRPCCode(code), msg).WithDetails(&pb.Error{Code: int32(code), Message: msg})
	return &Status{s}
}

func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Success.Code():
		statusCode = codes.OK
	case Unknown.Code():
		statusCode = codes.Unknown
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case NotFound.Code():
		statusCode = codes.NotFound
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case Fail.Code():
		statusCode = codes.FailedPrecondition
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	default:
		statusCode = codes.Unknown
	}
	return statusCode
}
