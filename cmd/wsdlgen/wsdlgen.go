package main // import "github.com/weasel-software/go-xml/cmd/wsdlgen"

import (
	"log"
	"os"

	"github.com/weasel-software/go-xml/wsdlgen"
	"github.com/weasel-software/go-xml/xsdgen"
)

func main() {
	log.SetFlags(0)
	var cfg wsdlgen.Config
	cfg.Option(wsdlgen.DefaultOptions...)
	cfg.XSDOption(xsdgen.DefaultOptions...)
	cfg.Option(wsdlgen.LogOutput(log.New(os.Stderr, "", 0)))

	if err := cfg.GenCLI(os.Args[1:]...); err != nil {
		log.Fatal(err)
	}
}
