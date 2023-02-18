package converter

import (
	"fmt"

	"github.com/rancher/lasso/pkg/log"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	samplev1 "github.com/harvester/webhook-sample/pkg/apis/sample.webhook.io/v1"
	samplev2 "github.com/harvester/webhook-sample/pkg/apis/sample.webhook.io/v2"
	"github.com/harvester/webhook/pkg/server/conversion"
)

type fooConverter struct{}

var _ conversion.Converter = &fooConverter{}

func NewFooConverter() *fooConverter {
	return &fooConverter{}
}

func (f *fooConverter) GroupResource() schema.GroupResource {
	return samplev1.Resource(samplev1.FooResourceName)
}

func (f *fooConverter) Convert(obj *unstructured.Unstructured, toVersion string) (*unstructured.Unstructured, error) {
	fromVersion := obj.GetAPIVersion()
	log.Debugf("convert %s from %q to %q", obj.GetKind(), fromVersion, toVersion)

	if fromVersion == toVersion {
		return nil, fmt.Errorf("conversion from a version to itself should not call the webhook: %s", toVersion)
	}

	convertedObj := obj.DeepCopy()
	convertedObj.SetAPIVersion(toVersion)

	switch toVersion {
	case samplev1.SchemeGroupVersion.String():
		convertedObj.Object["alias"] = ""
	case samplev2.SchemeGroupVersion.String():
		convertedObj.Object["alias"] = obj.GetName()
	default:
		return nil, fmt.Errorf("unexpected conversion version %q", toVersion)
	}

	return convertedObj, nil
}
