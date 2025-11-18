// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/ptz"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationOptionsResponse.
func Call_GetConfigurationOptions(ctx context.Context, dev *onvif.Device, request ptz.GetConfigurationOptions) (ptz.GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse ptz.GetConfigurationOptionsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetConfigurationOptionsResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetConfigurationOptions")
		return reply.Body.GetConfigurationOptionsResponse, errors.Annotate(err, "reply")
	}
}
