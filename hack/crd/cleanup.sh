#!/usr/bin/env bash

set -e

SCRIPT_DIR=$(dirname "$0")
pushd "$SCRIPT_DIR"

TEMPLATE_DIR="../../config/crd/base"
mv ${TEMPLATE_DIR}/*componentdefinitions.core.oam.dev.yaml ${TEMPLATE_DIR}/core.oam.dev_componentdefinitions.yaml
mv ${TEMPLATE_DIR}/*applications.core.oam.dev.yaml ${TEMPLATE_DIR}/core.oam.dev_applications.yaml

echo "clean up unused fields of CRDs"

for filename in `ls "$TEMPLATE_DIR"`; do

  sed -i.bak '/creationTimestamp: null/d' "${TEMPLATE_DIR}/$filename"

done

rm ${TEMPLATE_DIR}/*.bak

popd
