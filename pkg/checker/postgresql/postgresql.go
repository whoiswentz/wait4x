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

package postgresql

import (
	"context"
	"database/sql"
	"github.com/atkrad/wait4x/pkg/checker"

	// Needed for the PostgreSQL driver
	_ "github.com/lib/pq"
)

var OpenDatabaseErr = checker.NewError("open database error", "debug")
var PingErr = checker.NewError("ping error", "debug")

// PostgreSQL represents PostgreSQL checker
type PostgreSQL struct {
	dsn string
	*checker.LogAware
}

// New creates the PostgreSQL checker
func New(dsn string) checker.Checker {
	p := &PostgreSQL{
		dsn:      dsn,
		LogAware: &checker.LogAware{},
	}

	return p
}

// Check checks PostgreSQL connection
func (p *PostgreSQL) Check(ctx context.Context) error {
	p.Logger().Info("Checking PostgreSQL connection ...")
	db, err := sql.Open("postgres", p.dsn)
	if err != nil {
		return OpenDatabaseErr.WithWrap(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			p.Logger().Debug(err)
		}
	}()

	err = db.PingContext(ctx)
	if err != nil {
		return PingErr.WithWrap(err)
	}

	return nil
}
