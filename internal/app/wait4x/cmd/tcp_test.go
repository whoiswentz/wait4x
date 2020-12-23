package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/atkrad/wait4x/internal/pkg/errors"
	"github.com/atkrad/wait4x/internal/pkg/test"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)

	os.Exit(m.Run())
}

func TestTcpCommandInvalidArgument(t *testing.T) {
	rootCmd := NewRootCommand()
	rootCmd.AddCommand(NewTCPCommand())

	_, err := test.ExecuteCommand(rootCmd, "tcp")

	assert.Equal(t, "ADDRESS is required argument for the tcp command", err.Error())
}

func TestTcpConnectionSuccess(t *testing.T) {
	rootCmd := NewRootCommand()
	rootCmd.AddCommand(NewTCPCommand())

	_, err := test.ExecuteCommand(rootCmd, "tcp", "1.1.1.1:53")

	assert.Nil(t, err)
}

func TestTcpConnectionFail(t *testing.T) {
	rootCmd := NewRootCommand()
	rootCmd.AddCommand(NewTCPCommand())

	_, err := test.ExecuteCommand(rootCmd, "tcp", "127.0.0.1:8080", "-t", "2s")

	assert.Equal(t, errors.TimedOutErrorMessage, err.Error())
}