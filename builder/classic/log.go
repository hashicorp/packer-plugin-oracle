// Copyright IBM Corp. 2013, 2025
// SPDX-License-Identifier: MPL-2.0

package classic

import "log"

type Logger struct {
	Enabled bool
}

func (l *Logger) Log(input ...interface{}) {
	if !l.Enabled {
		return
	}
	log.Println(input...)
}
