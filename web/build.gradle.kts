import com.bmuschko.gradle.docker.tasks.image.DockerBuildImage
import com.moowork.gradle.node.npm.NpmTask

plugins {
  id("com.moowork.node") version "1.3.1"
  id("com.bmuschko.docker-remote-api") version "6.6.1"
}

node {
}

val killWebServer by tasks.creating(Exec::class) {
  commandLine("bash", "-c", "lsof -i:4200 -t | xargs kill")
  setIgnoreExitValue(true)
}

val runWeb by tasks.creating(NpmTask::class) {
  dependsOn(killWebServer)
  setArgs(mutableListOf("run", "start"))
}

tasks.create("buildWebImage", DockerBuildImage::class) {
  inputDir.set(file("./"))
  images.add("cloud-resource-dashboard/web:latest")
}
