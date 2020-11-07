![Api docker image](https://github.com/Max-Levitskiy/cloud-resource-dashboard/workflows/Api%20docker%20image/badge.svg)

# cloud-resource-dashboard
Collect information about your cloud infrastructure in one place

# Requirements
You need to have configured Kubernetes cluster

# Running the application

## Adding AWS secrets
You need to have aws-credentials secret in the k8s to mount in the API pod.
There is a helper script for it: charts/scripts/aws-credentials-to-k8s.sh 
This script will create aws-credentials secret from your ~/.aws/config and ~/.aws/credentials.

__IMPORTANT: Please check you connected to correct k8s cluster to avoid a leak of your credentials.__

## Adding GCP secrets
### Service accounts creation
Add service account in your [GCP console](https://console.cloud.google.com/iam-admin/serviceaccounts). For scan all resources you should select Project -> Viewer role. If you don't need all resources, just select required services and use viewer option. 

For scanning multiple projects you have two options:
- Create service account in every project
- Create service account in one project and add this role as an [IAM](https://console.cloud.google.com/iam-admin/iam) for other required projects.

### Service account keys to secret

#### Create service accounts 

After you have required service account(s) you need to create keys and create secret from them. On the [service account page](https://console.cloud.google.com/iam-admin/serviceaccounts) use Create Action -> Create key option, select json format and save the file in the ~/.gcp/credentials 
If you're using multiple service account save all keys in ~/.gcp/credentials. Names of the files is not important. 

#### Create secret in k8s

Use charts/scripts/gcp-credentials-to-k8s.sh script for create gcp-credentials secret.
 
 Alternatively, you can create secret manually with k8s command:
 ```shell script
GCPCREDENTIALS=$HOME/.gcp/credentials
FILENAME=fileWithKey.json
kubectl create secret generic gcp-credentials --from-file=$FILENAME=$GCPCREDENTIALS/$FILENAME
```

## From gradle
```shell script
./gradlew installElasticsearchChart
./gradlew helmInstall
```

## As a helm chart (Required helm installed)
From the charts/cloud-resource-dashboard run:
```shell script
helm dependency build
helm install cloud-resource-dashboard .
```

# Development

### Run elastic search (k8s)
```shell script
./gradlew installElasticsearchChart
```

### Run web interface (Using Angular CLI serve command)
```shell script
./gradlew runWeb
```
or (requires nodejs and npm installed)
```shell script
cd web
npm run start
```
or (requires previous + angular cli installed)
```shell script
cd web
ng s
```

### Run API server
It will be run with current AWS account settings
```shell script
./gradlew runWeb
```
