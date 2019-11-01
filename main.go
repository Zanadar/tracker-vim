package main

import (
	"errors"
	"fmt"
	"github.com/neovim/go-client/nvim"
	"github.com/neovim/go-client/nvim/plugin"
	"os/exec"
	"runtime"
)

var storyUrl = "somthine"

func main() {
	plugin.Main(handler)
}

func handler(p *plugin.Plugin) error {
	p.HandleFunction(&plugin.FunctionOptions{Name: "GotoStory"}, goToStory)

	return nil
}

func goToStory(vim *nvim.Nvim, args []interface{}) (string, error) {
	id, ok := args[0].(int)
	if !ok {
		return "test", errors.New("id should be a string")
	}

	url := fmt.Sprint("%s/%d", storyUrl, id)
	return "test", openbrowser(url)
}

func openbrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return err
	}
	return nil
}
