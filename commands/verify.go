package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	tab "github.com/xlab/tablewriter"
	"strings"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "sub-command verify, verify resource using terraform init, plan",
	Long:  `sub-command verify, verify resource using terraform init, plan`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("subcommand -> verify, takes one argument as tfvars filename")
		}
		err := runDry(args)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func runDry(args []string) error {
	var splitPath []string
	var backendConfig BackendConfig
	var err error
	if err = FetchEnvVariables(); err != nil {
		return err
	}
	splitPath = strings.Split(args[0], "/")
	if strings.Split(splitPath[len(splitPath) - 1], ".")[1] == "tfvars" {
		log.Infof("will run terraform init, plan on [%s]", splitPath[len(splitPath) - 1])
		fmt.Println()
		Key = strings.Split(args[0], ".")[0] + ".tfstate"
		if backendConfig, err = GenerateBackendConfig(); err != nil {
			return err
		}
		table := tab.CreateTable()
		table.AddHeaders("Resource Name", "Backend Bucket", "TFVars Name", "Backend Key")
		table.AddRow(strings.Split(splitPath[len(splitPath) - 1], ".")[0], backendConfig.S3Bucket, splitPath[len(splitPath) - 1], backendConfig.Key)
		fmt.Println(table.Render())
		fmt.Println()
		if err = TerraformInit(backendConfig); err != nil {
			return err
		}
		if err = TerraformPlan(backendConfig); err != nil {
			return err
		}
	} else if strings.Split(splitPath[len(splitPath) - 1], ".")[1] == "destroy" {
		log.Infof("will run terraform init, plan -destory on [%s]", splitPath[len(splitPath) - 1])
		Key = strings.Split(args[0], ".")[0] + ".tfstate"
		if backendConfig, err = GenerateBackendConfig(); err != nil {
			return err
		}
		if err = TerraformInit(backendConfig); err != nil {
			return err
		}
		if err = TerraformPlanDestroy(backendConfig); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("wrong file type [%s]", splitPath[len(splitPath) - 1])
	}
	return nil
}
