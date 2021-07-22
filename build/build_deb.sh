#!/bin/bash

BASE_DIR=monitor-client_0.0.1-1_arm64
BIN_DIR=${BASE_DIR}/usr/local/bin
mdkir -p $BIN_DIR

cp $BIN_OUTPUT $BASE_DIR
mv build/DEBIAN $BASE_DIR
dpkg-deb --build --root-owner-group $BASE_DIR
