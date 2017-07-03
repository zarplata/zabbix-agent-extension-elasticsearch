#!/bin/sh

prefix="$1"

if [ -z "$prefix" ]; then 
    echo "Not define prefix."
    exit 1
fi

sed "s/elasticsearch./$prefix.elasticsearch./g" -i template_elasticsearch_service.xml
sed "s/None_pfx/$prefix/g" -i template_elasticsearch_service.xml

echo "Done."

exit 0
