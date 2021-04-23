// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package pfcp

import (
	"log"

	"github.com/wmnsk/go-pfcp/internal/logger"
)

// SetLogger replaces the standard logger with arbitrary *log.Logger.
//
// This package prints just informational logs from goroutines working background
// that might help developers test the program but can be ignored safely. More
// important ones that need any action by the caller would be returned as errors.
func SetLogger(l *log.Logger) {
	logger.SetLogger(l)
}

// EnableLogging enables the logging from the package.
//
// If l is nil, it uses default logger provided by the package.
// Logging is enabled by default.
//
// See also: SetLogger.
func EnableLogging(l *log.Logger) {
	logger.EnableLogging(l)
}

// DisableLogging disables the logging from the package.
//
// Logging is enabled by default.
func DisableLogging() {
	logger.DisableLogging()
}
