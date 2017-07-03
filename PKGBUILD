# Maintainer: Andrey Kitsul <a.kitsul@office.ngs.ru>

pkgname=zabbix-agent-extension-elasticsearch
pkgver=20170630.3_52d780d
pkgrel=1
pkgdesc="Extension for zabbix-agentd for monitoring Elasticsearch service"
arch=('any')
license=('GPL')
makedepends=('go')
depends=('zabbix-agent')
install="install.sh"
source=("git+https://github.com/zarplata/$pkgname.git#branch=master")
md5sums=(
    'SKIP'
    )

make_zabbix_config() {
    userparam_string_discovery="UserParameter=elasticsearch.discovery[*], /usr/bin/${pkgname} --discovery --agg-group \$1"
    userparam_string="UserParameter=elasticsearch.stats[*], /usr/bin/${pkgname} --zabbix-host \$1 --zabbix-prefix \$2"

    echo "$userparam_string_discovery" > "$pkgname.conf"
    echo "$userparam_string" >> "$pkgname.conf"
}


pkgver() {
	cd "$pkgname"
    make ver

}

build() {
    make_zabbix_config

    cd "$srcdir/$pkgname"
    make 
}

package() {
	cd "$srcdir/$pkgname"
    ZBX_INC_DIR=/etc/zabbix/zabbix_agentd.conf.d/

    install -Dm 0755 .out/"${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
    install -Dm 0644 ../"${pkgname}.conf" "${pkgdir}${ZBX_INC_DIR}${pkgname}.conf"
    
}
