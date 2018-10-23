package env

import (
	"fmt"

	"github.com/pkg/errors"
	sls "github.com/serverless-operator/serverless-operator/pkg/apis/serverlessrelease/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetEnvVarValue returns the value referenced by the supplied EnvVarSource given the other supplied information.
func GetEnvVarValue(client kubernetes.Interface, namespace string, from *sls.EnvVarSource) (string, error) {
	if from.SecretKeyRef != nil {
		return getSecretValue(client, namespace, from.SecretKeyRef)
	}
	if from.ConfigMapKeyRef != nil {
		return getConfigValue(client, namespace, from.ConfigMapKeyRef)
	}

	return "", fmt.Errorf("invalid valueFrom")
}

// See https://github.com/kubernetes/kubernetes/blob/137193faafbc3eadf4ef5ce62bbfcd0f2e7ceca9/pkg/kubectl/cmd/set/env/env_resolve.go#L50
// not implementing the store so that we query every time
func getSecretValue(client kubernetes.Interface, namespace string, secretSelector *corev1.SecretKeySelector) (string, error) {
	secret, err := client.CoreV1().Secrets(namespace).Get(secretSelector.Name, metav1.GetOptions{})
	if err != nil {
		return "", errors.Wrap(err, "getting secret")
	}
	if data, ok := secret.Data[secretSelector.Key]; ok {
		return string(data), nil
	}
	return "", fmt.Errorf("key %s not found in secret %s", secretSelector.Key, secretSelector.Name)
}

// See https://github.com/kubernetes/kubernetes/blob/137193faafbc3eadf4ef5ce62bbfcd0f2e7ceca9/pkg/kubectl/cmd/set/env/env_resolve.go#L68
// not implementing the store so that we query time
func getConfigValue(client kubernetes.Interface, namespace string, configSelector *corev1.ConfigMapKeySelector) (string, error) {
	configMap, err := client.CoreV1().ConfigMaps(namespace).Get(configSelector.Name, metav1.GetOptions{})
	if err != nil {
		return "", errors.Wrap(err, "getting config map")
	}
	if data, ok := configMap.Data[configSelector.Key]; ok {
		return string(data), nil
	}
	return "", fmt.Errorf("key %s not found in configmap %s", configSelector.Key, configSelector.Name)
}
