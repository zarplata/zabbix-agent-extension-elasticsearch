#!/bin/bash

set -eou pipefail

package_name="zabbix-agent-extension-elasticsearch"
rm -rf *.tar.xz
makepkg -Cod; PKGVER=$(cd $(pwd)/src/$package_name/ && make ver) makepkg -esd

