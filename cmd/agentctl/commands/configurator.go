package commands

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"

	"github.com/ligato/vpp-agent/api/configurator"
	agentcli "github.com/ligato/vpp-agent/cmd/agentctl/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewConfiguratorCommand(cli agentcli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yaml",
		Short: "Import YAML",
		Example: `
To import a YAML file:
$ agentctl yaml ./network.yaml
		`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args[0]) == 0 {
				err = errors.New("No file specified")
				return
			}
			var file *os.File
			if file, err = os.Open(args[0]); err != nil {
				err = errors.Wrapf(err, "os.Open(%s)", args[0])
				return
			}

			return runConfiguratorCli(cli, file)
		},
	}
	return cmd
}

func runConfiguratorCli(cli agentcli.Cli, file *os.File) (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Read file
	var data []byte
	if data, err = ioutil.ReadAll(file); err != nil {
		err = errors.Wrap(err, "ioutil.ReadAll()")
		return
	}
	// Parse data
	config := configurator.Config{}
	if err = yaml.Unmarshal(data, &config); err != nil {
		err = errors.Wrap(err, "yaml.Unmarshal()")
		return
	}
	return cli.Client().Configure(ctx, config)
}
