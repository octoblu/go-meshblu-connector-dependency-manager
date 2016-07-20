package installer_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

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
			binPath string
			err     error
		)

		BeforeEach(func() {
			binPath, err = ioutil.TempDir(os.TempDir(), "installer-tests")
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			// os.RemoveAll(binPath)
		})

		Describe("InstallNPMWithoutDefaults", func() {
			Describe("In darwin amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNPMWithoutDefaults("v5.0.0", binPath, server.URL(), darwinX64)
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
						err = installer.InstallNPMWithoutDefaults("v5.0.0", binPath, server.URL(), linux386)
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
						err = installer.InstallNPMWithoutDefaults("v5.0.0", binPath, server.URL(), linuxArm)
					})

					It("should return no error (and make no http requests)", func() {
						Expect(err).To(BeNil())
					})
				})
			})

			Describe("In windows amd64", func() {
				Describe("When installing NPM", func() {
					BeforeEach(func() {
						var data []byte
						data, err = afero.ReadFile(afero.NewOsFs(), "fixtures/npm-3.10.0.zip")
						Expect(err).To(BeNil())

						server.Set("GET", "/npm/npm/archive/v3.10.0.zip", &testserver.Transaction{ResponseStatus: 200, ResponseBody: data})

						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNPMWithoutDefaults("v3.10.0", binPath, server.URL(), windowsX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /npm/npm/archive/v3.10.0.zip", func() {
						transaction := server.Get("GET", "/npm/npm/archive/v3.10.0.zip")
						Expect(transaction.Request).NotTo(BeNil())
					})

					It("should save the NPM response body on the filesystem", func() {
						expectedData, err := afero.ReadFile(afero.NewOsFs(), "fixtures/npm-3.10.0/bin/npm")
						Expect(err).To(BeNil())

						actualData, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(binPath, "npm"))
						Expect(err).To(BeNil())
						Expect(actualData).To(Equal(expectedData))
					})

					It("should save the NPM.cmd response body on the filesystem", func() {
						expectedData, err := afero.ReadFile(afero.NewOsFs(), "fixtures/npm-3.10.0/bin/npm.cmd")
						Expect(err).To(BeNil())

						actualData, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(binPath, "npm.cmd"))
						Expect(err).To(BeNil())
						Expect(actualData).To(Equal(expectedData))
					})
				})

				Describe("When the server responds with a non 200", func() {
					BeforeEach(func() {
						server.Set("GET", "/npm/npm/archive/v3.10.0.zip", &testserver.Transaction{ResponseStatus: 404})

						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNPMWithoutDefaults("v3.10.0", binPath, server.URL(), windowsX64)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
						Expect(err.Error()).To(ContainSubstring("Expected HTTP status code 200, received: 404"))
					})
				})

				Describe("When the server refuses the connection", func() {
					BeforeEach(func() {
						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNPMWithoutDefaults("v3.10.0", binPath, "http://0.0.0.0:0", windowsX64)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
					})
				})
			})
		})
	})
})
