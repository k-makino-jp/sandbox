#!/bin/bash

set -e

function am_i_root_user() {
    USER=$(whoami)
    if [[ $USER != "root" ]]; then
        echo "ERROR: ROOT user must be used. current user = $USER"
        exit 1
    fi
}

function create_directory() {
    if [ ! -d "$1" ]; then
        mkdir -p "$1"
        chmod 777 "$1"
    fi
}

echo "INFO: setup started."

am_i_root_user

readonly SERVICE_ACCOUNT_DIR="/var/run/secrets/kubernetes.io/serviceaccount"
readonly PODINFO_DIR="/etc/podinfo"
create_directory $SERVICE_ACCOUNT_DIR
create_directory $PODINFO_DIR

echo "INFO: setup completed."