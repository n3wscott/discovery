# Using KinD

The following is what is needed to use KinD to run the e2e tests. Base path is
assumed to be the top level project.

See .github/workflows/kind-e2e.yaml for minimum kind version

Startup a cluster:

```shell
./test/kind/bootstrap.sh
```

Install the Discovery Controllers:

```shell
./test/kind/install.sh
```

Run the end to end tests:

```shell
./test/kind/run-tests.sh
```

Cleanup:

```shell
./test/kind/bootstrap.sh --shutdown
```
