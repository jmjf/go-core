# Setup

## Devcontainer

I want to use VSCode devcontainers and keep the golang package there. It has taken a bit of work, but I've assembled a custom setup based on

* Understand what [this `devcontainer.json`](https://github.com/devcontainers/images/blob/main/src/go/.devcontainer/devcontainer.json) was doing
* Add a non-root user to the container based on [VSCode docs](https://code.visualstudio.com/remote/advancedcontainers/add-nonroot-user)
* Learn how to change a user's default shell using `chsh` [from Baeldung]](<https://www.baeldung.com/linux/change-default-shell>)
* Build a `Dockerfile` to establish the base container
* A lot of work figuring out how to put all those pieces together to install `nvm` and `node`, which I'll probably end up needing for something

* Build a `docker-compose.yml` to run Postgres and adminer with the dev container
  * [Example from Microsoft](https://github.com/microsoft/vscode-dev-containers/blob/main/containers/go-postgres/.devcontainer/docker-compose.yml)
  * [Postgres image docs](https://hub.docker.com/_/postgres)
  * Postgres is exposed on port 9432; adminer is exposed on port 9080
* Build `devcontainer.json` based on [Microsoft's example](https://github.com/microsoft/vscode-dev-containers/blob/main/containers/go-postgres/.devcontainer/devcontainer.json)
* Confirm that `go` is installed in the container
  * Running `go` outside the container says it isn't installed
* Confirm that adminer is exposed on `localhost:9080` outside the container
* Confirm that Postgres data is in the workspace `db` directory outside the container
  * `sudo chmod g+rwx db` to be able to `ls` the directory without `sudo`
* Confirm that git configuration is carried over in the container with `git config -l`

So, now my base environment should be set up. I'm writing this in the container, but it's immediately reflected in the hosting folder.

## Commit prefixes

* FEAT: (new feature for the user, not a new feature for build script)
* FIX: (bug fix for the user, not a fix to a build script)
* DOCS: (changes to the documentation)
* STYLE: (formatting, missing semi colons, etc; no production code change)
* REFACTOR: (refactoring production code, eg. renaming a variable)
* TEST: (adding missing tests, refactoring tests; no production code change)
* CHORE: (updating grunt tasks etc; no production code change)
