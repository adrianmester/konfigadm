package phases

import (
	"strings"
	"testing"

	"github.com/onsi/gomega"
)

func TestCommandRuntimeFlag(t *testing.T) {
}

func setupCommandFixture(t *testing.T, flag Flag) (string, *gomega.WithT) {
	cfg, g := NewFixture("commands.yml", t).WithFlags(flag).Build()
	_, commands, _ := cfg.ApplyPhases()
	return commands, g
}

func TestCommand(t *testing.T) {
	commands, g := setupCommandFixture(t, DEBIAN)
	g.Expect(commands).To(gomega.ContainSubstring("echo command"))
}

func TestPreCommand(t *testing.T) {
	commands, g := setupCommandFixture(t, DEBIAN)
	g.Expect(commands).To(gomega.ContainSubstring("echo pre"))
}
func TestPostCommand(t *testing.T) {
	commands, g := setupCommandFixture(t, DEBIAN)
	g.Expect(commands).To(gomega.ContainSubstring("echo post"))
}
func TestCommandInterpolation(t *testing.T) {
}

func (cfg *Config) FindCmd(prefix string) []*Command {
	cmds := []*Command{}

	for _, cmd := range cfg.PreCommands {
		if strings.HasPrefix(cmd.Cmd, prefix) {
			cmds = append(cmds, &cmd)
		}
	}
	return cmds
}
