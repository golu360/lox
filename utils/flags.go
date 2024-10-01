package utils

import "flag"

type CLIFlags struct {
	Output string
}

func (c *CLIFlags) Parse() {
	flag.StringVar(&c.Output, "out", "", "outputFilePath")

	flag.Parse()
}

func GetCLIFlags() CLIFlags {
	return CLIFlags{}
}
