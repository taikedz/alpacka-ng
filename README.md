# Next-gen Alpacka

[Original Alpacka project](https://gitlab.com/taikedz/alpacka) was a wrapper named 'paf' written in Bash with Bash Builder.

This Next-gen Alpacka aims to provide similar functionality, along with a requirements file format to use across package managers. It is re-written in Go for better portability, and zero-dependency at runtime.

## Why ?

A few reasons:

1. All package managers have their idiosyncracies even with regard to activities they have in common. Alpacka unifies the workflow.
2. Some package managers have dissociated commands, or complicated syntax
  * Alpacka overcomes the former, and provides a uniform syntax to get around the latter.
3. No package managers allow setting pre-action warnings - this feature helps avoid accidentally butter-fingering an upgrade and downtiming a server.
4. I'm lazy and don't like retyping a command simply to change one part. Alpacka only runs the last action specified amongst info, install, remove, or system upgrade...
  * Run `paf -s minetest` to show the package, then use up-arrow and add `-i` to install it (`paf -s minetest -i` , ignores prior `-s` flag)

## Example command line uses

Run the same command on Ubuntu, Fedora, Arch or openSUSE environments:

```sh
# Install some packages
paf -i package1 package2 ...

# Do an indices update, before installing
paf -u -i packages ...
# also
paf -ui packages ...

# Do an upgrade, accept changes
paf -u -g -y
# also
paf -ugy


# Install from a packages file
paf -m -M packages.yaml
```

## Package managers supported

* APT - Debian/Ubuntu family
* dnf/yum - Fedora/Red Hat family
* pacman - Arch family
* apk - Alpine
* Zypper - OpenSUSE


## Packages file format

(TBD)

We define a format that allows checking the contents of the `/etc/os-release` file.

Depending on what is found there, certain package groups' definitions are loaded. Variants are declared in-order. The first variant to match is applied.

```yaml
alpacka:
    # os-release key lookups and comparisons
    variants:
    - release: ID_LIKE=~ubuntu, VERSION_ID>=18.04
      # package group defs to use
      groups:
      - common
      - debian
      - newbuntu

    - release: ID_LIKE=~debian
      groups:
       - common
       - debian

    - release: ID_LIKE=~fedora
      groups:
      - common
      - fedora

    - release: ID_LIKE=arch
      groups:
      - common
      - debian

    # Package groups by name.
    package-groups:
      common:
      - php
      - sqlite

      debian:
      - apache2

      fedora:
      - httpd

      newbuntu:
      - pythonispython2

```

## Get warnings

(TBD)

You can also set a warning for any action. Warnings are messages that are displayed before an action is carried out. If a warning is set, a message is printed and nothing is performed.

To run an action bypassing the warning (execute anyway), use the long bypass option

  paf -g `--live-dangerously`

To set a warning:
    
    paf -w upgrade -W "Be careful when upgrading this server - it restarts the core service, which takes a while !"

To unset a warning:

    sudo paf -w upgrade -W .

To simply view an existing warning:

    paf -w upgrade

Warnings can be set for `upgrade` and `remove`

# License

(C) 2005 Tai Kedzierski

Provided under the terms of the GNU Lesser General Public License v3.0

See [LICENSE.txt](./LICENSE.txt) or [the LGPL online](https://www.gnu.org/licenses/lgpl-3.0.en.html)
