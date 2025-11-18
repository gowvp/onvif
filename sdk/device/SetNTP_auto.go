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

// Call_SetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNTPResponse.
func Call_SetNTP(ctx context.Context, dev *onvif.Device, request device.SetNTP) (device.SetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNTPResponse device.SetNTPResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNTPResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetNTP")
		return reply.Body.SetNTPResponse, errors.Annotate(err, "reply")
	}
}
