### Google Cloud account

If you don't have a Google Cloud account, grab a credit card and go to https://cloud.google.com/. Then, click the “Get started for free” button. Finally, choose whether you want a business account or an individual one.

Next step: On GCP, select an existing project or create a new one.

#### Enable the GKE API

(You can skip this step if GKE API is already enabled for this project)

With a GCP Project selected or created,

- Head to the [Google Developer Console](https://console.developers.google.com/apis/api/container.googleapis.com/overview).
- Enable the 'Kubernetes Engine API'.
- Enable the 'Cloud Resource Manager API'
- Enable the 'Compute Engine API'
- Enable the 'IAM Service Account Credentials API'

Alternatively you can enable these from the [gcloud](https://cloud.google.com/sdk/gcloud) command line;

```shell
# Setup if required
gcloud auth login (assuming you've not authenticated)
gcloud config set project <project_id>

# Enable the APIs
gcloud services enable cloudresourcemanager.googleapis.com
gcloud services enable iam.googleapis.com
gcloud services enable compute.googleapis.com
gcloud services enable container.googleapis.com
```

#### Create a Service Account

(You can skip this step if you already have a Service Account setup)

With the a GCP Project selected or created,

- Head to the [IAM Console](https://console.cloud.google.com/iam-admin/serviceaccounts).
- Click `Create service account`.
- Fill in the form with details with your team's service account.

#### Configure your Service Account permissions

(You can skip this step if you're Service Account has the `Owner` role)

- Assign the `Owner` role to your Service account.

#### Create a key and download it (as JSON)

(You can skip this step if you already have your Service Account key downloaded in JSON format)

Kore will use this key to access the Service Account.

This is the last step, create a key and download it in JSON format.

