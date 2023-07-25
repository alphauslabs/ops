`ops` is the command line client for the **Long Operations** service [(opsd)](https://github.com/mobingilabs/ouchan/tree/master/cloudrun/opsd).

To install using [HomeBrew](https://brew.sh/), run the following command:

```bash
$ brew install alphauslabs/tap/ops
```

To setup authentication, set your `GOOGLE_APPLICATION_CREDENTIALS` env variable using your credentials file. You also need to give your credentials file access to the `opsd-[next|prod]` service. To do so, try the following commands:

```bash
# Install the `iam` tool:
$ brew install alphauslabs/tap/iam

# Validate `iam` credentials:
$ iam whoami

# Request access to our `tucpd-[next|prod]` service (once only):
$ iam allow-me opsd-prod
```

Explore more available subcommands and flags though:

```bash
$ ops -h
# or
$ ops <subcmd> -h
```
