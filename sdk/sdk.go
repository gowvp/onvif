package sdk

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"log/slog"
	"net/http"

	"github.com/juju/errors"
)

func ReadAndParse(ctx context.Context, httpReply *http.Response, reply interface{}, tag string) error {
	slog.Debug("RPC", "msg", httpReply.Status, "status", httpReply.StatusCode, "action", tag)
	// TODO(jfsmig): extract the deadline from ctx.Deadline() and apply it on the reply reading
	b, err := ioutil.ReadAll(httpReply.Body)
	if err != nil {
		return errors.Annotate(err, "read")
	}

	httpReply.Body.Close()

	err = xml.Unmarshal(b, reply)
	return errors.Annotate(err, "decode")
}
