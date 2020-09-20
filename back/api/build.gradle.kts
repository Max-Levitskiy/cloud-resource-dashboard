class BuildGolan : Plugin<Project> {
//
//    private val dir: File =
//
//    @javax.inject.Inject
//    constructor()

    override fun apply(project: Project) {
        project.tasks.register("buildHandlers") {
            project.exec {
                commandLine("bash", "-c", "GOOS=linux go build -o bin/api api.go")
            }
//            val dir = project.file("handlers/")
//            dir.listFiles()?.forEach { file ->
//                val name = file.name
//                project.logger.info("Compiling file ${name}...")
//                project.exec {
//                    commandLine("bash", "-c", "GOOS=linux go build -o bin/${name}.go handlers/${name}/main.go")
//                }
//            }
        }
    }

}

apply<BuildGolan>()

val build: Task by tasks.creating {
    dependsOn("buildHandlers")
}
