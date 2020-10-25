# cloud-resource-dashboard
Collect information about your cloud infrastructure in one place

# Requirements
You need to have configured Kubernetes cluster

# Running the application

## Adding secrets
You need to have aws-credentials secret in the k8s to mount in the API pod.
There is a helper script for it: charts/scripts/aws-credentials-to-k8s.sh 
This script will create aws-credentials secret from your ~/.aws/config and ~/.aws/credentials.

__IMPORTANT: Please check you connected to correct k8s cluster to avoid a leak of your credentials.__

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
