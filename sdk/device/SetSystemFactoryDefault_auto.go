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

// Call_SetSystemFactoryDefault forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemFactoryDefaultResponse.
func Call_SetSystemFactoryDefault(ctx context.Context, dev *onvif.Device, request device.SetSystemFactoryDefault) (device.SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse device.SetSystemFactoryDefaultResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetSystemFactoryDefault")
		return reply.Body.SetSystemFactoryDefaultResponse, errors.Annotate(err, "reply")
	}
}
