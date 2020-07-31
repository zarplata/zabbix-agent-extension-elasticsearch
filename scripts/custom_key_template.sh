#!/bin/bash

if [[ -z $1 ]]; then
    echo "usage: ${0##*/} prefix"
    exit 1
fi

sed "s/elasticsearch\./$1\.elasticsearch\./g" <template_elasticsearch_service.xml \
    | sed "s/None_pfx/$1/g" >elasticsearch_service.xml

