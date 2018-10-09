package operator

import (
	"reflect"

	opkit "github.com/rook/operator-kit"
	slsapi "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease"
	slsapiv1 "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

var Resource = opkit.CustomResource{
	Name:    "serverlessrelease",
	Plural:  "serverlessreleases",
	Group:   slsapi.GroupName,
	Version: "v1alpha1",
	Scope:   apiextensionsv1beta1.NamespaceScoped,
	Kind:    reflect.TypeOf(slsapiv1.ServerlessRelease{}).Name(),
	ShortNames: []string{
		"sls",
		"slsrelease",
	},
}