package stash

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Rename = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Try to rename the stash.",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.
			EmptyCommit("blah").
			CreateFileAndAdd("file-1", "change to stash1").
			StashWithMessage("foo").
			CreateFileAndAdd("file-2", "change to stash2").
			StashWithMessage("bar")
	},
	Run: func(shell *Shell, input *Input, assert *Assert, keys config.KeybindingConfig) {
		input.SwitchToStashWindow()
		assert.CurrentViewName("stash")

		assert.SelectedLine(Equals("On master: bar"))
		input.NextItem()
		assert.SelectedLine(Equals("On master: foo"))
		input.Press(keys.Stash.RenameStash)

		input.Prompt(Equals("Rename stash: stash@{1}"), " baz")

		assert.SelectedLine(Equals("On master: foo baz"))
	},
})