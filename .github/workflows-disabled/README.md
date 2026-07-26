# Disabled workflows

`deploy-docker.yml` was removed from the active workflow set. It released the
`multiversx/events-notifier` image to Docker Hub, which is outside the Xorewa
NewArc scope. This repository must not publish, deploy, or otherwise mutate
MultiversX infrastructure.

The active workflows are restricted to pushes and pull requests targeting the
`NewArc` branch and only build, test, or scan this repository.
