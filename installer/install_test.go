package installer_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/octoblu/go-meshblu-connector-dependency-manager/installer"
	"github.com/octoblu/go-meshblu-connector-installer/osruntime"
	"github.com/spf13/afero"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Transaction struct {
	Request        *http.Request
	ResponseBody   string
	ResponseStatus int
}

var _ = Describe("Install", func() {
	var (
		server       *httptest.Server
		transactions map[string]*Transaction
	)

	BeforeEach(func() {
		transactions = make(map[string]*Transaction)

		server = httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			key := fmt.Sprintf("%v %v", request.Method, request.URL.Path)

			transaction, ok := transactions[key]
			if !ok {
				Fail(fmt.Sprintf("Received an unexpected message: %v", key))
			}

			transaction.Request = request
			response.Header().Add("Content-Type", "application/octet-stream")
			response.WriteHeader(transaction.ResponseStatus)
			response.Write([]byte(transaction.ResponseBody))
		}))
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
						transactions["GET /dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz"] = &Transaction{ResponseStatus: 204}
						darwinX64 := osruntime.OSRuntime{GOOS: "darwin", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL, Fs, darwinX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz", func() {
						transaction := transactions["GET /dist/v5.0.0/node-v5.0.0-darwin-x64.tar.gz"]
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})

			Describe("In linux 386", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						transactions["GET /dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz"] = &Transaction{ResponseStatus: 204}

						linux386 := osruntime.OSRuntime{GOOS: "linux", GOARCH: "386"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL, Fs, linux386)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz", func() {
						transaction := transactions["GET /dist/v5.0.0/node-v5.0.0-linux-x86.tar.gz"]
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})

			Describe("In linux arm", func() {
				Describe("When everything goes well", func() {
					BeforeEach(func() {
						transactions["GET /dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz"] = &Transaction{ResponseStatus: 204}

						linuxArm := osruntime.OSRuntime{GOOS: "linux", GOARCH: "arm"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL, Fs, linuxArm)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz", func() {
						transaction := transactions["GET /dist/v5.0.0/node-v5.0.0-linux-armv71.tar.gz"]
						Expect(transaction.Request).NotTo(BeNil())
					})
				})
			})

			Describe("In linux sparc", func() {
				Describe("When called", func() {
					BeforeEach(func() {
						linuxSparc := osruntime.OSRuntime{GOOS: "linux", GOARCH: "sparc"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL, Fs, linuxSparc)
					})

					It("should return an error", func() {
						Expect(err).NotTo(BeNil())
					})
				})
			})

			Describe("In windows amd64", func() {
				Describe("When installing node", func() {
					BeforeEach(func() {
						transactions["GET /dist/v5.0.0/win-x64/node.exe"] = &Transaction{ResponseStatus: 204}

						windowsX64 := osruntime.OSRuntime{GOOS: "windows", GOARCH: "amd64"}
						err = installer.InstallNodeWithoutDefaults("v5.0.0", server.URL, Fs, windowsX64)
					})

					It("should return no error", func() {
						Expect(err).To(BeNil())
					})

					It("should GET /dist/v5.0.0/win-x64/node.exe", func() {
						transaction := transactions["GET /dist/v5.0.0/win-x64/node.exe"]
						Expect(transaction.Request).NotTo(BeNil())
					})
				})

				// Describe("When installing npm", func() {
				// 	BeforeEach(func() {
				// 		transactions["GET /dist/v3.0.0/node-v5.0.0-x64.zip"] = &Transaction{ResponseStatus: 204}
				// 		err = sut.Install("npm", "v3.10.0")
				// 	})
				//
				// 	It("should return no error", func() {
				// 		Expect(err).To(BeNil())
				// 	})
				//
				// 	It("should GET /dist/v5.0.0/node-v5.0.0-x64.zip", func() {
				// 		transaction := transactions["GET /dist/v5.0.0/node-v5.0.0-x64.zip"]
				// 		Expect(transaction.Request).NotTo(BeNil())
				// 	})
				// })
			})
		})
	})
})
