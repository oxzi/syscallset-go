#!/usr/bin/env bash

# SPDX-FileCopyrightText: Alvar Penning
#
# SPDX-License-Identifier: BSD-3-Clause

# This test script executes every Go integration test (tm) within the
# subdirectories and errors if it does not exit with the expected exit code,
# defined as TEST_ERRORS.

for d in "$(dirname "$0")"/*/; do
  echo -n "${d} "
  go run "${d}/test.go" > /dev/null 2>&1
  [[ "$?" -eq "$(sed -n -E 's/.*TEST_ERRORS=([01])$/\1/p' "${d}/test.go")" ]] \
    && echo "good" || { echo "fail"; exit 1; }
done
