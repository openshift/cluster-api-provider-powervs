#!/bin/bash

set -e

if [ $# -lt 1 ]; then
    echo "usage: $0 <filename>"
    exit 1
fi

if [ -z "$IBMCLOUD_API_KEY" ]; then
    echo "error: IBMCLOUD_API_KEY is not set in the environment" 2>&1
    exit 1
fi

x=$(echo -n "$IBMCLOUD_API_KEY" | base64)

sed -e "s/IBMCLOUD_API_KEY:.*/IBMCLOUD_API_KEY: $x/" \
    "$1"
