import org.unbrokendome.gradle.plugins.helm.command.tasks.HelmInstall
import org.unbrokendome.gradle.plugins.helm.command.tasks.HelmUninstall

plugins {
    id("org.unbroken-dome.helm") version "1.2.0"
}

buildscript {
    repositories {
        gradlePluginPortal()
    }
}

helm {
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
    values.put("persistence.enabled", "false")
    values.put("replicas", "1")
    values.put("service.type", "NodePort")
    values.put("service.nodePort", "30920")
}

val uninstallElasticsearchChart by tasks.creating(HelmUninstall::class) {
    releaseName.value("elasticsearch")
}

task("clean") {
    dependsOn(uninstallElasticsearchChart)
}
