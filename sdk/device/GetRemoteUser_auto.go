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

// Call_GetRemoteUser forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRemoteUserResponse.
func Call_GetRemoteUser(ctx context.Context, dev *onvif.Device, request device.GetRemoteUser) (device.GetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteUserResponse device.GetRemoteUserResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRemoteUserResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRemoteUser")
		return reply.Body.GetRemoteUserResponse, errors.Annotate(err, "reply")
	}
}
