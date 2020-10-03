import com.bmuschko.gradle.docker.tasks.image.DockerBuildImage
import com.moowork.gradle.node.npm.NpmTask

plugins {
  id("com.moowork.node") version "1.3.1"
  id("com.bmuschko.docker-remote-api") version "6.6.1"
}

node {}

val killWebServer by tasks.creating(Exec::class) {
  commandLine("sh", "-c", "lsof -i:4200 -t | xargs kill")
  setIgnoreExitValue(true)
}

val build by tasks.creating(NpmTask::class) {
  setArgs(mutableListOf("install"))
}

val runWeb by tasks.creating(NpmTask::class) {
  dependsOn(killWebServer, build)
  setArgs(mutableListOf("run", "start"))
}

val buildWebImage by tasks.creating(DockerBuildImage::class) {
  inputDir.set(file("./"))
  images.add("cloudresourcedashboard/web:$version")
}
