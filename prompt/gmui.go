package main

import (
	"fmt"
	"strings"

	. "test/gocui"

	prompt "github.com/c-bata/go-prompt"
)

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

var commands = []prompt.Suggest{
	{Text: "get", Description: "Display one or many resources"},
	{Text: "describe", Description: "Show details of a specific resource or group of resources"},
	{Text: "create", Description: "Create a resource by filename or stdin"},
	{Text: "replace", Description: "Replace a resource by filename or stdin."},
	{Text: "patch", Description: "Update field(s) of a resource using strategic merge patch."},
	{Text: "delete", Description: "Delete resources by filenames, stdin, resources and names, or by resources and label selector."},
	{Text: "edit", Description: "Edit a resource on the server"},
	{Text: "apply", Description: "Apply a configuration to a resource by filename or stdin"},
	{Text: "namespace", Description: "SUPERSEDED: Set and view the current Kubernetes namespace"},
	{Text: "logs", Description: "Print the logs for a container in a pod."},
	{Text: "rolling-update", Description: "Perform a rolling update of the given ReplicationController."},
	{Text: "scale", Description: "Set a new size for a Deployment, ReplicaSet, Replication Controller, or Job."},
	{Text: "cordon", Description: "Mark node as unschedulable"},
	{Text: "drain", Description: "Drain node in preparation for maintenance"},
	{Text: "uncordon", Description: "Mark node as schedulable"},
	{Text: "attach", Description: "Attach to a running container."},
	{Text: "exec", Description: "Execute a command in a container."},
	{Text: "port-forward", Description: "Forward one or more local ports to a pod."},
	{Text: "proxy", Description: "Run a proxy to the Kubernetes API server"},
	{Text: "run", Description: "Run a particular image on the cluster."},
	{Text: "expose", Description: "Take a replication controller, service, or pod and expose it as a new Kubernetes Service"},
	{Text: "autoscale", Description: "Auto-scale a Deployment, ReplicaSet, or ReplicationController"},
	{Text: "rollout", Description: "rollout manages a deployment"},
	{Text: "label", Description: "Update the labels on a resource"},
	{Text: "annotate", Description: "Update the annotations on a resource"},
	{Text: "config", Description: "config modifies kubeconfig files"},
	{Text: "cluster-info", Description: "Display cluster info"},
	{Text: "api-versions", Description: "Print the supported API versions on the server, in the form of 'group/version'."},
	{Text: "version", Description: "Print the client and server version information."},
	{Text: "explain", Description: "Documentation of resources."},
	{Text: "convert", Description: "Convert config files between different API versions"},
	{Text: "top", Description: "Display Resource (CPU/Memory/Storage) usage"},

	// Custom command.
	{Text: "exit", Description: "Exit this program"},
}

var resourceTypes = []prompt.Suggest{
	{Text: "clusters"}, // valid only for federation apiservers
	{Text: "componentstatuses"},
	{Text: "configmaps"},
	{Text: "daemonsets"},
	{Text: "deployments"},
	{Text: "endpoints"},
	{Text: "events"},
	{Text: "horizontalpodautoscalers"},
	{Text: "ingresses"},
	{Text: "jobs"},
	{Text: "cronjobs"},
	{Text: "limitranges"},
	{Text: "namespaces"},
	{Text: "networkpolicies"},
	{Text: "nodes"},
	{Text: "persistentvolumeclaims"},
	{Text: "persistentvolumes"},
	{Text: "pod"},
	{Text: "podsecuritypolicies"},
	{Text: "podtemplates"},
	{Text: "replicasets"},
	{Text: "replicationcontrollers"},
	{Text: "resourcequotas"},
	{Text: "secrets"},
	{Text: "serviceaccounts"},
	{Text: "services"},
	{Text: "statefulsets"},
	{Text: "storageclasses"},
	{Text: "thirdpartyresources"},

	// aliases
	{Text: "cs"},
	{Text: "cm"},
	{Text: "ds"},
	{Text: "deploy"},
	{Text: "ep"},
	{Text: "hpa"},
	{Text: "ing"},
	{Text: "limits"},
	{Text: "ns"},
	{Text: "no"},
	{Text: "pvc"},
	{Text: "pv"},
	{Text: "po"},
	{Text: "psp"},
	{Text: "rs"},
	{Text: "rc"},
	{Text: "quota"},
	{Text: "sa"},
	{Text: "svc"},
}

func argumentsCompleter(namespace string, args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "root":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(rootOptions, second, true)
		}
	case "get":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "componentstatuses"},
				{Text: "configmaps"},
				{Text: "daemonsets"},
				{Text: "deployments"},
				{Text: "endpoints"},
				{Text: "events"},
				{Text: "horizontalpodautoscalers"},
				{Text: "ingresses"},
				{Text: "jobs"},
				{Text: "cronjobs"},
				{Text: "limitranges"},
				{Text: "namespaces"},
				{Text: "networkpolicies"},
				{Text: "nodes"},
				{Text: "persistentvolumeclaims"},
				{Text: "persistentvolumes"},
				{Text: "pod"},
				{Text: "podsecuritypolicies"},
				{Text: "podtemplates"},
				{Text: "replicasets"},
				{Text: "replicationcontrollers"},
				{Text: "resourcequotas"},
				{Text: "secrets"},
				{Text: "serviceaccounts"},
				{Text: "services"},
				{Text: "statefulsets"},
				{Text: "storageclasses"},
				{Text: "thirdpartyresources"},
				// aliases
				{Text: "cs"},
				{Text: "cm"},
				{Text: "ds"},
				{Text: "deploy"},
				{Text: "ep"},
				{Text: "hpa"},
				{Text: "ing"},
				{Text: "limits"},
				{Text: "ns"},
				{Text: "no"},
				{Text: "pvc"},
				{Text: "pv"},
				{Text: "po"},
				{Text: "psp"},
				{Text: "rs"},
				{Text: "rc"},
				{Text: "quota"},
				{Text: "sa"},
				{Text: "svc"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

		// third := args[2]
		// if len(args) == 3 {
		// 	switch second {
		// 	case "componentstatuses", "cs":
		// 		return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
		// 	case "configmaps", "cm":
		// 		return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
		// 	case "daemonsets", "ds":
		// 		return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
		// 	case "deploy", "deployments":
		// 		return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
		// 	case "endpoints", "ep":
		// 		return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
		// 	case "ingresses", "ing":
		// 		return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
		// 	case "limitranges", "limits":
		// 		return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
		// 	case "namespaces", "ns":
		// 		return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
		// 	case "no", "nodes":
		// 		return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
		// 	case "po", "pod", "pods":
		// 		return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
		// 	case "persistentvolumeclaims", "pvc":
		// 		return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
		// 	case "persistentvolumes", "pv":
		// 		return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
		// 	case "podsecuritypolicies", "psp":
		// 		return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
		// 	case "podtemplates":
		// 		return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
		// 	case "replicasets", "rs":
		// 		return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
		// 	case "replicationcontrollers", "rc":
		// 		return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
		// 	case "resourcequotas", "quota":
		// 		return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
		// 	case "secrets":
		// 		return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
		// 	case "sa", "serviceaccounts":
		// 		return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
		// 	case "svc", "services":
		// 		return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
		// 	case "job", "jobs":
		// 		return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
		// 	}
		// }
	case "describe":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, second, true)
		}

	// 	third := args[2]
	// 	if len(args) == 3 {
	// 		switch second {
	// 		case "componentstatuses", "cs":
	// 			return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
	// 		case "configmaps", "cm":
	// 			return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
	// 		case "daemonsets", "ds":
	// 			return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
	// 		case "deploy", "deployments":
	// 			return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
	// 		case "endpoints", "ep":
	// 			return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
	// 		case "ingresses", "ing":
	// 			return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
	// 		case "limitranges", "limits":
	// 			return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
	// 		case "namespaces", "ns":
	// 			return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
	// 		case "no", "nodes":
	// 			return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
	// 		case "po", "pod", "pods":
	// 			return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
	// 		case "persistentvolumeclaims", "pvc":
	// 			return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
	// 		case "persistentvolumes", "pv":
	// 			return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
	// 		case "podsecuritypolicies", "psp":
	// 			return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
	// 		case "podtemplates":
	// 			return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
	// 		case "replicasets", "rs":
	// 			return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
	// 		case "replicationcontrollers", "rc":
	// 			return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
	// 		case "resourcequotas", "quota":
	// 			return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
	// 		case "secrets":
	// 			return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
	// 		case "sa", "serviceaccounts":
	// 			return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
	// 		case "svc", "services":
	// 			return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
	// 		case "job", "jobs":
	// 			return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
	// 		}
	// 	}
	case "create":
		subcommands := []prompt.Suggest{
			{Text: "configmap", Description: "Create a configmap from a local file, directory or literal value"},
			{Text: "deployment", Description: "Create a deployment with the specified name."},
			{Text: "namespace", Description: "Create a namespace with the specified name"},
			{Text: "quota", Description: "Create a quota with the specified name."},
			{Text: "secret", Description: "Create a secret using specified subcommand"},
			{Text: "service", Description: "Create a service using specified subcommand."},
			{Text: "serviceaccount", Description: "Create a service account with the specified name"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subcommands, args[1], true)
		}
	case "delete":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, second, true)
		}

	// 	third := args[2]
	// 	if len(args) == 3 {
	// 		switch second {
	// 		case "componentstatuses", "cs":
	// 			return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
	// 		case "configmaps", "cm":
	// 			return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
	// 		case "daemonsets", "ds":
	// 			return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
	// 		case "deploy", "deployments":
	// 			return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
	// 		case "endpoints", "ep":
	// 			return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
	// 		case "ingresses", "ing":
	// 			return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
	// 		case "limitranges", "limits":
	// 			return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
	// 		case "namespaces", "ns":
	// 			return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
	// 		case "no", "nodes":
	// 			return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
	// 		case "po", "pod", "pods":
	// 			return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
	// 		case "persistentvolumeclaims", "pvc":
	// 			return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
	// 		case "persistentvolumes", "pv":
	// 			return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
	// 		case "podsecuritypolicies", "psp":
	// 			return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
	// 		case "podtemplates":
	// 			return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
	// 		case "replicasets", "rs":
	// 			return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
	// 		case "replicationcontrollers", "rc":
	// 			return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
	// 		case "resourcequotas", "quota":
	// 			return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
	// 		case "secrets":
	// 			return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
	// 		case "sa", "serviceaccounts":
	// 			return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
	// 		case "svc", "services":
	// 			return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
	// 		case "job", "jobs":
	// 			return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
	// 		}
	// 	}
	case "edit":
		if len(args) == 2 {
			return prompt.FilterHasPrefix(resourceTypes, args[1], true)
		}

	// 	if len(args) == 3 {
	// 		third := args[2]
	// 		switch args[1] {
	// 		case "componentstatuses", "cs":
	// 			return prompt.FilterContains(getComponentStatusCompletions(c.client), third, true)
	// 		case "configmaps", "cm":
	// 			return prompt.FilterContains(getConfigMapSuggestions(c.client, namespace), third, true)
	// 		case "daemonsets", "ds":
	// 			return prompt.FilterContains(getDaemonSetSuggestions(c.client, namespace), third, true)
	// 		case "deploy", "deployments":
	// 			return prompt.FilterContains(getDeploymentSuggestions(c.client, namespace), third, true)
	// 		case "endpoints", "ep":
	// 			return prompt.FilterContains(getEndpointsSuggestions(c.client, namespace), third, true)
	// 		case "ingresses", "ing":
	// 			return prompt.FilterContains(getIngressSuggestions(c.client, namespace), third, true)
	// 		case "limitranges", "limits":
	// 			return prompt.FilterContains(getLimitRangeSuggestions(c.client, namespace), third, true)
	// 		case "namespaces", "ns":
	// 			return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), third, true)
	// 		case "no", "nodes":
	// 			return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
	// 		case "po", "pod", "pods":
	// 			return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
	// 		case "persistentvolumeclaims", "pvc":
	// 			return prompt.FilterContains(getPersistentVolumeClaimSuggestions(c.client, namespace), third, true)
	// 		case "persistentvolumes", "pv":
	// 			return prompt.FilterContains(getPersistentVolumeSuggestions(c.client), third, true)
	// 		case "podsecuritypolicies", "psp":
	// 			return prompt.FilterContains(getPodSecurityPolicySuggestions(c.client), third, true)
	// 		case "podtemplates":
	// 			return prompt.FilterContains(getPodTemplateSuggestions(c.client, namespace), third, true)
	// 		case "replicasets", "rs":
	// 			return prompt.FilterContains(getReplicaSetSuggestions(c.client, namespace), third, true)
	// 		case "replicationcontrollers", "rc":
	// 			return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), third, true)
	// 		case "resourcequotas", "quota":
	// 			return prompt.FilterContains(getResourceQuotasSuggestions(c.client, namespace), third, true)
	// 		case "secrets":
	// 			return prompt.FilterContains(getSecretSuggestions(c.client, namespace), third, true)
	// 		case "sa", "serviceaccounts":
	// 			return prompt.FilterContains(getServiceAccountSuggestions(c.client, namespace), third, true)
	// 		case "svc", "services":
	// 			return prompt.FilterContains(getServiceSuggestions(c.client, namespace), third, true)
	// 		case "job", "jobs":
	// 			return prompt.FilterContains(getJobSuggestions(c.client, namespace), third, true)
	// 		}
	// 	}

	// case "namespace":
	// 	if len(args) == 2 {
	// 		return prompt.FilterContains(getNameSpaceSuggestions(c.namespaceList), args[1], true)
	// 	}
	// case "logs":
	// 	if len(args) == 2 {
	// 		return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
	// 	}
	// case "rolling-update", "rollingupdate":
	// 	if len(args) == 2 {
	// 		return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), args[1], true)
	// 	} else if len(args) == 3 {
	// 		return prompt.FilterContains(getReplicationControllerSuggestions(c.client, namespace), args[2], true)
	// 	}
	// case "scale", "resize":
	// 	if len(args) == 2 {
	// 		// Deployment, ReplicaSet, Replication Controller, or Job.
	// 		r := getDeploymentSuggestions(c.client, namespace)
	// 		r = append(r, getReplicaSetSuggestions(c.client, namespace)...)
	// 		r = append(r, getReplicationControllerSuggestions(c.client, namespace)...)
	// 		return prompt.FilterContains(r, args[1], true)
	// 	}
	case "cordon":
		fallthrough
	case "drain":
		fallthrough
	// case "uncordon":
	// 	if len(args) == 2 {
	// 		return prompt.FilterHasPrefix(getNodeSuggestions(c.client), args[1], true)
	// 	}
	// case "attach":
	// 	if len(args) == 2 {
	// 		return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
	// 	}
	// case "exec":
	// 	if len(args) == 2 {
	// 		return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
	// 	}
	// case "port-forward":
	// 	if len(args) == 2 {
	// 		return prompt.FilterContains(getPodSuggestions(c.client, namespace), args[1], true)
	// 	}
	// 	if len(args) == 3 {
	// 		return prompt.FilterHasPrefix(getPortsFromPodName(namespace, args[1]), args[2], true)
	// 	}
	case "rollout":
		subCommands := []prompt.Suggest{
			{Text: "history", Description: "view rollout history"},
			{Text: "pause", Description: "Mark the provided resource as paused"},
			{Text: "resume", Description: "Resume a paused resource"},
			{Text: "undo", Description: "undoes a previous rollout"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "annotate":
	case "config":
		subCommands := []prompt.Suggest{
			{Text: "current-context", Description: "Displays the current-context"},
			{Text: "delete-cluster", Description: "Delete the specified cluster from the kubeconfig"},
			{Text: "delete-context", Description: "Delete the specified context from the kubeconfig"},
			{Text: "get-clusters", Description: "Display clusters defined in the kubeconfig"},
			{Text: "get-contexts", Description: "Describe one or many contexts"},
			{Text: "set", Description: "Sets an individual value in a kubeconfig file"},
			{Text: "set-cluster", Description: "Sets a cluster entry in kubeconfig"},
			{Text: "set-context", Description: "Sets a context entry in kubeconfig"},
			{Text: "set-credentials", Description: "Sets a user entry in kubeconfig"},
			{Text: "unset", Description: "Unsets an individual value in a kubeconfig file"},
			{Text: "use-context", Description: "Sets the current-context in a kubeconfig file"},
			{Text: "view", Description: "Display merged kubeconfig settings or a specified kubeconfig file"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
		// if len(args) == 3 {
		// 	third := args[2]
		// 	switch args[1] {
		// 	case "use-context":
		// 		return prompt.FilterContains(getContextSuggestions(), third, true)
		// 	}
		// }
	case "cluster-info":
		subCommands := []prompt.Suggest{
			{Text: "dump", Description: "Dump lots of relevant info for debugging and diagnosis"},
		}
		if len(args) == 2 {
			return prompt.FilterHasPrefix(subCommands, args[1], true)
		}
	case "explain":
		return prompt.FilterHasPrefix(resourceTypes, args[1], true)
	case "top":
		second := args[1]
		if len(args) == 2 {
			subcommands := []prompt.Suggest{
				{Text: "nodes"},
				{Text: "pod"},
				// aliases
				{Text: "no"},
				{Text: "po"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}

		// third := args[2]
		// if len(args) == 3 {
		// 	switch second {
		// 	case "no", "node", "nodes":
		// 		return prompt.FilterContains(getNodeSuggestions(c.client), third, true)
		// 	case "po", "pod", "pods":
		// 		return prompt.FilterContains(getPodSuggestions(c.client, namespace), third, true)
		// 	}
		// }
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}

/* Option arguments */

// var yamlFileCompleter = completer.FilePathCompleter{
// 	IgnoreCase: true,
// 	Filter: func(fi os.FileInfo) bool {
// 		if fi.IsDir() {
// 			return true
// 		}
// 		if strings.HasSuffix(fi.Name(), ".yaml") || strings.HasSuffix(fi.Name(), ".yml") {
// 			return true
// 		}
// 		return false
// 	},
// }

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	filtered := make([]string, 0, l)

	shouldSkipNext := []string{
		"-f",
		"--filename",
		"-n",
		"--namespace",
		"-s",
		"--server",
		"--kubeconfig",
		"--cluster",
		"--user",
		"--output",
		"-o",
	}

	var skipNextArg bool
	for i := 0; i < len(args); i++ {
		if skipNextArg {
			skipNextArg = false
			continue
		}

		for _, s := range shouldSkipNext {
			if strings.HasPrefix(args[i], s) {
				if strings.Contains(args[i], "=") {
					// we can specify option value like '-o=json'
					skipNextArg = false
				} else {
					skipNextArg = true
				}
				continue
			}
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}

		filtered = append(filtered, args[i])
	}
	return filtered, skipNextArg
}

func checkNamespaceArg(d prompt.Document) string {
	args := strings.Split(d.Text, " ")
	var found bool
	for i := 0; i < len(args); i++ {
		if found {
			return args[i]
		}
		if args[i] == "--namespace" || args[i] == "-n" {
			found = true
			continue
		}
	}
	return ""
}

/* NameSpaces */

func getNameSpaceSuggestions(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "default", Description: "Store the username and age"},
		{Text: "ares", Description: "Store the article text posted by user"},
		{Text: "kube-system", Description: "Store the text commented to articles"},
		{Text: "monitor", Description: "Combine users with specific rules"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func getPreviousOption(d prompt.Document) (cmd, option string, found bool) {
	args := strings.Split(d.TextBeforeCursor(), " ")
	l := len(args)
	if l >= 2 {
		option = args[l-2]
	}
	if strings.HasPrefix(option, "-") {
		return args[0], option, true
	}
	return "", "", false
}

func completeOptionArguments(d prompt.Document) ([]prompt.Suggest, bool) {
	cmd, option, found := getPreviousOption(d)
	if !found {
		return []prompt.Suggest{}, false
	}
	switch cmd {
	case "get", "describe", "create", "delete", "replace", "patch",
		"edit", "apply", "expose", "rolling-update", "rollout",
		"label", "annotate", "scale", "convert", "autoscale", "top",
		"root":
		switch option {
		// case "-f", "--filename":
		// 	return yamlFileCompleter.Complete(d), true
		case "-n", "--namespace":
			return prompt.FilterHasPrefix(
				getNameSpaceSuggestions(d),
				d.GetWordBeforeCursor(),
				true,
			), true
		}
	}
	return []prompt.Suggest{}, false
}

func executor(in string) {
	fmt.Println("Your input: " + in)
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	} else if in == "root admin" {
		Run()
	}
	LivePrefixState.LivePrefix = in + "> "
	LivePrefixState.IsEnable = true
	fmt.Println("executor")
}

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	switch commandArgs[0] {
	case "root":
		suggests = rootOptions
	case "get":
		suggests = getOptions
	case "describe":
		suggests = describeOptions
	case "create":
		suggests = createOptions
	case "replace":
		suggests = replaceOptions
	case "patch":
		suggests = patchOptions
	case "delete":
		suggests = deleteOptions
	case "edit":
		suggests = editOptions
	case "apply":
		suggests = applyOptions
	case "logs":
		suggests = logsOptions
	case "rolling-update":
		suggests = rollingUpdateOptions
	case "scale", "resize":
		suggests = scaleOptions
	case "attach":
		suggests = attachOptions
	case "exec":
		suggests = execOptions
	case "port-forward":
		suggests = portForwardOptions
	case "proxy":
		suggests = proxyOptions
	case "run", "run-container":
		suggests = runOptions
	case "expose":
		suggests = exposeOptions
	case "autoscale":
		suggests = autoscaleOptions
	case "rollout":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "history":
				suggests = rolloutHistoryOptions
			case "pause":
				suggests = rolloutPauseOptions
			case "resume":
				suggests = rolloutResumeOptions
			case "status":
				suggests = rolloutStatusOptions
			case "undo":
				suggests = rolloutUndoOptions
			}
		}
	case "label":
		suggests = labelOptions
	case "cluster-info":
		suggests = clusterInfoOptions
	case "explain":
		suggests = explainOptions
	case "cordon":
		suggests = cordonOptions
	case "drain":
		suggests = drainOptions
	case "uncordon":
		suggests = uncordonOptions
	case "annotate":
		suggests = annotateOptions
	case "convert":
		suggests = convertOptions
	case "top":
		if len(commandArgs) >= 2 {
			switch commandArgs[1] {
			case "no", "node", "nodes":
				suggests = topNodeOptions
			case "po", "pod", "pods":
				suggests = topPodOptions
			}
		}
	case "config":
		if len(commandArgs) == 2 {
			switch commandArgs[1] {
			case "get-contexts":
				suggests = configGetContextsOptions
			case "view":
				suggests = configViewOptions
			case "set-cluster":
				suggests = configSetClusterOptions
			case "set-credentials":
				suggests = configSetCredentialsOptions
			case "set":
				suggests = configSetOptions
			}
		}
	default:
		suggests = optionHelp
	}

	suggests = append(suggests, globalOptions...)
	if long {
		return prompt.FilterContains(
			prompt.FilterHasPrefix(suggests, "--", false),
			strings.TrimLeft(args[l-1], "--"),
			true,
		)
	}
	return prompt.FilterContains(suggests, strings.TrimLeft(args[l-1], "-"), true)
}

var optionHelp = []prompt.Suggest{
	{Text: "-h"},
	{Text: "--help"},
}

var globalOptions = []prompt.Suggest{
	{Text: "--namespace", Description: "temporarily set the namespace for a request"},
	{Text: "-n", Description: "temporarily set the namespace for a request"},
	{Text: "--server", Description: "specify the address and port of the Kubernetes API server"},
	{Text: "-s", Description: "specify the address and port of the Kubernetes API server"},
	{Text: "--user", Description: "take the user if this flag exists."},
	{Text: "--cluster", Description: "take the cluster if this flag exists."},
}

func completer(in prompt.Document) []prompt.Suggest {
	// fmt.Println(fmt.Sprintf("Before %s After %s GetWordBeforeCursor %s", in.TextBeforeCursor(), in.TextAfterCursor(), in.GetWordBeforeCursor()))
	// if in.TextBeforeCursor() == "" {
	// 	return []prompt.Suggest{}
	// }
	// s := []prompt.Suggest{
	// 	{Text: "users", Description: "Store the username and age"},
	// 	{Text: "articles", Description: "Store the article text posted by user"},
	// 	{Text: "comments", Description: "Store the text commented to articles"},
	// 	{Text: "groups", Description: "Combine users with specific rules"},
	// }
	// fmt.Println("completer")
	// return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)

	// 如果输入值为空 返回空字符串
	if in.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	// 	获取所有输入字符串并以空格分割
	args := strings.Split(in.TextBeforeCursor(), " ")
	// 获取当前输入字符
	current := in.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(current, "-") {
		return optionCompleter(args, strings.HasPrefix(current, "--"))
	}

	// Return suggestions for option
	if suggests, found := completeOptionArguments(in); found {
		return suggests
	}

	namespace := checkNamespaceArg(in)
	if namespace == "" {
		namespace = "default"
	}
	commandArgs, skipNext := excludeOptions(args)
	if skipNext {
		// when type 'get pod -o ', we don't want to complete pods. we want to type 'json' or other.
		// So we need to skip argumentCompleter.
		return []prompt.Suggest{}
	}
	return argumentsCompleter(namespace, commandArgs)
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("live-prefix-example"),
	)
	p.Run()
}

var rootOptions = []prompt.Suggest{
	prompt.Suggest{Text: "root", Description: "嵌套页面"},
	prompt.Suggest{Text: "users", Description: "用户测试页面"},
	prompt.Suggest{Text: "admin", Description: "用户管理界面"},
}

var annotateOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources, including uninitialized ones, in the namespace of the specified resource types."},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to update the annotation"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to update the annotation"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the mainctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the mainctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "--local", Description: "If true, annotation will NOT contact api-server but run locally."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://mainrnetes.io/docs/user-guide/mainctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://mainrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://mainrnetes.io/docs/user-guide/mainctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://mainrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--overwrite", Description: "If true, allow annotations to be overwritten, otherwise reject annotation updates that overwrite existing annotations."},
	prompt.Suggest{Text: "--record", Description: "Record current mainctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--resource-version", Description: "If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)."},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

var applyOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources in the namespace of the specified resource types."},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--cascade", Description: "Only relevant during a prune or a force apply. If true, cascade the deletion of the resources managed by pruned or deleted resources (e.g. Pods created by a ReplicationController)."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files that contains the configuration to apply"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files that contains the configuration to apply"},
	prompt.Suggest{Text: "--force", Description: "Delete and re-create the specified resource, when PATCH encounters conflict and has retried for 5 times."},
	prompt.Suggest{Text: "--grace-period", Description: "Only relevant during a prune or a force apply. Period of time in seconds given to pruned or deleted resources to terminate gracefully. Ignored if negative."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the promptctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the promptctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "--openapi-patch", Description: "If true, use openapi to calculate diff when the openapi presents and the resource can be found in the openapi spec. Otherwise, fall back to use baked-in types."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--overwrite", Description: "Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration"},
	prompt.Suggest{Text: "--prune", Description: "Automatically delete resource objects, including the uninitialized ones, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all."},
	prompt.Suggest{Text: "--prune-whitelist", Description: "Overwrite the default whitelist with <group/version/kind> for --prune"},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--timeout", Description: "Only relevant during a force apply. The length of time to wait before giving up on a delete of the old resource, zero means determine a timeout from the size of the object. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h)."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var attachOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-c", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	prompt.Suggest{Text: "--container", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	prompt.Suggest{Text: "-i", Description: "Pass stdin to the container"},
	prompt.Suggest{Text: "--stdin", Description: "Pass stdin to the container"},
	prompt.Suggest{Text: "-t", Description: "Stdin is a TTY"},
	prompt.Suggest{Text: "--tty", Description: "Stdin is a TTY"},
}

var autoscaleOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--cpu-percent", Description: "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, a default autoscaling policy will be used."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to autoscale."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to autoscale."},
	prompt.Suggest{Text: "--generator", Description: "The name of the API generator to use. Currently there is only 1 generator."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--max", Description: "The upper limit for the number of pods that can be set by the autoscaler. Required."},
	prompt.Suggest{Text: "--min", Description: "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value."},
	prompt.Suggest{Text: "--name", Description: "The name for the newly created object. If not specified, the name of the input resource will be used."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform promptctl apply on this object in the future."},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

var clusterInfoOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
}

var configGetContextsOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
}

var configSetClusterOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--embed-certs", Description: "embed-certs for the cluster entry in promptconfig"},
}

var configSetCredentialsOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--auth-provider", Description: "Auth provider for the user entry in promptconfig"},
	prompt.Suggest{Text: "--auth-provider-arg", Description: "'key=value' arguments for the auth provider"},
	prompt.Suggest{Text: "--embed-certs", Description: "Embed client cert/key for the user entry in promptconfig"},
}

var configSetOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--set-raw-bytes", Description: "When writing a []byte PROPERTY_VALUE, write the given string directly without base64 decoding."},
}

var configViewOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--flatten", Description: "Flatten the resulting promptconfig file into self-contained output (useful for creating portable promptconfig files)"},
	prompt.Suggest{Text: "--merge", Description: "Merge the full hierarchy of promptconfig files"},
	prompt.Suggest{Text: "--minify", Description: "Remove all information not used by current-context from the output"},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--raw", Description: "Display raw byte data"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

var convertOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to need to get converted."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to need to get converted."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--local", Description: "If true, convert will NOT try to contact api-server but run locally."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output-version", Description: "Output the formatted object with the given group version (for ex: 'extensions/v1beta1').)"},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var cordonOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on"},
}

var createOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--edit", Description: "Edit the API resource before creating"},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to use to create the resource"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to use to create the resource"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--raw", Description: "Raw URI to POST to the server.  Uses the transport specified by the promptconfig file."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform promptctl apply on this object in the future."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
	prompt.Suggest{Text: "--windows-line-endings", Description: "Only relevant if --edit=true. Defaults to the line ending native to your platform."},
}

var deleteOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Delete all resources, including uninitialized ones, in the namespace of the specified resource types."},
	prompt.Suggest{Text: "--cascade", Description: "If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController).  Default true."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files containing the resource to delete."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files containing the resource to delete."},
	prompt.Suggest{Text: "--force", Description: "Immediate deletion of some resources may result in inconsistency or data loss and requires confirmation."},
	prompt.Suggest{Text: "--grace-period", Description: "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative."},
	prompt.Suggest{Text: "--ignore-not-found", Description: "Treat \"resource not found\" as a successful delete. Defaults to \"true\" when --all is specified."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the promptctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the promptctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "--now", Description: "If true, resources are signaled for immediate shutdown (same as --grace-period=1)."},
	prompt.Suggest{Text: "-o", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--output", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, not including uninitialized ones."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, not including uninitialized ones."},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object"},
}

var describeOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files containing the resource to describe"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files containing the resource to describe"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the promptctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the promptctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--show-events", Description: "If true, display events related to the described object."},
}

var drainOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--delete-local-data", Description: "Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained)."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--force", Description: "Continue even if there are pods not managed by a ReplicationController, ReplicaSet, Job, DaemonSet or StatefulSet."},
	prompt.Suggest{Text: "--grace-period", Description: "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used."},
	prompt.Suggest{Text: "--ignore-daemonsets", Description: "Ignore DaemonSet-managed pods."},
	prompt.Suggest{Text: "--pod-selector", Description: "Label selector to filter pods on the node"},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up, zero means infinite"},
}

var editOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to use to edit the resource"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to use to edit the resource"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the promptctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the promptctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: yaml|json."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: yaml|json."},
	prompt.Suggest{Text: "--output-patch", Description: "Output the patch if the resource is edited."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform promptctl apply on this object in the future."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
	prompt.Suggest{Text: "--windows-line-endings", Description: "Defaults to the line ending native to your platform."},
}

var execOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-c", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	prompt.Suggest{Text: "--container", Description: "Container name. If omitted, the first container in the pod will be chosen"},
	prompt.Suggest{Text: "-p", Description: "Pod name"},
	prompt.Suggest{Text: "--pod", Description: "Pod name"},
	prompt.Suggest{Text: "-i", Description: "Pass stdin to the container"},
	prompt.Suggest{Text: "--stdin", Description: "Pass stdin to the container"},
	prompt.Suggest{Text: "-t", Description: "Stdin is a TTY"},
	prompt.Suggest{Text: "--tty", Description: "Stdin is a TTY"},
}

var explainOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--api-version", Description: "Get different explanations for particular API version"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--recursive", Description: "Print the fields of fields (Currently only 1 level deep)"},
}

var exposeOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--cluster-ip", Description: "ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service."},
	prompt.Suggest{Text: "--container-port", Description: "Synonym for --target-port"},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--external-ip", Description: "Additional external IP address (not managed by promptrnetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to expose a service"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to expose a service"},
	prompt.Suggest{Text: "--generator", Description: "The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'."},
	prompt.Suggest{Text: "-l", Description: "Labels to apply to the service created by this call."},
	prompt.Suggest{Text: "--labels", Description: "Labels to apply to the service created by this call."},
	prompt.Suggest{Text: "--load-balancer-ip", Description: "IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used (cloud-provider specific)."},
	prompt.Suggest{Text: "--name", Description: "The name for the newly created object."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--overrides", Description: "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field."},
	prompt.Suggest{Text: "--port", Description: "The port that the service should serve on. Copied from the resource being exposed, if unspecified"},
	prompt.Suggest{Text: "--protocol", Description: "The network protocol for the service to be created. Default is 'TCP'."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform promptctl apply on this object in the future."},
	prompt.Suggest{Text: "--selector", Description: "A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.)"},
	prompt.Suggest{Text: "--session-affinity", Description: "If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--target-port", Description: "Name or number for the port on the container that the service should direct traffic to. Optional."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--type", Description: "Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'."},
}
var getOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--chunk-size", Description: "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future."},
	prompt.Suggest{Text: "--experimental-server-print", Description: "If true, have the server return the appropriate table output. Supports extension APIs and CRD. Experimental."},
	prompt.Suggest{Text: "--export", Description: "If true, use 'export' for the resources.  Exported resources are stripped of cluster-specific information."},
	prompt.Suggest{Text: "--field-selector", Description: "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--ignore-not-found", Description: "If the requested object does not exist the command will return exit code 0."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the promptctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the promptctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "-L", Description: "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2..."},
	prompt.Suggest{Text: "--label-columns", Description: "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2..."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--raw", Description: "Raw URI to request from the server.  Uses the transport specified by the promptconfig file."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-kind", Description: "If present, list the resource type for the requested object(s)."},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--use-openapi-print-columns", Description: "If true, use x-promptrnetes-print-column metadata (if present) from the OpenAPI schema for displaying a resource."},
	prompt.Suggest{Text: "-w", Description: "After listing/getting the requested object, watch for changes. Uninitialized objects are excluded if no object name is provided."},
	prompt.Suggest{Text: "--watch", Description: "After listing/getting the requested object, watch for changes. Uninitialized objects are excluded if no object name is provided."},
	prompt.Suggest{Text: "--watch-only", Description: "Watch for changes to the requested object(s), without listing/getting first."},
}

var labelOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources, including uninitialized ones, in the namespace of the specified resource types"},
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to update the labels"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to update the labels"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--include-uninitialized", Description: "If true, the promptctl command applies to uninitialized objects. If explicitly set to false, this flag overrides other flags that make the promptctl commands apply to uninitialized objects, e.g., \"--all\". Objects with empty metadata.initializers are regarded as initialized."},
	prompt.Suggest{Text: "--list", Description: "If true, display the labels for a given resource."},
	prompt.Suggest{Text: "--local", Description: "If true, label will NOT contact api-server but run locally."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--overwrite", Description: "If true, allow labels to be overwritten, otherwise reject label updates that overwrite existing labels."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--resource-version", Description: "If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)."},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}

var logsOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-c", Description: "Print the logs of this container"},
	prompt.Suggest{Text: "--container", Description: "Print the logs of this container"},
	prompt.Suggest{Text: "-f", Description: "Specify if the logs should be streamed."},
	prompt.Suggest{Text: "--follow", Description: "Specify if the logs should be streamed."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--interactive", Description: "If true, prompt the user for input when required."},
	prompt.Suggest{Text: "--limit-bytes", Description: "Maximum bytes of logs to return. Defaults to no limit."},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	prompt.Suggest{Text: "-p", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	prompt.Suggest{Text: "--previous", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on."},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on."},
	prompt.Suggest{Text: "--since", Description: "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used."},
	prompt.Suggest{Text: "--since-time", Description: "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used."},
	prompt.Suggest{Text: "--tail", Description: "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided."},
	prompt.Suggest{Text: "--timestamps", Description: "Include timestamps on each line in the log output"},
}

var patchOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to update"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to update"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--local", Description: "If true, patch will operate on the content of the file, not the server-side resource."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "-p", Description: "The patch to be applied to the resource JSON file."},
	prompt.Suggest{Text: "--patch", Description: "The patch to be applied to the resource JSON file."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--type", Description: "The type of patch being provided; one of [json merge strategic]"},
}

var portForwardOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
}

var proxyOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--accept-hosts", Description: "Regular expression for hosts that the proxy should accept."},
	prompt.Suggest{Text: "--accept-paths", Description: "Regular expression for paths that the proxy should accept."},
	prompt.Suggest{Text: "--address", Description: "The IP address on which to serve on."},
	prompt.Suggest{Text: "--api-prefix", Description: "Prefix to serve the proxied API under."},
	prompt.Suggest{Text: "--disable-filter", Description: "If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to XSRF attacks, when used with an accessible port."},
	prompt.Suggest{Text: "-p", Description: "The port on which to run the proxy. Set to 0 to pick a random port."},
	prompt.Suggest{Text: "--port", Description: "The port on which to run the proxy. Set to 0 to pick a random port."},
	prompt.Suggest{Text: "--reject-methods", Description: "Regular expression for HTTP methods that the proxy should reject (example --reject-methods='POST,PUT,PATCH')."},
	prompt.Suggest{Text: "--reject-paths", Description: "Regular expression for paths that the proxy should reject. Paths specified here will be rejected even accepted by --accept-paths."},
	prompt.Suggest{Text: "-u", Description: "Unix socket on which to run the proxy."},
	prompt.Suggest{Text: "--unix-socket", Description: "Unix socket on which to run the proxy."},
	prompt.Suggest{Text: "-w", Description: "Also serve static files from the given directory under the specified prefix."},
	prompt.Suggest{Text: "--www", Description: "Also serve static files from the given directory under the specified prefix."},
	prompt.Suggest{Text: "-P", Description: "Prefix to serve static files under, if static file directory is specified."},
	prompt.Suggest{Text: "--www-prefix", Description: "Prefix to serve static files under, if static file directory is specified."},
}

var replaceOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--cascade", Description: "Only relevant during a force replace. If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController)."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to use to replace the resource."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to use to replace the resource."},
	prompt.Suggest{Text: "--force", Description: "Delete and re-create the specified resource"},
	prompt.Suggest{Text: "--grace-period", Description: "Only relevant during a force replace. Period of time in seconds given to the old resource to terminate gracefully. Ignored if negative."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "-o", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--output", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform promptctl apply on this object in the future."},
	prompt.Suggest{Text: "--timeout", Description: "Only relevant during a force replace. The length of time to wait before giving up on a delete of the old resource, zero means determine a timeout from the size of the object. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h)."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var rollingUpdateOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--container", Description: "Container name which will have its image upgraded. Only relevant when --image is specified, ignored otherwise. Required when using --image on a multi-container pod"},
	prompt.Suggest{Text: "--deployment-label-key", Description: "The key to use to differentiate between two different controllers, default 'deployment'.  Only relevant when --image is specified, ignored otherwise"},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename or URL to file to use to create the new replication controller."},
	prompt.Suggest{Text: "--filename", Description: "Filename or URL to file to use to create the new replication controller."},
	prompt.Suggest{Text: "--image", Description: "Image to use for upgrading the replication controller. Must be distinct from the existing image (either new image or new image tag).  Can not be used with --filename/-f"},
	prompt.Suggest{Text: "--image-pull-policy", Description: "Explicit policy for when to pull container images. Required when --image is same as existing image, ignored otherwise."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--poll-interval", Description: "Time delay between polling for replication controller status after the update. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	prompt.Suggest{Text: "--rollback", Description: "If true, this is a request to abort an existing rollout that is partially rolled out. It effectively reverses current and next and runs a rollout"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "--timeout", Description: "Max time to wait for a replication controller to update before giving up. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	prompt.Suggest{Text: "--update-period", Description: "Time to wait between updating pods. Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}

var rolloutHistoryOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--revision", Description: "See the details, including podTemplate of the revision specified"},
}

var rolloutPauseOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
}

var rolloutResumeOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
}

var rolloutStatusOptions = []prompt.Suggest{
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--revision", Description: "Pin to a specific revision for showing its status. Defaults to 0 (last revision)."},
	prompt.Suggest{Text: "-w", Description: "Watch the status of the rollout until it's done."},
	prompt.Suggest{Text: "--watch", Description: "Watch the status of the rollout until it's done."},
}

var rolloutUndoOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--to-revision", Description: "The revision to rollback to. Default to 0 (last revision)."},
}

var runOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	prompt.Suggest{Text: "--attach", Description: "If true, wait for the Pod to start running, and then attach to the Pod as if 'promptctl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true. With '--restart=Never' the exit code of the container process is returned."},
	prompt.Suggest{Text: "--command", Description: "If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default."},
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "--env", Description: "Environment variables to set in the container"},
	prompt.Suggest{Text: "--expose", Description: "If true, a public, external service is created for the container(s) which are run"},
	prompt.Suggest{Text: "--generator", Description: "The name of the API generator to use, see http://promptrnetes.io/docs/user-guide/promptctl-conventions/#generators for a list."},
	prompt.Suggest{Text: "--hostport", Description: "The host port mapping for the container port. To demonstrate a single-machine container."},
	prompt.Suggest{Text: "--image", Description: "The image for the container to run."},
	prompt.Suggest{Text: "--image-pull-policy", Description: "The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "-l", Description: "Comma separated labels to apply to the pod(s). Will override previous values."},
	prompt.Suggest{Text: "--labels", Description: "Comma separated labels to apply to the pod(s). Will override previous values."},
	prompt.Suggest{Text: "--leave-stdin-open", Description: "If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes."},
	prompt.Suggest{Text: "--limits", Description: "The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges."},
	prompt.Suggest{Text: "--no-headers", Description: "When using the default or custom-column output format, don't print headers (default print headers)."},
	prompt.Suggest{Text: "-o", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--output", Description: "Output format. One of: json|yaml|wide|name|custom-columns=...|custom-columns-file=...|go-template=...|go-template-file=...|jsonpath=...|jsonpath-file=... See custom columns [http://promptrnetes.io/docs/user-guide/promptctl-overview/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [http://promptrnetes.io/docs/user-guide/jsonpath]."},
	prompt.Suggest{Text: "--overrides", Description: "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field."},
	prompt.Suggest{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	prompt.Suggest{Text: "--port", Description: "The port that this container exposes.  If --expose is true, this is also the port used by the service that is created."},
	prompt.Suggest{Text: "--quiet", Description: "If true, suppress prompt messages."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-r", Description: "Number of replicas to create for this container. Default is 1."},
	prompt.Suggest{Text: "--replicas", Description: "Number of replicas to create for this container. Default is 1."},
	prompt.Suggest{Text: "--requests", Description: "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges."},
	prompt.Suggest{Text: "--restart", Description: "The restart policy for this Pod.  Legal values [Always, OnFailure, Never].  If set to 'Always' a deployment is created, if set to 'OnFailure' a job is created, if set to 'Never', a regular pod is created. For the latter two --replicas must be 1.  Default 'Always', for CronJobs `Never`."},
	prompt.Suggest{Text: "--rm", Description: "If true, delete resources created in this command for attached containers."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform promptctl apply on this object in the future."},
	prompt.Suggest{Text: "--schedule", Description: "A schedule in the Cron format the job should be run with."},
	prompt.Suggest{Text: "--service-generator", Description: "The name of the generator to use for creating a service.  Only used if --expose is true"},
	prompt.Suggest{Text: "--service-overrides", Description: "An inline JSON override for the generated service object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.  Only used if --expose is true."},
	prompt.Suggest{Text: "--serviceaccount", Description: "Service account to set in the pod spec"},
	prompt.Suggest{Text: "-a", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-all", Description: "When printing, show all resources (default show all pods including terminated one.)"},
	prompt.Suggest{Text: "--show-labels", Description: "When printing, show all labels as the last column (default hide labels column)"},
	prompt.Suggest{Text: "--sort-by", Description: "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string."},
	prompt.Suggest{Text: "-i", Description: "Keep stdin open on the container(s) in the pod, even if nothing is attached."},
	prompt.Suggest{Text: "--stdin", Description: "Keep stdin open on the container(s) in the pod, even if nothing is attached."},
	prompt.Suggest{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
	prompt.Suggest{Text: "-t", Description: "Allocated a TTY for each container in the pod."},
	prompt.Suggest{Text: "--tty", Description: "Allocated a TTY for each container in the pod."},
}

var scaleOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all", Description: "Select all resources in the namespace of the specified resource types"},
	prompt.Suggest{Text: "--current-replicas", Description: "Precondition for current size. Requires that the current size of the resource match this value in order to scale."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to set a new size"},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to set a new size"},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "-o", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--output", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--record", Description: "Record current promptctl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--replicas", Description: "The new desired number of replicas. Required."},
	prompt.Suggest{Text: "--resource-version", Description: "Precondition for resource version. Requires that the current resource version match this value in order to scale."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--timeout", Description: "The length of time to wait before giving up on a scale operation, zero means don't wait. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h)."},
}

var topNodeOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in"},
	prompt.Suggest{Text: "--heapster-port", Description: "Port name in service to use"},
	prompt.Suggest{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as"},
	prompt.Suggest{Text: "--heapster-service", Description: "Name of Heapster service"},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
}

var topPodOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--all-namespaces", Description: "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace."},
	prompt.Suggest{Text: "--containers", Description: "If present, print usage of containers within a pod."},
	prompt.Suggest{Text: "--heapster-namespace", Description: "Namespace Heapster service is located in"},
	prompt.Suggest{Text: "--heapster-port", Description: "Port name in service to use"},
	prompt.Suggest{Text: "--heapster-scheme", Description: "Scheme (http or https) to connect to Heapster as"},
	prompt.Suggest{Text: "--heapster-service", Description: "Name of Heapster service"},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)"},
}

var uncordonOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	prompt.Suggest{Text: "-l", Description: "Selector (label query) to filter on"},
	prompt.Suggest{Text: "--selector", Description: "Selector (label query) to filter on"},
}
