import com.wiredforcode.gradle.spawn.*
import com.bmuschko.gradle.docker.tasks.image.DockerBuildImage

plugins {
    id("com.bmuschko.docker-remote-api") version "6.6.1"
}

val build by tasks.creating(Exec::class) {
    commandLine("sh","-c", "go build -o bin/api api.go")
}

val runApi by tasks.creating(Exec::class) {
    dependsOn(build)
    commandLine("bin/api")

}

val buildApiImage by tasks.creating(DockerBuildImage::class) {
    inputDir.set(file("./"))
    images.add("cloudresourcedashboard/api:latest")
}
