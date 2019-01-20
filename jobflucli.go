// jobflucli is an application I made to test the capabilities of Go and TUI.
// It is also a useful terminal frontend to read the JobFluent public feed
// so that you can view interesting computer engineering job offers without
// leaving your warm terminal.
package main

func main() {
	context := new(Context)
	ui := NewUserInterface(context)
	ui.SwitchToLocations()
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
