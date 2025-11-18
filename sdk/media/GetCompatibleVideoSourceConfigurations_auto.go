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

// Call_GetCompatibleVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoSourceConfigurationsResponse.
func Call_GetCompatibleVideoSourceConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleVideoSourceConfigurations) (media.GetCompatibleVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoSourceConfigurationsResponse media.GetCompatibleVideoSourceConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoSourceConfigurations")
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, errors.Annotate(err, "reply")
	}
}
