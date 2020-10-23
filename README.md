# cloud-resource-dashboard
Collect information about your cloud infrastructure in one place

# Requirements
You need to have configured Kubernetes cluster

# Running the application

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
