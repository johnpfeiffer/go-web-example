#!/bin/bash
# Warning: running this will delete all data from the database

# TODO: accept optional parameter to filter scope and leverage the test -run feature

TEST_INTEGRATION=true go test -v

