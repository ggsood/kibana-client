package kbapi

import (
	"github.com/go-resty/resty/v2"
)

// API handle the API specification
type API struct {
	KibanaSpaces *KibanaSpacesAPI
}

// KibanaSpacesAPI handle the spaces API
type KibanaSpacesAPI struct {
	Get              KibanaSpaceGet
	List             KibanaSpaceList
	Create           KibanaSpaceCreate
	Delete           KibanaSpaceDelete
	Update           KibanaSpaceUpdate
	CopySavedObjects KibanaSpaceCopySavedObjects
}

// New initialise the API implementation
func New(c *resty.Client) *API {
	return &API{
		KibanaSpaces: &KibanaSpacesAPI{
			Get:              newKibanaSpaceGetFunc(c),
			List:             newKibanaSpaceListFunc(c),
			Create:           newKibanaSpaceCreateFunc(c),
			Update:           newKibanaSpaceUpdateFunc(c),
			Delete:           newKibanaSpaceDeleteFunc(c),
			CopySavedObjects: newKibanaSpaceCopySavedObjectsFunc(c),
		},
	}
}
