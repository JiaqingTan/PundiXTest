package util

import "pundixtest/constant"

func GetFormattedFxcoredCommand(arguments []string) string {
	cmd := constant.FXCoredCommand + " "

	for _, argument := range arguments {
		cmd += argument + " "
	}

	cmd += constant.NodeFlag

	return cmd
}