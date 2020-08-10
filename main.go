package main

import (
    "github.com/c-bata/go-prompt"
    "os"
    "strings"
)

func completer(d prompt.Document) []prompt.Suggest {
    s := []prompt.Suggest{
        {Text: "users", Description: "Store the username and age"},
        {Text: "articles", Description: "Store the article text posted by username"},
        {Text: "comments", Description: "Store the text commented to articles"},
    }
    return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func executor(t string) {
    cmd := strings.Split(t, " ")[0]

    switch cmd {
        case "exit": {
            os.Exit(0)
        }
        case "import": {

        }
        default: {
            LivePrefixState.LivePrefix = "... "
            LivePrefixState.IsEnable = true
        }
    }
    return
}

var LivePrefixState struct {
    LivePrefix string
    IsEnable   bool
}

func changeLivePrefix() (string, bool) {
    return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}
func main() {

    p := prompt.New(
        executor,
        completer,
        prompt.OptionPrefix(">>> "),
        prompt.OptionLivePrefix(changeLivePrefix),
    )
    p.Run()

}
