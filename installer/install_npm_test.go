package installer_test

import (
	"github.com/octoblu/go-meshblu-connector-dependency-manager/installer"
	"github.com/octoblu/go-meshblu-connector-installer/osruntime"
	"github.com/octoblu/go-test-server/testserver"
	"github.com/spf13/afero"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InstallNPM", func() {
	var (
		server testserver.Server
	)

	BeforeEach(func() {
		server = testserver.New(Fail)
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("InstallNPM", func() {
		It("should exist", func() {
			Expect(installer.InstallNPM).NotTo(BeNil())
		})
	})

	Describe("InstallNPMWithoutDefaults", func() {
		var (
			err error
			Fs  afero.Fs
		)

		BeforeEach(func() {
			Fs = afero.NewMemMapFs()
		})

		Describe("InstallNPMWithoutDefaults", func() {
			Describe("In darwin amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNPMWithoutDefaults("v5.0.0", server.URL(), Fs, darwinX64)
					})

					It("should return no error (and make no http requests)", func() {
						Expect(err).To(BeNil())
					})
				})
			})

			Describe("In linux 386", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						linux386 := osruntime.OSRuntime{GOOS: "linux", GOARCH: "386"}
						err = installer.InstallNPMWithoutDefaults("v5.0.0", server.URL(), Fs, linux386)
					})

					It("should return no error (and make no http requests)", func() {
						Expect(err).To(BeNil())
					})
				})
			})

			Describe("In linux arm", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						linuxArm := osruntime.OSRuntime{GOOS: "linux", GOARCH: "arm"}
						err = installer.InstallNPMWithoutDefaults("v5.0.0", server.URL(), Fs, linuxArm)
					})

					It("should return no error (and make no http requests)", func() {
						Expect(err).To(BeNil())
					})
				})
			})

			Describe("In windows amd64", func() {
				Describe("When installing NPM", func() {
					BeforeEach(func() {
						server.Set("GET", "/npm/npm/archive/v3.10.0.zip", &testserver.Transaction{ResponseStatus: 204})

						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNPMWithoutDefaults("v3.10.0", server.URL(), Fs, windowsX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /npm/npm/archive/v3.10.0.zip", func() {
						transaction := server.Get("GET", "/npm/npm/archive/v3.10.0.zip")
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})
		})
	})
})
