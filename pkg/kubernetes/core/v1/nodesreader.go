// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kubernetes "k8s.io/client-go/kubernetes"
	"time"
)

type NodesReader interface {
	Get(name string) (*Node, error)
	List() ([]*Node, error)
}

func NewNodesReader(client resource.Client, filter resource.Filter) NodesReader {
	return &nodesReader{
		Client: client,
		filter: filter,
	}
}

type nodesReader struct {
	resource.Client
	filter resource.Filter
}

func (c *nodesReader) Get(name string) (*Node, error) {
	node := &corev1.Node{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), NodeKind.Scoped).
		Resource(NodeResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(node)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   NodeKind.Group,
			Version: NodeKind.Version,
			Kind:    NodeKind.Kind,
		}, node.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    NodeKind.Group,
				Resource: NodeResource.Name,
			}, name)
		}
	}
	return NewNode(node, c.Client), nil
}

func (c *nodesReader) List() ([]*Node, error) {
	list := &corev1.NodeList{}
	client, err := kubernetes.NewForConfig(c.Config())
	if err != nil {
		return nil, err
	}
	err = client.CoreV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), NodeKind.Scoped).
		Resource(NodeResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*Node, 0, len(list.Items))
	for _, node := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   NodeKind.Group,
			Version: NodeKind.Version,
			Kind:    NodeKind.Kind,
		}, node.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := node
			results = append(results, NewNode(&copy, c.Client))
		}
	}
	return results, nil
}
