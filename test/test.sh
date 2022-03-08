#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2022 Alvar Penning
#
# SPDX-License-Identifier: BSD-3-Clause

# This test script executes every Go integration test (tm) within the test
# directory and errors if it does not exit with the expected exit code, defined
# as TEST_ERRORS.

set -u

for f in "$(dirname "$0")"/*.go; do
  export "$(grep "TEST_ERRORS" "$f" | sed -E 's/\/{2}\s*(.*)$/\1/g')"

  go run "$f" > /dev/null 2>&1
  [[ "$?" -eq "$TEST_ERRORS" ]] || { echo "$f failed"; exit 1; }

  unset TEST_ERRORS
done
