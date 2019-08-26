package controller

import (
	"github.com/openshift/rbac-permissions-operator/pkg/controller/subjectpermission"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, subjectpermission.Add)
}
