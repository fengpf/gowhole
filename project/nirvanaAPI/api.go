package nirvanaAPI

import (
	"context"

	"github.com/caicloud/nirvana/definition"
)

//EchoDesc for test api.
var EchoDesc = definition.Descriptor{
	Path:        "/echo",
	Description: "Echo API",
	Definitions: []definition.Definition{
		{
			Method:   definition.Get,
			Function: EchoMsg,
			Consumes: []string{definition.MIMEAll},
			Produces: []string{definition.MIMEText},
			Parameters: []definition.Parameter{
				{
					Source:      definition.Query,
					Name:        "msg",
					Description: "Corresponding to the second parameter",
				},
			},
			Results: []definition.Result{
				{
					Destination: definition.Data,
					Description: "Corresponding to the first result",
				},
				{
					Destination: definition.Error,
					Description: "Corresponding to the second result",
				},
			},
		},
	},
}

//EchoMsg for test
func EchoMsg(ctx context.Context, msg string) (string, error) {
	return msg, nil
}
