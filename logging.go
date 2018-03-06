package main

import (
	lorg "github.com/kovetskiy/lorg"
)

func mustSetupLogger(debugMode bool) {
	logger = lorg.NewLog()
	defaultFormatting := lorg.NewFormat(
		`${time:15:04:05} ${level:[%s]:right:short} %s `,
	)

	logger.SetFormat(defaultFormatting)
	logger.SetLevel(lorg.LevelInfo)

	if debugMode {
		logger.SetLevel(lorg.LevelDebug)
	}
}
