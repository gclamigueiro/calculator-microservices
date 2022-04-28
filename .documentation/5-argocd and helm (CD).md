
# Deploy and update kuberentes resources with ArgoCd and Helm

For application deployment with ArgoCD we need tu use CRDs (Custom Resource Definitions) to configure and manage de release. The two types of CRDs are:

## Project resource

Provides a logical grouping of applications, including access to source and destination repositories, and permissions to resources within the cluster. Was created in [2.1 - Configure ArgoCD](./2.1-configure%20argocd.md)

## Aplication resource

It stores the configuration of how an application should be deployed and managed. Can be configured to use Helm to manage the manifest.

You can check an example in ```./.iaas/argocd/``` folder inside de projects

