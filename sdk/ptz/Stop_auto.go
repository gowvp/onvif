// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/ptz"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_Stop forwards the call to dev.CallMethod() then parses the payload of the reply as a StopResponse.
func Call_Stop(ctx context.Context, dev *onvif.Device, request ptz.Stop) (ptz.StopResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopResponse ptz.StopResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.StopResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Stop")
		return reply.Body.StopResponse, errors.Annotate(err, "reply")
	}
}
