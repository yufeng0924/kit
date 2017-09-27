package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/kujtimiihoxha/kit/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var methods []string
var initserviceCmd = &cobra.Command{
	Use:     "service",
	Short:   "Initiate a service",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logrus.Error("You must provide a name for the service")
			return
		}
		if viper.GetString("g_s_transport") == "grpc" {
			if !checkProtoc() {
				return
			}
		}
		emw := false
		smw := false
		if viper.GetBool("g_s_dmw") {
			emw = true
			smw = true
		} else {
			emw = viper.GetBool("g_s_endpoint_mdw")
			smw = viper.GetBool("g_s_svc_mdw")
		}
		g := generator.NewGenerateService(
			args[0],
			viper.GetString("g_s_transport"),
			smw,
			emw,
			methods,
		)
		if err := g.Generate(); err != nil {
			logrus.Error(err)
		}
	},
}

func init() {
	generateCmd.AddCommand(initserviceCmd)
	initserviceCmd.Flags().StringP("transport", "t", "http", "The transport you want your service to be initiated with")
	initserviceCmd.Flags().Bool("dmw", false, "Generate default middleware for service and endpoint")
	initserviceCmd.Flags().StringArrayVarP(&methods, "methods", "m", []string{}, "Specify methods to be generated")
	initserviceCmd.Flags().Bool("svc-mdw", false, "If set a default Logging and Instrumental middleware will be created and attached to the service")
	initserviceCmd.Flags().Bool("endpoint-mdw", false, "If set a default Logging and Tracking middleware will be created and attached to the endpoint")
	viper.BindPFlag("g_s_transport", initserviceCmd.Flags().Lookup("transport"))
	viper.BindPFlag("g_s_dmw", initserviceCmd.Flags().Lookup("dmw"))
	viper.BindPFlag("g_s_svc_mdw", initserviceCmd.Flags().Lookup("svc-mdw"))
	viper.BindPFlag("g_s_endpoint_mdw", initserviceCmd.Flags().Lookup("endpoint-mdw"))
}
