#!/bin/bash

set -e
set -o pipefail

bin_name="subs"
destination="/usr/local/bin/"

rm -rf $destination/$bin_name

echo "Uninstalled successfully!"