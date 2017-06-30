post_install() {
    systemctl restart zabbix-agentd
}

post_upgrade() {
    systemctl restart zabbix-agentd
}

post_remove() {
    systemctl restart zabbix-agentd
}
