# Next-gen Alpacka

Original Alpacka project was a wrapper named 'paf' for apt/yum/pacman and other tools to make some lazy command-line operation easier for installation and search, and was written in Bash with Bash Builder.

Next-gen Alpacka aims to provide similar functionality, along with a requirements file format to use across package managers.

## Example command line uses

Next Gen Alpacka provides the `paka` command.

```sh
# Install some packages
paka -i package1 package2 ...

# Do an indices update, before installing
paka -u -i packages ...

# Do an upgrade, accept changes
paka -u -g -y


# Install from a packages file
paka -p packages.yaml
```

## Packages file format

We define a format that allows checking the contents of the `/etc/-os-release` file.

Depending on what is found there, certain package groups' definitions are loaded. Variants are declared in-order. The first variant to match is applied.

```yaml
alpacka:
    # Packages needed on specific variants
    variants:
    # os-release key lookups and comparisons
    - release: ID_LIKE=debian
      # package group defs to use
      packages: common, debian
      # manager engine to use
      manager: apt

    - release: ID_LIKE=fedora, VERSION_ID>=22
      packages: common, fedora
      manager: dnf

    - release: ID_LIKE=fedora, VERSION_ID<22
      packages: common, fedora
      manager: yum

    - release: ID_LIKE=arch
      packages: common, debian
      manager: pacman

    # Package groups by name.
    package-groups:
      common:
      - php
      - sqlite

      debian:
      - apache2

      fedora:
      - httpd

```

Supported package manager engines need to be implemented before they can be used.

Package manager engines should be implemented to this tool's repo ideally, but can also be loaded from `/usr/local/lib/alpacka/engines/*.py` files.

# License

Provided under the terms of the GNU Lesser General Public License v3.0

See [LICENSE.txt](./LICENSE.txt) or [the LGPL online](https://www.gnu.org/licenses/lgpl-3.0.en.html)
