post_install() {
    systemctl restart zabbix-agent
}

post_upgrade() {
    systemctl restart zabbix-agent
}

post_remove() {
    systemctl restart zabbix-agent
}
