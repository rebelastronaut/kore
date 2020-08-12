
# Local Quick Start Guide

In this guide, we'll walk you through how to use the Kore CLI to set up a sandbox team environment locally and deploy a sample application.

We'll showcase how Kore can give you a head start with setting up [clusters](https://www.redhat.com/en/topics/containers/what-is-a-kubernetes-cluster), team members and environments.

**NOTE** The installation of Kore created by the `kore alpha local` command in this quick start is suitable for testing and proof-of-concept work only. Bootstrapping a production installation of Kore is coming soon with [issue/340](https://github.com/appvia/kore/issues/340).

## Getting Started

- [Docker](#docker)
- [Start Kore Locally with CLI](#start-kore-locally-with-cli)
- [Login as Admin with CLI](#login-as-admin-with-cli)
- [Create a Team with CLI](#create-a-team-with-cli)
- [Enable Kore to Set up Team Environments on GKE](enable-kore-to-set-up-team-environments-on-gke)
- [Provision a Sandbox Env with CLI](#provision-a-sandbox-env-with-cli)
- [Deploy An App to the Sandbox](#deploy-an-app-to-the-sandbox)
- [Cleaning Up](#cleaning-up)

### Docker

Please ensure you have the following installed on your machine,

- Docker: installation instructions can be found [here]([https://docs.docker.com/install/](https://docs.docker.com/install/)
- Kubectl: installation instructions can be found [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/)


### Run Kore locally

This should provision a local kubernetes installation and deploy the official helm release.

```shell
$ kore alpha local up
✅ Performing preflight checks for installation
   ◉ Checking for kubectl binary requirement
✅ Passed preflight checks for kore installation
   ◉ Single-sign on is disabled, using kore managed users
   ◉ Local admin not set, generating admin user password
✅ Persisting the values to local file: "/home/jest/.kore/values.yaml"
✅ Performing preflight checks for local cluster provider
   ◉ Kind binary requirement found in $PATH
   ◉ Docker binary requirement found in $PATH
✅ Passed preflight checks for cluster provider
✅ Attempting to build the local kubernetes cluster
   ◉ Checking if kind cluster: "kore" already exists
   ◉ Using Kind image: "kindest/node:v1.16.9"
   ◉ Provisioning a kind cluster: "kore" (usually takes 1-3mins)
   ◉ Still building the kind cluster "kore": 20s
   ◉ Still building the kind cluster "kore": 40s
   ◉ Built local kind cluster in 61s
   ◉ Exporting kubeconfig from kind cluster: "kore"
✅ Exported the kubeconfig from provisioned cluster
✅ Provisioned a local kubernetes cluster
✅ Switched the kubectl context: "kind-kore"
✅ Attempting to deploy the Kore release
   ◉ Using the official Helm chart for deployment
   ◉ Kore Helm chart has been installed at /home/jest/.kore/charts
   ◉ Waiting for kubernetes controlplane to become available
   ◉ Creating the kore namespace in cluster
   ◉ Deploying the kore installation to cluster
✅ Deployed the Kore release into the cluster
✅ Waiting for deployment to rollout successfully (5m0s timeout)
   ◉ Deployed Kore installation to cluster in 104s
✅ Successfully deployed the kore release to cluster

Access the Kore portal via http://localhost:3000 [ admin | VssJHJVQ ]
Configure your CLI via $ kore login -a http://localhost:10080
```

Note: you can now view the UI from http://localhost:3000 _(credentials will be rendered to screen)_, or use the CLI below.

### Login as Admin with CLI

You now have to login to be able to create teams and provision environments.

As you're the only user, you'll be assigned Admin privileges.

```shell script
✔ Please enter the Kore API (e.g https://api.example.com) : http://localhost:10080
Please enter your username : admin
Please confirm password for  : ********
$ kore whoami
```

Note you can also enable single-sign-on for the UI and all clusters; an example of how to configure an IDP provider can be found [here](docs/setup-auth0.md). To enable the feature on the local demo add `kore alpha local up --enable-sso` which will prompt for your OpenID settings _(you can do this as any point)_.

### Create a Team with CLI

Let's create a team with the CLI. In local mode, you'll be assigned as team member to this team.

As a team member, you'll be able to provision environments on behalf of team.

```shell script
$ kore create team --description 'The Appvia product team, working on project Q.' team-appvia
# "team-appvia" team was successfully created
```

To ensure the team was created,

```shell script
$ kore get teams team-appvia
# Name            Description
# team-appvia     The Appvia product team, working on project Q.
```

### Enable Kore to Set up Team Environments on GKE

We now need to give Kore the credentials it needs to build a cluster on our behalf. This command imports a set of credentials into kore
and allows your new team to use them to make clusters.

We'll then use these to create a cluster to host our sandbox environment. You can follow [here](docs/setup-gcp.md] to see how to configure a token, but essentially we need the service account json which has owner in the project.

- GKE Project ID.
- Path to the service account key JSON file.

```shell script
$ kore create gkecredentials gke --description "GKE Credentials" -p <gcp-project-id> --cred-file <path-to-json-service-account> --allocate team-appvia
# Storing credentials in Kore
# Waiting for resource "gke" to provision (you can background with ctrl-c)
# Successfully provisioned the resource: "gke"
# Storing credential allocation in Kore
# Waiting for resource "gke" to provision (you can background with ctrl-c)
# Successfully provisioned the resource: "gke"
```

### Provision a Sandbox Env with CLI

Its time to use the Kore CLI To provision our Sandbox environment,

```shell script
$ kore create cluster appvia-trial -t team-appvia --plan gke-development -a gke --namespaces sandbox
# Attempting to create cluster: "appvia-trial", plan: gke-development
# Waiting for "appvia-trial" to provision (usually takes around 5 minutes, ctrl-c to background)
# Cluster appvia-sdbox has been successfully provisioned
# --> Attempting to create namespace: sandbox

# You can update your kubeconfig via: $ kore kubeconfig -t team-appvia
# Then use 'kubectl' to interact with your team's cluster
```

There's a lot to unpack here. So, lets walk through it,

- `create cluster`, we create a [cluster](https://www.redhat.com/en/topics/containers/what-is-a-kubernetes-cluster) to host our sandbox environment.

- `appvia-trial`, the name of the cluster.

- `-t team-appvia`, the team for which we are creating the sandbox environment.

- `--plan gke-development`, a Kore predefined plan called `gke-development`. This creates a cluster ideal for non-prod use.

- `-a gke`, the `gke` allocated credential to use for creating this cluster.

- `--namespace sandbox`, creates an environment called `sandbox` in the `appvia-trial` where we can deploy our apps, servers, etc..

You now have a sandbox environment locally provisioned for your team. 🎉

### Deploy An App to the Sandbox

We'll be using `kubectl`, the Kubernetes CLI, to make the deployment. If you don't have it already, [please install and setup kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-macos).

Now we have to configure our `kubectl` kubeconfig in ~/.kube/config with our new GKE cluster.

```shell script
./kore kubeconfig -t team-appvia
# Successfully added team [team-appvia] provisioned clusters to your kubeconfig
# Context        Cluster
# appvia-trial   appvia-trial
```

Switch the current `kubectl` context to `appvia-trial`,

```shell script
kubectl config use-context appvia-trial --namespace=sandbox
# + kubectl config use-context appvia-trial --namespace=sandbox
# Switched to context "appvia-trial".
```

Deploy the GKE example web application container available from the Google Cloud Repository

```shell script
kubectl create deployment hello-server --image=gcr.io/google-samples/hello-app:1.0
# + kubectl create deployment hello-server --image=gcr.io/google-samples/hello-app:1.0
# deployment.apps/hello-server created

kubectl expose deployment hello-server --type LoadBalancer --port 80 --target-port 8080
# + kubectl expose deployment hello-server --type LoadBalancer --port 80 --target-port 8080
# service/hello-server exposed
```

Get the `EXTERNAL-IP` for `hello-server` service

```shell script
kubectl get service hello-server
# + kubectl get services
# NAME           TYPE           CLUSTER-IP     EXTERNAL-IP          PORT(S)        AGE
# hello-server   LoadBalancer   10.70.10.119   <35.242.154.199>     80:31319/TCP   23s
```

Now navigate to the `EXTERNAL-IP` as a url

```shell script
open http://35.242.154.199
```

You should see this on the webpage

```text
Hello, world!
Version: 1.0.0
Hostname: hello-server-7f8fd4d44b-hpxls
```

### Cleaning Up

To avoid incurring charges to your Google Cloud account for the resources used in this quickstart, follow these steps.

#### Delete the app from the sandbox environment

```shell script
kubectl delete service hello-server
```

#### Delete the sandbox environment

You can now use kore to destroy the cluster:

```shell script
./kore delete --team team-appvia cluster appvia-trial
# "appvia-trial" was successfully deleted
```

You can check for the cluster deletion completing by retrieving the cluster:

```shell script
./kore get cluster appvia-trial --team team-appvia
# Name            Kind    API Endpoint           Auth Proxy Endpoint    Status
# appvia-trial    GKE     https://1.2.3.4        5.6.7.8                Deleting
```

Once the deletion is complete, the cluster will disappear from Kore:

```shell script
./kore get cluster appvia-trial --team team-appvia
# Error: "appvia-trial" does not exist
```

Finally, after waiting for your cluster to delete, you may stop your local kore environment:

```shell script
./kore local stop
# ...Stopping Kore.
# ...Kore is now stopped.
```
