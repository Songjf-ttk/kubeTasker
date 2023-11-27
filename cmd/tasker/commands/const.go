package commands

import rbacv1 "k8s.io/api/rbac/v1"

const (
	// kubeTasker controller resource constants
	KubeTaskerControllerServiceAccount     = "tasker"
	KubeTaskerControllerClusterRole        = "tasker-cluster-role"
	KubeTaskerControllerClusterRoleBinding = "tasker-binding"

	// Argo UI resource constants
	ArgoUIServiceAccount     = "argo-ui"
	ArgoUIClusterRole        = "argo-ui-cluster-role"
	ArgoUIClusterRoleBinding = "argo-ui-binding"
	ArgoUIDeploymentName     = "argo-ui"
	ArgoUIServiceName        = "argo-ui"
)

var (
	KubeTaskerControllerPolicyRules = []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			// TODO(jesse): remove exec privileges when issue #499 is resolved
			Resources: []string{"pods", "pods/exec"},
			Verbs:     []string{"create", "get", "list", "watch", "update", "patch"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"configmaps"},
			Verbs:     []string{"get", "watch", "list"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"persistentvolumeclaims"},
			Verbs:     []string{"create", "delete"},
		},
		{
			APIGroups: []string{"kubetasker.io"},
			Resources: []string{"workflows"},
			Verbs:     []string{"get", "list", "watch", "update", "patch"},
		},
	}

	ArgoUIPolicyRules = []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"pods", "pods/exec", "pods/log"},
			Verbs:     []string{"get", "list", "watch"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"secrets"},
			Verbs:     []string{"get"},
		},
		{
			APIGroups: []string{"kubetasker.io"},
			Resources: []string{"workflows"},
			Verbs:     []string{"get", "list", "watch"},
		},
	}
)
