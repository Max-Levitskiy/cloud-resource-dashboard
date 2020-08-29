import com.moowork.gradle.node.npm.NpmTask

plugins {
    id("com.moowork.node") version "1.3.1"
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
