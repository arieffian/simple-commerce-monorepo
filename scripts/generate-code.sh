#!/usr/bin/env bash
set -e

oapi-codegen --config api/users/types.cfg.yml api/users/main.all.yml
# oapi-codegen --config api/products/types.cfg.yml api/products/main.all.yml
# oapi-codegen --config api/transactions/types.cfg.yml api/transactions/main.all.yml
