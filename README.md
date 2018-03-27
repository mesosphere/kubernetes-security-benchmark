# Kubernetes Security Benchmark

This project aims to provide a simple way to evaluate the security of
your Kubernetes deployment against sets of best practices defined by
various community sources.

```shell
$ kubernetes-security-benchmark --help
Run security benchmarks against your Kubernetes clusters.

Usage:
  kubernetes-security-benchmark [command]

Available Commands:
  cis         Run Kubernetes CIS Benchmark tests
  help        Help about any command
  version     Print the version number of Kubernetes Security Benchmark

Flags:
  -h, --help   help for kubernetes-security-benchmark

Use "kubernetes-security-benchmark [command] --help" for more information about a command.
```

## CIS Kubernetes Benchmark

The [Center for Internet Security](https://www.cisecurity.org/) (CIS)
publishes [a benchmark for Kubernetes](https://www.cisecurity.org/benchmark/kubernetes/).
Tests are specified against the various components of a Kubernetes deployment and as such need to be run on the machine (container, VM, or bare-metal) that the component is running on. This project enables a very flexible way to run these tests to match your deployment.

```shell
$ kubernetes-security-benchmark cis --help
Run Kubernetes CIS Benchmark tests.

Usage:
  kubernetes-security-benchmark cis [flags]
  kubernetes-security-benchmark cis [command]

Available Commands:
  api-server                        Run the API server specific benchmarks
  control-plane-configuration-files Run the control plane configuration files specific benchmarks
  controller-manager                Run the controller manager specific benchmarks
  etcd                              Run the etcd specific benchmarks
  federation-api-server             Run the federation API server specific benchmarks
  federation-controller-manager     Run the federation controller manager specific benchmarks
  general-security-primitives       Run the general security primitives specific benchmarks
  kubelet                           Run the kubelet specific benchmarks
  node-configuration-files          Run the node configuration files specific benchmarks
  scheduler                         Run the scheduler specific benchmarks
  version                           Prints the version of the Kubernetes CIS Benchmark

Flags:
  -h, --help                           help for cis
      --spec.dryRun                    If set, ginkgo will walk the test hierarchy without actually running anything.  Best paired with -v.
      --spec.failFast                  If set, ginkgo will stop running a test suite after a failure occurs.
      --spec.failOnMissingProcess      Whether the tests should fail if the relevant process is not running
      --spec.failOnPending             If set, ginkgo will mark the test suite as failed if any specs are pending.
      --spec.flakeAttempts int         Make up to this many attempts to run each spec. Please note that if any of the attempts succeed, the suite will not be failed. But any failures will still be recorded. (default 1)
      --spec.focus string              If set, ginkgo will only run specs that match this regular expression.
      --spec.noColor                   If set, suppress color output in default reporter. (default true)
      --spec.noisyPendings             If set, default reporter will shout about pending tests.
      --spec.noisySkippings            If set, default reporter will shout about skipping tests.
      --spec.progress                  If set, ginkgo will emit progress information as each spec runs to the GinkgoWriter.
      --spec.randomizeAllSpecs         If set, ginkgo will randomize all specs together.  By default, ginkgo only randomizes the top level Describe, Context and When groups.
      --spec.regexScansFilePath        If set, ginkgo regex matching also will look at the file path (code location).
      --spec.seed int                  The seed used to randomize the spec suite. (default 1522082832)
      --spec.skip string               If set, ginkgo will only run specs that do not match this regular expression.
      --spec.skipMeasurements          If set, ginkgo will skip any measurement specs.
      --spec.slowSpecThreshold float   (in seconds) Specs that take longer to run than this threshold are flagged as slow by the default reporter. (default 5)
      --spec.succinct                  If set, default reporter prints out a very succinct report (default true)
      --spec.trace                     If set, default reporter prints out the full stack trace when a failure occurs
      --spec.v                         If set, default reporter print out all specs as they begin.

Use "kubernetes-security-benchmark cis [command] --help" for more information about a command.
```

### Running all tests

In order to run all tests, run:

```shell
$ kubernetes-security-benchmark cis
```

This will run all tests against the machine the binary is run on. This is a very unusual setup because Kubernetes is normally deployed in a distributed fashion, but can be useful for all-in-one deployments such as [Minikube](https://kubernetes.io/docs/getting-started-guides/minikube/).

### Running specific tests

Specific tests can be run via the `--spec.focus` flag. For example, to only run `1.1.1 Ensure that the --anonymous-auth argument is set to false`, you can run:

```shell
$ kubernetes-security-benchmark cis --spec.focus='\[1\.1\.1\]'
```

**Note:** that the `--spec.focus` flag value is a regular expression that matches against the spec description, hence the need to escape the square brackets and dot.

### Running tests targeting a specific component

As a convenience, subcommands are provided to run targeted test suites against specific components, e.g.:

```shell
$ kubernetes-security-benchmark cis api-server
```

This is easier to remember than the equivalent command:

```shell
$ kubernetes-security-benchmark cis --spec.focus='\[1\.1\]'
```
