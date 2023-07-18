package cmd

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/tensorchord/openmodelz/agent/api/types"
)

var (
	// Used for flags.
	serverListQuiet   bool
	serverListVerbose bool
)

// serverListCmd represents the server list command
var serverListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all servers in the cluster",
	Long:    `List all servers in the cluster`,
	Example: `  mdz server list`,
	PreRunE: getAgentClient,
	RunE:    commandServerList,
}

func init() {
	serverCmd.AddCommand(serverListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverListCmd.Flags().BoolVarP(&serverListQuiet, "quiet", "q", false, "Quiet mode - print out only the server names")
	serverListCmd.Flags().BoolVarP(&serverListVerbose, "verbose", "v", false, "Verbose mode - print out all server details")
}

func commandServerList(cmd *cobra.Command, args []string) error {
	servers, err := agentClient.ServerList(cmd.Context())
	if err != nil {
		return err
	}

	if serverListQuiet {
		for _, server := range servers {
			cmd.Printf("%s\n", server.Spec.Name)
		}
	} else if serverListVerbose {
		t := table.NewWriter()
		t.SetStyle(table.Style{
			Box:     table.StyleBoxDefault,
			Color:   table.ColorOptionsDefault,
			Format:  table.FormatOptionsDefault,
			HTML:    table.DefaultHTMLOptions,
			Options: table.OptionsNoBordersAndSeparators,
			Title:   table.TitleOptionsDefault,
		})
		t.AppendHeader(table.Row{"Name", "Phase", "Allocatable", "Capacity", "Distribution", "OS", "Kernel"})

		for _, server := range servers {
			t.AppendRow(table.Row{server.Spec.Name, server.Status.Phase,
				resourceListString(server.Status.Allocatable),
				resourceListString(server.Status.Capacity),
				server.Status.System.OSImage,
				server.Status.System.OperatingSystem,
				server.Status.System.KernelVersion})
		}
		cmd.Println(t.Render())
	} else {
		t := table.NewWriter()
		t.SetStyle(table.Style{
			Box:     table.StyleBoxDefault,
			Color:   table.ColorOptionsDefault,
			Format:  table.FormatOptionsDefault,
			HTML:    table.DefaultHTMLOptions,
			Options: table.OptionsNoBordersAndSeparators,
			Title:   table.TitleOptionsDefault,
		})
		t.AppendHeader(table.Row{"Name", "Phase", "Allocatable", "Capacity"})

		for _, server := range servers {
			t.AppendRow(table.Row{server.Spec.Name, server.Status.Phase,
				resourceListString(server.Status.Allocatable),
				resourceListString(server.Status.Capacity)})
		}
		cmd.Println(t.Render())
	}
	return nil
}

func resourceListString(l types.ResourceList) string {
	res := fmt.Sprintf("cpu: %s\nmem: %s", l["cpu"], l["memory"])
	if l["nvidia.com/gpu"] != "" {
		res += fmt.Sprintf("\ngpu: %s", l["nvidia.com/gpu"])
	}
	return res
}
