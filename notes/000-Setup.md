# Setup

## Build devcontainer

I want to use VSCode devcontainers and keep the golang package there. It has taken a bit of work, but I've assembled a custom setup based on

* Understand what [this `devcontainer.json`](https://github.com/devcontainers/images/blob/main/src/go/.devcontainer/devcontainer.json) is doing, roughly, based on the things it's adding (non-root user, etc.)
* Add a non-root user to the container based on [VSCode docs](https://code.visualstudio.com/remote/advancedcontainers/add-nonroot-user)
* Learn how to change a user's default shell [using `chsh`](<https://www.baeldung.com/linux/change-default-shell>)
* Build `./.devcontainer/Dockerfile` to define the base container
* A lot of work figuring out how to put all those pieces together to install `nvm` and `node`, which I'll probably end up needing for something

* Build `./.devcontainer/docker-compose.yml` to run Postgres and adminer with the dev container
  * [Example from Microsoft](https://github.com/microsoft/vscode-dev-containers/blob/main/containers/go-postgres/.devcontainer/docker-compose.yml)
  * [Postgres image docs](https://hub.docker.com/_/postgres)
  * Postgres is exposed on port 9432; adminer is exposed on port 9080
* Build `./.devcontainer/devcontainer.json` based on [Microsoft's example](https://github.com/microsoft/vscode-dev-containers/blob/main/containers/go-postgres/.devcontainer/devcontainer.json)
* Confirm that `go` is installed in the container
  * Running `go` outside the container says it isn't installed
  * Running `go` inside the container shows help, so it is
* Confirm that adminer is exposed on `localhost:9080` outside the container and I can login to the database
* Confirm that Postgres data is in the workspace `db` directory outside the container
  * `sudo chmod g+rwx db` to be able to `ls` the directory without `sudo`
* Confirm that git configuration is carried over in the container with `git config -l`

In the devcontainer, the non-root default user (no password) is `dev` and `dev` is in sudoers so can sudo.

So, now my base environment should be set up. I'm writing this in the container, but it's immediately reflected in the hosting folder.

## Note on compatibility

**TL;DR:** Using VS Code and Docker snap packages seemed to cause issues, including failure to update to newer versions. Removing snaps and installing from apt repos solved the problems.

I built the containers and repo on an Ubuntu 22.04 machine (bare metal). Later, I pulled the repo into an Ubuntu 22.04 VM on a Windows 10 machine (Virtual Box). VS Code tried to start the dev container and failed.

* I could start `docker-compose.yml` with `docker compose` from the command line or VS Code terminal.
* Running `docker ps` showed the containers up.
* The Docker extension didn't complain about connecting to Docker, but did not see containers or images that `docker ps` and `docker image ls` showed.

After reinstalling extensions, change various settings, etc., I was frustrated. So I did the following:

* `snap list` -- both Docker and VS Code were installed from snap; noted that Docker was 20.x vs. 24.x (current as of this writing), VS Code version wasn't standard, so couldn't tell
* `snap refresh` did not update Docker of VS Code versions.
* Ran `snap remove code` to remove snap VS Code and ran `snap saved` and `snap forget <snapshot-number>` to remove the snapshot.
* Removed and pruned all Docker containers, images and volumes.
* Ran `snap remove docker` to remove snap Docker and removed the snapshot.
* Added Docker and Microsoft apt repos and installed Docker and VS Code from them.
  * [official Docker guide](https://docs.docker.com/engine/install/ubuntu/)
  * [guide used for VS Code](https://itslinuxfoss.com/how-to-install-visual-studio-code-on-ubuntu-22-04/)
* Started the containers with `docker compose` from the `.devcontainer` directory.
* Started VS Code and installed the Docker extension.
  * Could not connect to Docker (improvement over last time)
  * Checked and saw it kept the old configuration settings that pointed the docker command to `/snap/bin/docker`
  * Changed configuration to `/usr/bin/docker`
  * Containers, images, etc., visible
* Stopped the containers
* Opened the folder in the container
  * The container took a while to set up because it was installing VS Code server but it started eventually
  * After starting, VS Code prompted to install `gopls`, which I allowed.
* Repeated tests from previous section (go, adminer, postgres, git configuration) to verify it's all working as expected.

TODO: Check the bare metal machine for snap vs. apt packages. (But I prefer apt repo packages anyway, so not going back.)

## Commit prefixes

* FEAT: (new feature for the user, not a new feature for build script)
* FIX: (bug fix for the user, not a fix to a build script)
* DOCS: (changes to the documentation)
* STYLE: (formatting, missing semi colons, etc; no production code change)
* REFACTOR: (refactoring production code, eg. renaming a variable)
* TEST: (adding missing tests, refactoring tests; no production code change)
* CHORE: (updating grunt tasks etc; no production code change)

## Abbreviations

* Ctrl -> controller, as the model-view-controller pattern
