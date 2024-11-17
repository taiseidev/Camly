fastlane documentation
----

# Installation

Make sure you have the latest version of the Xcode command line tools installed:

```sh
xcode-select --install
```

For _fastlane_ installation instructions, see [Installing _fastlane_](https://docs.fastlane.tools/#installing-fastlane)

# Available Actions

## iOS

### ios dev_fetch_cert_and_profile

```sh
[bundle exec] fastlane ios dev_fetch_cert_and_profile
```

Dev: Fetch Certificate and Profile

### ios stg_fetch_adhoc_cert_and_profile

```sh
[bundle exec] fastlane ios stg_fetch_adhoc_cert_and_profile
```

Stg: Fetch Ad Hoc Certificate and Profile for broader testing

### ios prod_fetch_cert_and_profile

```sh
[bundle exec] fastlane ios prod_fetch_cert_and_profile
```

Prod: Fetch Certificate and Profile

### ios readonly_dev_fetch_cert_and_profile

```sh
[bundle exec] fastlane ios readonly_dev_fetch_cert_and_profile
```

Readonly: Fetch Certificate and Profiles for Dev and Stg

### ios readonly_stg_fetch_adhoc_cert_and_profile

```sh
[bundle exec] fastlane ios readonly_stg_fetch_adhoc_cert_and_profile
```

Readonly: Fetch Ad Hoc Certificate and Profile for Stg

### ios readonly_prod_fetch_cert_and_profile

```sh
[bundle exec] fastlane ios readonly_prod_fetch_cert_and_profile
```

Readonly: Fetch Certificate and Profile for Prod

----

This README.md is auto-generated and will be re-generated every time [_fastlane_](https://fastlane.tools) is run.

More information about _fastlane_ can be found on [fastlane.tools](https://fastlane.tools).

The documentation of _fastlane_ can be found on [docs.fastlane.tools](https://docs.fastlane.tools).
