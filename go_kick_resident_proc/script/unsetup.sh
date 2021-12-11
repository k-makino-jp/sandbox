#!/bin/bash

set -e

function am_i_root_user() {
    USER=$(whoami)
    if [[ $USER != "root" ]]; then
        echo "ERROR: ROOT user must be used. current user = $USER"
        exit 1
    fi
}

function delete_directory() {
    rm -rf "$1"
}

echo "INFO: unsetup started."

am_i_root_user

readonly SECRETS_DIR="/var/run/secrets/kubernetes.io/serviceaccount"
readonly PODINFO_DIR="/etc/podinfo"
delete_directory $SECRETS_DIR
delete_directory $PODINFO_DIR

echo "INFO: unsetup completed."