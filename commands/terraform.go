package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"strings"
)

func TerraformInit(b BackendConfig) error {
	s3 := "-backend-config=bucket=" + b.S3Bucket
	dynamodb := "-backend-config=dynamodb_table=" + b.DynamoDB
	key := "-backend-config=key=" + b.Key
	region := "-backend-config=region=" + b.Region
	cmdArgs := []string{"init", s3, dynamodb, region, key}
	cmd := exec.Command("terraform", cmdArgs...)
	log.Debugf("Running initialization using [ %s ] command [from function terraformInit()]", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize terraform for [%s]", strings.Split(b.Key, ".")[0] + "tfvars")
	}
	return nil
}

func TerraformApply(b BackendConfig) error {
	cmdArgs := []string{"apply", "-auto-approve", "-var-file=" + strings.Split(b.Key, ".")[0] + ".tfvars"}
	cmd := exec.Command("terraform", cmdArgs...)
	log.Debugf("Running apply using [ %s ] command [from function TerraformApply()]", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to terraform apply for [%s]", strings.Split(b.Key, ".")[0] + ".tfvars")
	}
	return nil
}

func TerraformDestroy(b BackendConfig) error {
	cmdArgs := []string{"destroy", "-auto-approve", "-var-file=" + strings.Split(b.Key, ".")[0] + ".destroy"}
	cmd := exec.Command("terraform", cmdArgs...)
	log.Debugf("Running terraform destroy using [ %s ] command [from function TerraformDestroy()]", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to terraform destroy for [%s]", strings.Split(b.Key, ".")[0] + ".destroy")
	}
	return nil
}

func TerraformPlan(b BackendConfig) error {
	cmdArgs := []string{"plan", "-var-file=" + strings.Split(b.Key, ".")[0] + ".tfvars"}
	cmd := exec.Command("terraform", cmdArgs...)
	log.Debugf("Running terraform plan using [ %s ] command [from function TerraformPlan()]", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to terraform plan for [%s]", strings.Split(b.Key, ".")[0] + ".tfvars")
	}
	return nil
}

func TerraformPlanDestroy(b BackendConfig) error {
	cmdArgs := []string{"plan", "-destroy", "-var-file=" + strings.Split(b.Key, ".")[0] + ".destroy"}
	cmd := exec.Command("terraform", cmdArgs...)
	log.Debugf("Running terraform plan destroy using [ %s ] command [from function TerraformPlanDestroy()]", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to terraform plan destroy terraform for [%s]", strings.Split(b.Key, ".")[0] + ".destroy")
	}
	return nil
}