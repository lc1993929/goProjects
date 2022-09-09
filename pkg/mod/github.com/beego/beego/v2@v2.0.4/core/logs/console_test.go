// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Try each log level in decreasing order of priority.
func testConsoleCalls(bl *BeeLogger) {
	bl.Emergency("emergency")
	bl.Alert("alert")
	bl.Critical("critical")
	bl.Error("error")
	bl.Warning("warning")
	bl.Notice("notice")
	bl.Informational("informational")
	bl.Debug("debug")
}

// Test console logging by visually comparing the lines being output with and
// without a log level specification.
func TestConsole(t *testing.T) {
	log1 := NewLogger(10000)
	log1.EnableFuncCallDepth(true)
	log1.SetLogger("console", "")
	testConsoleCalls(log1)

	log2 := NewLogger(100)
	log2.SetLogger("console", `{"level":3}`)
	testConsoleCalls(log2)
}

// Test console without color
func TestConsoleNoColor(t *testing.T) {
	log := NewLogger(100)
	log.SetLogger("console", `{"color":false}`)
	testConsoleCalls(log)
}

// Test console async
func TestConsoleAsync(t *testing.T) {
	log := NewLogger(100)
	log.SetLogger("console")
	log.Async()
	// log.Close()
	testConsoleCalls(log)
	for len(log.msgChan) != 0 {
		time.Sleep(1 * time.Millisecond)
	}
}

func TestFormat(t *testing.T) {
	log := newConsole()
	lm := &LogMsg{
		Level:      LevelDebug,
		Msg:        "Hello, world",
		When:       time.Date(2020, 9, 19, 20, 12, 37, 9, time.UTC),
		FilePath:   "/user/home/main.go",
		LineNumber: 13,
		Prefix:     "Cus",
	}
	res := log.Format(lm)
	assert.Equal(t, "2020/09/19 20:12:37.000 \x1b[1;44m[D]\x1b[0m Cus Hello, world", res)
	err := log.WriteMsg(lm)
	assert.Nil(t, err)
}
