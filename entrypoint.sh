#!/bin/bash

echo "Running unit tests"
make unit-test

echo "Running integration tests against ${ACCOUNTS_API_BASE_URL}"
make integration-test
