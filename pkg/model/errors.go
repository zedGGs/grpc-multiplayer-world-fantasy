package model

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrName = errors.New("name contains not valid character(s)")

	ErrNameProfane = errors.New("name unavailable")

	ErrRealm = errors.New("not valid realm")

	ErrGender = errors.New("not valid gender")

	ErrServerLocation = errors.New("not valid server location")

	ErrUnauthorized  = status.Error(codes.Unauthenticated, "not authorized")
	ErrDoesNotExist  = status.Error(codes.NotFound, "does not exist")
	ErrHandleRequest = status.Error(codes.Internal, "unable to handle request")
)
