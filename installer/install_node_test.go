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
			err     error
			binPath string
		)

		BeforeEach(func() {
			binPath, err = ioutil.TempDir(os.TempDir(), "installer-tests")
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.RemoveAll(binPath)
		})

		Describe("Install", func() {
			Describe("In darwin amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						var data []byte
						data, err = afero.ReadFile(afero.NewOsFs(), "fixtures/node-v5.0.0-darwin-x64.tar.gz")
						Expect(err).To(BeNil())

						transaction := &testserver.Transaction{ResponseStatus: 200, ResponseBody: data}
						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz", transaction)

						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, server.URL(), darwinX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz", func() {
						transaction := server.Get("GET", "/dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz")
						Expect(transaction.Request).NotTo(BeNil())
					})

					It("should save node on the filesystem", func() {
						exists, err := afero.Exists(afero.NewOsFs(), filepath.Join(binPath, "node"))
						Expect(err).To(BeNil())
						Expect(exists).To(BeTrue())
					})

					It("should save the node response body on the filesystem", func() {
						expectedData, err := afero.ReadFile(afero.NewOsFs(), "fixtures/node-v5.0.0-darwin-x64/bin/node")
						Expect(err).To(BeNil())

						actualData, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(binPath, "node"))
						Expect(err).To(BeNil())
						Expect(actualData).To(Equal(expectedData))
					})

					It("should save npm on the filesystem", func() {
						exists, err := afero.Exists(afero.NewOsFs(), filepath.Join(binPath, "npm"))
						Expect(err).To(BeNil())
						Expect(exists).To(BeTrue())
					})

					It("should save the npm response body on the filesystem", func() {
						expectedData, err := afero.ReadFile(afero.NewOsFs(), "fixtures/node-v5.0.0-darwin-x64/bin/npm")
						Expect(err).To(BeNil())

						actualData, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(binPath, "npm"))
						Expect(err).To(BeNil())
						Expect(actualData).To(Equal(expectedData))
					})
				})

				Describe("When downloading returns a 404", func() {
					BeforeEach(func() {
						transaction := &testserver.Transaction{ResponseStatus: 404}
						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz", transaction)

						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, server.URL(), darwinX64)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
						Expect(err.Error()).To(ContainSubstring("Expected HTTP status code 200, received: 404"))
					})
				})

				Describe("When the server refuses the connection", func() {
					BeforeEach(func() {
						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, "http://0.0.0.0:0", darwinX64)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
					})
				})
			})

			Describe("In linux 386", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						var data []byte
						data, err = afero.ReadFile(afero.NewOsFs(), "fixtures/node-v5.0.0-linux-x86.tar.gz")
						Expect(err).To(BeNil())

						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz", &testserver.Transaction{ResponseStatus: 200, ResponseBody: data})

						linux386 := osruntime.OSRuntime{GOOS: "linux", GOARCH: "386"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, server.URL(), linux386)
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
						var data []byte
						data, err = afero.ReadFile(afero.NewOsFs(), "fixtures/node-v5.0.0-linux-armv71.tar.gz")
						Expect(err).To(BeNil())

						server.Set("GET", "/dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz", &testserver.Transaction{ResponseStatus: 200, ResponseBody: data})

						linuxArm := osruntime.OSRuntime{GOOS: "linux", GOARCH: "arm"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, server.URL(), linuxArm)
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
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, server.URL(), linuxSparc)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
					})
				})
			})

			Describe("In windows amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						var data []byte
						data, err = afero.ReadFile(afero.NewOsFs(), "fixtures/node.exe")
						Expect(err).To(BeNil())

						server.Set("GET", "/dist/v5.0.0/win-x64/node.exe", &testserver.Transaction{ResponseStatus: 200, ResponseBody: data})

						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", binPath, server.URL(), windowsX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/win-x64/node.exe", func() {
						transaction := server.Get("GET", "/dist/v5.0.0/win-x64/node.exe")
						Expect(transaction.Request).NotTo(BeNil())
					})

					It("should save the node response body on the filesystem", func() {
						expectedData, err := afero.ReadFile(afero.NewOsFs(), "fixtures/node.exe")
						Expect(err).To(BeNil())

						actualData, err := afero.ReadFile(afero.NewOsFs(), filepath.Join(binPath, "node.exe"))
						Expect(err).To(BeNil())
						Expect(actualData).To(Equal(expectedData))
					})
				})
			})
		})
	})
})
