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

// Call_PullMessages forwards the call to dev.CallMethod() then parses the payload of the reply as a PullMessagesResponse.
func Call_PullMessages(ctx context.Context, dev *onvif.Device, request event.PullMessages) (event.PullMessagesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			PullMessagesResponse event.PullMessagesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.PullMessagesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "PullMessages")
		return reply.Body.PullMessagesResponse, errors.Annotate(err, "reply")
	}
}
