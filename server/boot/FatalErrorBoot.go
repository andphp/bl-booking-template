package boot

import "os"

var FatalErrorChan chan error
var Quit chan os.Signal

func init() {
	FatalErrorChan = make(chan error)
	Quit = make(chan os.Signal)
}
