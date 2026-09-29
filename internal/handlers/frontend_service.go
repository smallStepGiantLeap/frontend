package handlers

import (
	"context"

	frontendv1 "github.com/smallStepGiantLeap/frontend/client/gen/frontend/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FrontendService implements frontend.v1.FrontendService. Generated once by vikrant: this file belongs to the
// service team.
type FrontendService struct {
	frontendv1.UnimplementedFrontendServiceServer
}

// NewFrontendService returns the service implementation main registers.
func NewFrontendService() *FrontendService { return &FrontendService{} }

// Home is unary and not marked idempotent, so the mesh retries it only when
// the request never reached this server.
func (s *FrontendService) Home(ctx context.Context, req *frontendv1.HomeRequest) (*frontendv1.HomeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "frontend.v1.FrontendService/Home is not implemented yet")
}
