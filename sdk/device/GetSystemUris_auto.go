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

// Call_GetSystemUris forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemUrisResponse.
func Call_GetSystemUris(ctx context.Context, dev *onvif.Device, request device.GetSystemUris) (device.GetSystemUrisResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemUrisResponse device.GetSystemUrisResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemUrisResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSystemUris")
		return reply.Body.GetSystemUrisResponse, errors.Annotate(err, "reply")
	}
}
