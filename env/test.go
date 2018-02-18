// +build test

package env

func initialize() {
	OnDevelopment = true
	configfilename = "test"
}
