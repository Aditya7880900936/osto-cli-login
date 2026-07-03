package cli

import "github.com/chzyer/readline"

var Completer = readline.NewPrefixCompleter(
	readline.PcItem("register"),
	readline.PcItem("login"),
	readline.PcItem("whoami"),
	readline.PcItem("logout"),
	readline.PcItem("enable-2fa"),
	readline.PcItem("disable-2fa"),
	readline.PcItem("help"),
	readline.PcItem("exit"),
)