package k8s

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.mondoo.io/mondoo/motor/asset"
	"go.mondoo.io/mondoo/motor/providers"
	"go.mondoo.io/mondoo/motor/providers/k8s"

	batchv1 "k8s.io/api/batch/v1"
)

// ListJobs list all jobs in the cluster.
func ListJobs(p k8s.KubernetesProvider, connection *providers.Config, clusterIdentifier string, namespaceFilter []string) ([]*asset.Asset, error) {
	namespaces, err := p.Namespaces()
	if err != nil {
		return nil, errors.Wrap(err, "could not list kubernetes namespaces")
	}

	jobs := []batchv1.Job{}
	for i := range namespaces {
		namespace := namespaces[i]
		if !isIncluded(namespace.Name, namespaceFilter) {
			log.Info().Str("namespace", namespace.Name).Strs("filter", namespaceFilter).Msg("namespace not included")
			continue
		}

		jobsPerNamespace, err := p.Jobs(namespace)
		if err != nil {
			return nil, errors.Wrap(err, "failed to list Jobs")
		}

		jobs = append(jobs, jobsPerNamespace...)
	}

	assets := []*asset.Asset{}
	for i := range jobs {
		job := jobs[i]
		platformData := p.PlatformInfo()
		platformData.Version = job.APIVersion
		platformData.Build = job.ResourceVersion
		platformData.Labels = map[string]string{
			"namespace": job.Namespace,
			"uid":       string(job.UID),
		}
		platformData.Kind = providers.Kind_KIND_K8S_OBJECT
		asset := &asset.Asset{
			PlatformIds: []string{k8s.NewPlatformWorkloadId(clusterIdentifier, "jobs", job.Namespace, job.Name)},
			Name:        job.Namespace + "/" + job.Name,
			Platform:    platformData,
			Connections: []*providers.Config{connection},
			State:       asset.State_STATE_ONLINE,
			Labels:      job.Labels,
		}
		if asset.Labels == nil {
			asset.Labels = map[string]string{
				"namespace": job.Namespace,
			}
		} else {
			asset.Labels["namespace"] = job.Namespace
		}
		log.Debug().Str("name", job.Name).Str("connection", asset.Connections[0].Host).Msg("resolved Job")

		assets = append(assets, asset)
	}

	return assets, nil
}
