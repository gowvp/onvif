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

// Call_SetUser forwards the call to dev.CallMethod() then parses the payload of the reply as a SetUserResponse.
func Call_SetUser(ctx context.Context, dev *onvif.Device, request device.SetUser) (device.SetUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetUserResponse device.SetUserResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetUserResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetUser")
		return reply.Body.SetUserResponse, errors.Annotate(err, "reply")
	}
}
