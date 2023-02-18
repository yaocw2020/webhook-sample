package admitter

import (
	"github.com/sirupsen/logrus"
	admissionregv1 "k8s.io/api/admissionregistration/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/harvester/webhook/pkg/server/admission"

	samplev1 "github.com/harvester/webhook-sample/pkg/apis/sample.webhook.io/v1"
)

type Validator struct {
	admission.DefaultValidator
}

var _ admission.Validator = &Validator{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Create(request *admission.Request, newObj runtime.Object) error {
	foo := newObj.(*samplev1.Foo)
	logrus.Infof("create a new foo %s/%s", foo.Namespace, foo.Name)
	return nil
}

func (v *Validator) Resource() admission.Resource {
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
