// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/device"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_GetZeroConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetZeroConfigurationResponse.
func Call_GetZeroConfiguration(ctx context.Context, dev *onvif.Device, request device.GetZeroConfiguration) (device.GetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetZeroConfigurationResponse device.GetZeroConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetZeroConfigurationResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetZeroConfiguration")
		return reply.Body.GetZeroConfigurationResponse, errors.Annotate(err, "reply")
	}
}
