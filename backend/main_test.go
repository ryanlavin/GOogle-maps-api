package main

import (
	//"context"
	//"fmt"
	"fmt"
	//"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Testing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler testing")
}

var _ = Describe("Initializing server for testing", func() {
	var s server
	s.ConfigFile = "config.json"
	fmt.Println("Config file not loaded")
	s.Config = s.LoadConfig()
	fmt.Println("Config server not initialized")
	s.Router = s.InitServer()
	fmt.Println("Config server initialized")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Config server initialized")
	Context("Test server initialization", func() {
		It("Testing something", func() {
			Expect(s.Config.IndexFile).Should(Equal("../frontend/index.html"))
			Expect(s.Config.IndexURL).Should(Equal("/"))
		})

	})
})
