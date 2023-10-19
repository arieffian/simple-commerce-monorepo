#!/usr/bin/env bash
set -e

oapi-codegen --config api/users/types.cfg.yml api/users/main.all.yml
oapi-codegen --config api/products/types.cfg.yml api/products/main.all.yml
# oapi-codegen --config api/orders/types.cfg.yml api/orders/main.all.yml
# oapi-codegen --config api/payments/types.cfg.yml api/payments/main.all.yml
