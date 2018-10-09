package operator

import (
	opkit "github.com/christopherhein/operator-kit"
	sls "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease/v1alpha1"
	slsclient "github.com/serverless-operator/serverless-operator/pkg/client/clientset/versioned/typed/serverlessrelease/v1alpha1"
	"github.com/serverless-operator/serverless-operator/pkg/config"
	"k8s.io/client-go/tools/cache"
)

type Controller struct {
	config       *config.Config
	context      *opkit.Context
	slsClientSet slsclient.ReleaseV1alpha1Interface
}

func NewController(config *config.Config, context *opkit.Context, slsClientSet slsclient.ReleaseV1alpha1Interface) *Controller {
	return &Controller{
		config:       config,
		context:      context,
		slsClientSet: slsClientSet,
	}
}

// StartWatch watches for instances of Object Store custom resources and acts on them
func (c *Controller) StartWatch(namespace string, stopCh <-chan struct{}) error {

	resourceHandlers := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	}
	restClient := c.slsClientSet.RESTClient()
	watcher := opkit.NewWatcher(sls.ServerlessRelease, namespace, resourceHandlers, restClient)
	go watcher.Watch(&sls.ServerlessRelease{}, stopCh)

	return nil
}

func (c *Controller) onAdd(obj interface{}) {
	slsRelease := obj.(*sls.ServerlessRelease).DeepCopy()

	c.config.Logger.Infof("Serverless release added %s and package %s\n", slsRelease.Spec.ReleaseName, slsRelease.Spec.PackagePath)
}

func (c *Controller) onUpdate(oldObj, newObj interface{}) {
	slsReleaseOld := obj.(*oldObj.ServerlessRelease).DeepCopy()
	slsReleaseNew := obj.(*newObj.ServerlessRelease).DeepCopy()

	c.config.Logger.Infof("Serverless release updated from %s to package %s\n", slsReleaseOld.Spec.PackagePath, slsReleaseNew.Spec.PackagePath)
}

func (c *Controller) onDelete(obj interface{}) {
	slsRelease := obj.(*sls.ServerlessRelease).DeepCopy()

	c.config.Logger.Infof("Serverless release deleted %s and package %s\n", slsRelease.Spec.ReleaseName, slsRelease.Spec.PackagePath)
}
