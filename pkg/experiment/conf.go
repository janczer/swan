package experiment

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/intelsdi-x/swan/pkg/conf"
	"github.com/intelsdi-x/swan/pkg/utils/errutil"
)

var (
	// Both flags are defined using directly go native "flag" package to not be registered as experiment configuration.
	dumpConfig             = flag.Bool("config-dump", false, "Dump configuration as environment script.")
	dumpConfigExperimentID = flag.String("config-dump-experiment-id", "", "Dump configuration based on experiment ID.")
)

// Configure handles configuration parsing, generation and restoration based on config-* flags.
// Note: exits if configuration generation was requested.
// This function must reside in experiment package because depends on metadata access.
// Returns information about current log level.
func Configure() bool {

	err := conf.ParseFlags()
	errutil.Check(err)

	level, err := conf.LogLevel()
	errutil.Check(err)
	logrus.SetLevel(level)

	if *dumpConfig {
		previousExperimentID := *dumpConfigExperimentID
		if previousExperimentID != "" {
			metadata := NewMetadata(previousExperimentID, MetadataConfigFromFlags())
			err := metadata.Connect()
			errutil.Check(err)
			flags, err := metadata.GetGroup("flags")
			errutil.Check(err)
			fmt.Println(conf.DumpConfigMap(flags))
		} else {
			fmt.Println(conf.DumpConfig())
		}
		os.Exit(0)
	}
	return level == logrus.ErrorLevel
}