package v1

import (
	"context"
	"errors"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	profilev1 "github.com/harmony-development/legato/gen/profile"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/responses"
	"google.golang.org/protobuf/types/known/emptypb"
)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    4,
		},
		Auth: true,
	}, "/protocol.profile.v1.ProfileService/StatusUpdate")
}

// StatusUpdate handles the protocol's StatusUpdate request
func (v1 *V1) StatusUpdate(c context.Context, r *profilev1.StatusUpdateRequest) (*empty.Empty, error) {
	ctx := c.(middleware.HarmonyContext)
	if err := r.Validate(); err != nil {
		return nil, err
	}
	if err := v1.DB.SetStatus(ctx.UserID, r.NewStatus); err != nil {
		v1.Logger.Exception(err)
		return nil, errors.New(responses.UnknownError)
	}
	return &emptypb.Empty{}, nil
}
