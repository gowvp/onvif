// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/media"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationsResponse.
func Call_GetVideoSourceConfigurations(ctx context.Context, dev *onvif.Device, request media.GetVideoSourceConfigurations) (media.GetVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationsResponse media.GetVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetVideoSourceConfigurations")
		return reply.Body.GetVideoSourceConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
