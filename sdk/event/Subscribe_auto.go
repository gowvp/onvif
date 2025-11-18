// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package event

import (
	"context"

	"github.com/gowvp/onvif"
	"github.com/gowvp/onvif/event"
	"github.com/gowvp/onvif/sdk"
	"github.com/juju/errors"
)

// Call_Subscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a SubscribeResponse.
func Call_Subscribe(ctx context.Context, dev *onvif.Device, request event.Subscribe) (event.SubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SubscribeResponse event.SubscribeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SubscribeResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "Subscribe")
		return reply.Body.SubscribeResponse, errors.Annotate(err, "reply")
	}
}
