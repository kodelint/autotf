package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	EnvS3Bucket = "S3_BUCKET"
	EnvDynamoDB = "DYNAMODB"
	EnvRegion = "REGION"
)

var (
	S3Bucket string
	DynamoDB string
	Region string
	Key string
)

var rootCmd = &cobra.Command{
	Use:   "autotf",
	Short: "autotf is wrapper tool which runs terraform commands based on the provided tfvars file as argument",
	Long: `autotf is wrapper tool which runs terraform commands based on the provided tfvars file as an argument. Based 
on the tfvars filepath, environment variables, it automatically generates terraform backend configuration. This way autotf 
maintains a 1-to-1 relationship between tfvars and tfstate file`,
}

type BackendConfig struct {
	S3Bucket string
	DynamoDB string
	Key      string
	Region   string
}

type LoggerLogFormatter struct {
	TimestampFormat string
	LevelDesc []string
}

func (f *LoggerLogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(f.TimestampFormat))
	return []byte(fmt.Sprintf("%s %s %s\n", f.LevelDesc[entry.Level], timestamp, entry.Message)), nil
}

func init() {
	rootCmd.AddCommand(verifyCmd, deployCmd)
}

func Execute() {
	logFormat := new(LoggerLogFormatter)
	logFormat.TimestampFormat = "2006-01-02 15:04:05"
	logFormat.LevelDesc = []string{"PANIC", "FAIL", "ERROR", "WARN", "INFO", "DEBUG"}
	log.SetFormatter(logFormat)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}