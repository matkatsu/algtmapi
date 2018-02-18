// +build !production,!test

package env

func initialize() {
	OnDevelopment = true
	configfilename = "local"
}
