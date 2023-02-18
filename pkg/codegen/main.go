package main

import (
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"

	"github.com/harvester/webhook-sample/pkg/apis/sample.webhook.io/v1"
	"github.com/harvester/webhook-sample/pkg/apis/sample.webhook.io/v2"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/harvester/webhook-sample/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"sample.webhook.io": {
				Types: []interface{}{
					v1.Foo{},
					v2.Foo{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
		},
	})
}
