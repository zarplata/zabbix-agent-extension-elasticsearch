# Maintainer: Andrey Kitsul <a.kitsul@zarplata.ru>
# Contributor: Andrey Platonov <a.platonov@zarplata.ru>

pkgname=zabbix-agent-extension-elasticsearch
pkgver=20180723.22_75f1391
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


pkgver() {
	cd "$pkgname"
    make ver

}

build() {
    cd "$srcdir/$pkgname"
    make 
}

package() {
	cd "$srcdir/$pkgname"
    ZBX_INC_DIR=/etc/zabbix/zabbix_agentd.conf.d/

    install -Dm 0755 .out/"${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
    install -Dm 0644 "${pkgname}.conf" "${pkgdir}${ZBX_INC_DIR}${pkgname}.conf"
    
}
