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

// Call_StartFirmwareUpgrade forwards the call to dev.CallMethod() then parses the payload of the reply as a StartFirmwareUpgradeResponse.
func Call_StartFirmwareUpgrade(ctx context.Context, dev *onvif.Device, request device.StartFirmwareUpgrade) (device.StartFirmwareUpgradeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartFirmwareUpgradeResponse device.StartFirmwareUpgradeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StartFirmwareUpgradeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "StartFirmwareUpgrade")
		return reply.Body.StartFirmwareUpgradeResponse, errors.Annotate(err, "reply")
	}
}
