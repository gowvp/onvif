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

// Call_GetSystemBackup forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemBackupResponse.
func Call_GetSystemBackup(ctx context.Context, dev *onvif.Device, request device.GetSystemBackup) (device.GetSystemBackupResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemBackupResponse device.GetSystemBackupResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetSystemBackupResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetSystemBackup")
		return reply.Body.GetSystemBackupResponse, errors.Annotate(err, "reply")
	}
}
