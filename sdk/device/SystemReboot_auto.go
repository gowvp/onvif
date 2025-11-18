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

// Call_SystemReboot forwards the call to dev.CallMethod() then parses the payload of the reply as a SystemRebootResponse.
func Call_SystemReboot(ctx context.Context, dev *onvif.Device, request device.SystemReboot) (device.SystemRebootResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SystemRebootResponse device.SystemRebootResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SystemRebootResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SystemReboot")
		return reply.Body.SystemRebootResponse, errors.Annotate(err, "reply")
	}
}
