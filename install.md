### Tekton

#### Pipelines
`kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/previous/v0.48.0/release.yaml`

#### Pipelines as Code
`kubectl apply -f https://raw.githubusercontent.com/openshift-pipelines/pipelines-as-code/stable/release.k8s.yaml`

Patch the PaC service to Loadbalancer so that Metallb assings an external IP. This external IP needs to be used during the GitHub application setup process where we provide the IP as the webhook URL.

`kubectl patch svc pipelines-as-code-controller -n pipelines-as-code -p '{"spec": {"type": "LoadBalancer"}}'`

Create a GitHub Application by following the instaructions from [here](https://pipelinesascode.com/docs/install/github_apps/).

Webhook URL in the github app will be the externalIP that the metallb assigned along with the port 8080 - <external-ip>:8080
