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

## Dev mode

### Run elastic search
```shell script
./gradlew installElasticsearchChart

```

### Run web interface
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
```shell script
./gradlew runWeb
```
