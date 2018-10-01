// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/client-go/servicecertsigner/clientset/versioned/typed/servicecertsigner/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServicecertsignerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeServicecertsignerV1alpha1) ServiceCertSignerOperatorConfigs() v1alpha1.ServiceCertSignerOperatorConfigInterface {
	return &FakeServiceCertSignerOperatorConfigs{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeServicecertsignerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
