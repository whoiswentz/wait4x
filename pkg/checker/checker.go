// Copyright 2020 Mohammad Abdolirad
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checker

import (
	"context"
	"github.com/atkrad/wait4x/pkg/log"
)

// Checker is the interface that wraps the basic checker methods.
type Checker interface {
	SetLogger(logger log.Logger)
	Logger() log.Logger
	Check(ctx context.Context) error
}

// LogAware represents log object.
type LogAware struct {
	logger log.Logger
}

// SetLogger sets default logger
func (la *LogAware) SetLogger(logger log.Logger) {
	la.logger = logger
}

// Logger gets default logger
func (la *LogAware) Logger() log.Logger {
	return la.logger
}
