package installer_test

import (
	"github.com/octoblu/go-meshblu-connector-dependency-manager/installer"
	"github.com/octoblu/go-meshblu-connector-installer/osruntime"
	"github.com/octoblu/go-test-server/testserver"
	"github.com/spf13/afero"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InstallNode", func() {
	var (
		server testserver.Server
	)

	BeforeEach(func() {
		server = testserver.New(Fail)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("InstallNode", func() {
		It("should exist", func() {
			Expect(installer.InstallNode).NotTo(BeNil())
		})
	})

	Describe("InstallNodeWithoutDefaults", func() {
		var (
			err error
			Fs  afero.Fs
		)

		BeforeEach(func() {
			Fs = afero.NewMemMapFs()
		})

		Describe("Install", func() {
			Describe("In darwin amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz", &testserver.Transaction{ResponseStatus: 204})
						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL(), Fs, darwinX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz", func() {
						transaction := server.Get("GET", "/dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz")
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})

			Describe("In linux 386", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz", &testserver.Transaction{ResponseStatus: 204})

						linux386 := osruntime.OSRuntime{GOOS: "linux", GOARCH: "386"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL(), Fs, linux386)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz", func() {
						transaction := server.Get("GET", "/dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz")
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})

			Describe("In linux arm", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz", &testserver.Transaction{ResponseStatus: 204})

						linuxArm := osruntime.OSRuntime{GOOS: "linux", GOARCH: "arm"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL(), Fs, linuxArm)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz", func() {
						transaction := server.Get("GET", "/dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz")
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})

			Describe("In linux sparc", func() {
				Describe("When called", func() {
					BeforeEach(func() {
						linuxSparc := osruntime.OSRuntime{GOOS: "linux", GOARCH: "sparc"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL(), Fs, linuxSparc)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
					})
				})
			})

			Describe("In windows amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						server.Set("GET", "/dist/v5.0.0/win-x64/node.exe", &testserver.Transaction{ResponseStatus: 204})

						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL(), Fs, windowsX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/win-x64/node.exe", func() {
						transaction := server.Get("GET", "/dist/v5.0.0/win-x64/node.exe")
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})
		})
	})
})
