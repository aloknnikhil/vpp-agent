package client

import (
	"context"

	"github.com/ligato/vpp-agent/api/configurator"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func (c *Client) Configure(ctx context.Context, opts configurator.Config) (err error) {
	var conn *grpc.ClientConn
	if conn, err = c.GRPCConn(); err != nil {
		err = errors.Wrap(err, "c.GRPCConn()")
		return
	}
	grpcClient := configurator.NewConfiguratorClient(conn)
	_, err = grpcClient.Update(ctx, &configurator.UpdateRequest{Update: &opts})
	return
}
