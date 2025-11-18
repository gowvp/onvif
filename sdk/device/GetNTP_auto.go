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

// Call_GetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNTPResponse.
func Call_GetNTP(ctx context.Context, dev *onvif.Device, request device.GetNTP) (device.GetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNTPResponse device.GetNTPResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNTPResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNTP")
		return reply.Body.GetNTPResponse, errors.Annotate(err, "reply")
	}
}
