import org.unbrokendome.gradle.plugins.helm.command.tasks.HelmInstall
import org.unbrokendome.gradle.plugins.helm.command.tasks.HelmUninstall

plugins {
    id("org.unbroken-dome.helm") version "1.3.0"
    id("org.unbroken-dome.helm-releases") version "1.3.0"
    id("com.wiredforcode.spawn") version "0.8.2"
}

buildscript {
    repositories {
        gradlePluginPortal()
    }
}

helm {
    val cloudResourceDashboardChart by charts.creating {
        chartName.set("cloud-resource-dashboard")
        sourceDir.set(file("charts/cloud-resource-dashboard"))
    }

    releases {
        create("cloud-resource-dashboard") {
            from(cloudResourceDashboardChart)
            values.put("web.enabled", project.property("chart.web.enabled").toString().toBoolean())
            values.put("api.enabled", project.findProperty("chart.api.enabled").toString().toBoolean())
            values.put("elasticsearch.enabled", project.findProperty("chart.elasticsearch.enabled").toString().toBoolean())
        }
    }

    downloadClient {
        enabled.set(true)
    }
    repositories {
        helmStable()
        helmIncubator()
        bitnami()
        create("elastic") {
            url("https://helm.elastic.co")
        }
    }
}


val installElasticsearchChart by tasks.creating(HelmInstall::class) {
    chart.value("elastic/elasticsearch")
    releaseName.value("elasticsearch")
    valueFiles.from("charts/es/es-values.yaml")
}

val uninstallElasticsearchChart by tasks.creating(HelmUninstall::class) {
    releaseName.value("elasticsearch")
}

 val installKibanaChart by tasks.creating(HelmInstall::class) {
     chart.value("elastic/kibana")
     releaseName.value("kibana")
     values.put("service.type", "NodePort")
     values.put("resources.requests.cpu", "100m")
     values.put("resources.requests.memory", "100Mi")
     values.put("service.nodePort", "30561")
     values.put("elasticsearchURL", "http://localhost:30920")
 }

 val uninstallKibanaChart by tasks.creating(HelmUninstall::class) {
     releaseName.value("kibana")
 }

val installLogstashChart by tasks.creating(HelmInstall::class) {
    chart.value("elastic/logstash")
    releaseName.value("logstash")
    values.put("service.type", "NodePort")
    values.put("resources.requests.cpu", "100m")
    values.put("resources.requests.memory", "100Mi")
    values.put("service.nodePort", "30960")
}

val uninstallLogstashChart by tasks.creating(HelmUninstall::class) {
    releaseName.value("logstash")
}

val installElkCharts: Task by tasks.creating {
    dependsOn(
//            installLogstashChart,
            installElasticsearchChart
    )
}

val uninstallElkCharts: Task by tasks.creating() {
    dependsOn(uninstallLogstashChart, uninstallElasticsearchChart)
}

val clean: Task by tasks.creating {
    dependsOn(uninstallElkCharts)
}

val runAll by tasks.creating {
    dependsOn(
            installElkCharts,
            "back:api:buildApiImage",
            "web:buildWebImage"
//            "web:runWeb",
//            "back:api:runApi"
    )
}
