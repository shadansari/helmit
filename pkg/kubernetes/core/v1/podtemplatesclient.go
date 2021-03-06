// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
)

type PodTemplatesClient interface {
	PodTemplates() PodTemplatesReader
}

func NewPodTemplatesClient(resources resource.Client, filter resource.Filter) PodTemplatesClient {
	return &podTemplatesClient{
		Client: resources,
		filter: filter,
	}
}

type podTemplatesClient struct {
	resource.Client
	filter resource.Filter
}

func (c *podTemplatesClient) PodTemplates() PodTemplatesReader {
	return NewPodTemplatesReader(c.Client, c.filter)
}
