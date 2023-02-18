package admitter

import (
	"github.com/sirupsen/logrus"
	admissionregv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/harvester/webhook/pkg/server/admission"

	samplev1 "github.com/harvester/webhook-sample/pkg/apis/sample.webhook.io/v1"
)

type Mutator struct {
	admission.DefaultMutator
}

var _ admission.Mutator = &Mutator{}

func NewMutator() *Mutator {
	return &Mutator{}
}

func (m *Mutator) Create(request *admission.Request, newObj runtime.Object) (admission.Patch, error) {
	foo := newObj.(*samplev1.Foo)

	labels := foo.GetLabels()
	if labels == nil {
		labels = make(map[string]string)
	}

	labels["sample"] = "sample"

	logrus.Info("add sample label")

	return admission.Patch{
		admission.PatchOp{
			Op:    admission.PatchOpReplace,
			Path:  "/metadata/labels",
			Value: labels,
		},
	}, nil
}

func (m *Mutator) Resource() admission.Resource {
	return admission.Resource{
		Names:      []string{"foos"},
		Scope:      admissionregv1.NamespacedScope,
		APIGroup:   samplev1.SchemeGroupVersion.Group,
		APIVersion: samplev1.SchemeGroupVersion.Version,
		ObjectType: &samplev1.Foo{},
		OperationTypes: []admissionregv1.OperationType{
			admissionregv1.Create,
		},
	}
}
