package common

import (
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

func TestGetEnv(t *testing.T) {
	RegisterTestingT(t)
	Expect(GetEnv("DOKKU_ROOT")).To(Equal("/home/dokku"))
}

func TestGetAppImageRepo(t *testing.T) {
	RegisterTestingT(t)
	Expect(GetAppImageRepo("testapp")).To(Equal("dokku/testapp"))
}

func TestVerifyImageInvalid(t *testing.T) {
	RegisterTestingT(t)
	Expect(VerifyImage("testapp")).To(Equal(false))
}

func TestVerifyAppName(t *testing.T) {
	RegisterTestingT(t)
	dir := "/home/dokku/testApp"
	os.MkdirAll(dir, 0644)
	Expect(VerifyAppName("testApp")).To(Equal(true))
	os.RemoveAll(dir)
}
