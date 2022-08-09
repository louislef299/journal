# Simple To Do API server

The goal of this project is to create a simple todo server that can persist data in a k8s cluster as well as run some example deployments as I learn some CI/CD automations. Hoping to have a POC to roll out new changes via deloyments and github actions with this project. For reference, this is an example diagram of what I am initially thinking of:

![todo automation](./.images/todo_api.png)

To run a live development session with air, run:
`docker run -it --rm -w $(pwd) -v $(pwd):$(pwd) -dp 8080:8080 cosmtrek/air`

This works because of using a [bind mount](https://docs.docker.com/get-started/06_bind_mounts/) shortens the development time to see changes instantly withouts needing all of the build tools and environments installed. Since the environment is technically your project, and since there are tools that will automatically update for filesystem changes, this makes developing at speed much faster! 