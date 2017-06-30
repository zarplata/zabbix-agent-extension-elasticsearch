# Maintainer: Andrey Kitsul <a.kitsul@office.ngs.ru>

pkgname=zabbix-agent-extension-elasticsearch
pkgver=20170630.9_98df361
pkgrel=1
pkgdesc="Extension for zabbix-agentd for monitoring Elasticsearch service"
arch=('any')
license=('GPL')
depends=('zabbix-agent')
install="install.sh"
source=("git+ssh://git@git.rn:7999/devops/${pkgname}.git#branch=dev")
md5sums=(
    'SKIP'
    )

make_zabbix_config() {
    userparam_string_discovery="UserParameter=elasticsearch.discovery[*], /usr/bin/${pkgname} --discovery --agg-group \$1"
    userparam_string="UserParameter=elasticsearch.stats[*], /usr/bin/${pkgname} --zabbix-host \$1 --zabbix-prefix \$2"

    echo "$userparam_string_discovery" > "$pkgname.conf"
    echo "$userparam_string" >> "$pkgname.conf"
}

make_vendor() {
    git submodule update --init
    ln -s "${srcdir}"/"${pkgname}"/vendor/ "${srcdir}"/"${pkgname}"/vendor/src
}

pkgver() {
	cd "$pkgname"
	local date=$(git log -1 --format="%cd" --date=short | sed s/-//g)
	local count=$(git rev-list --count HEAD)
	local commit=$(git rev-parse --short HEAD)
    echo "$date.${count}_$commit"
}

build() {
    make_zabbix_config

    cd "${srcdir}"/"${pkgname}"

    make_vendor
    
    GOPATH="${srcdir}"/"${pkgname}"/vendor go build -ldflags "-X main.version=${pkgver}" -o out/"${pkgname}"
}

package() {
	cd "${srcdir}/"${pkgname}
    ZBX_INC_DIR=/etc/zabbix/zabbix_agentd.conf.d/

    install -Dm 0755 out/"${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
    install -Dm 0644 ../"${pkgname}.conf" "${pkgdir}${ZBX_INC_DIR}${pkgname}.conf"
    
}
