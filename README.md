# Next-gen Alpacka

Alpacka/NG is a re-write of my original [Alpacka][alpacka] bash script. It is a single-file deployable written in Go aimed at helping with distro-hopping and unifying package specs where package names differ between distros (looking at you Apache on Debian/Fedora!)

Alpacka provides:

* a single command that works across major distros and derivatives (Debian+, Fedora+, SuSE+, Arch+)
  * distro hoppers' boon!
* a single command for all common operations (no `apt-cache`/`apt-get` differentiation)
* intuitive flags and corresponding short-flags (looking at you `pacman`) for common operations
* no runtime dependencies (limitation of original alpacka)
* manifest spec to help achieve same install when same packages differ in names across distros
  * e.g. python=python2 on old ubuntu and python=python3 on new ubuntu
  * e.g. Apache web server is `apache2` on Debian derivatives and `httpd` on Fedora/Red Hat derivatives
* simple warning system - block actions by setting a warning, bypass the warning with `--ignore-warnings`
  * prevent butter-fingers from operating on shared systems, allow automation to proceed

[alpacka]: https://gitlab.com/taikedz/alpacka

## Download + Install

Download the appropriate binary:

```sh
# Use the version for you
version=0.0.3

# set to "paf-alpine" for alpine deployments
variant=paf

# On Ubuntu/Fedora/etc ... various GNU/Linux:
wget https://github.com/taikedz/alpacka-ng/releases/download/v${version}/${variant} -O /tmp/paf-dl

# Install to a location in root's `PATH` list - for example `/usr/local/bin`

chmod 755 /tmp/paf-dl
sudo mv -i /tmp/paf-dl /usr/local/bin/paf
```

## Build

You need the go build tools on your system. Please see <https://go.dev/doc/install>

Then run

```sh
./build.sh
```

A new build will be produced in the `./bin/` output directory.

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
* Zypper - OpenSUSE
* apk - Alpine (note: must be compiled in an Alpine environment to actually support Alpine ...)

May support in future:

* chocolatey
* winget
* homebrew

Notably, `chocolatey` and `homebrew` should supercede the native package manager : by default these superceders are used, except if `-N` (`native`) is used.

## Packages file format

Alpacka defines a YAML file format that allows checking the contents of the `/etc/os-release` file.

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

Comparisons supported:

* `>=` - greater than or equal to specified value
* `<=` - less than or equal to specified value
* `==` - exactly equal to specified value
* `=~` - release file contains specified value

## Get warnings

You can set a warning for any action. Warnings are messages that are displayed instead of carrying out an option. If a warning is set, a message is printed and `paf` exits without carrying out the activity. If no warning is set, or if prior warning was unset, action proceeds as normal.

To run an action bypassing the warning (execute anyway), use the long bypass option:

  paf -g --ignore-warnings

To set a warning:
    
    paf -w -A upgrade -W "Be careful when upgrading this server - it restarts the core service, which takes a while !"

To unset a warning:

    sudo paf -w -A upgrade -W .

To simply view an existing warning:

    paf -w -A upgrade

Warnings can be set for `upgrade` and `remove`

# License

(C) 2005 Tai Kedzierski

Provided under the terms of the GNU Lesser General Public License v3.0

See [LICENSE.txt](./LICENSE.txt) or [the LGPL online](https://www.gnu.org/licenses/lgpl-3.0.en.html)
