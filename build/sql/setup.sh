#!/bin/bash
set -e
service postgresql start
psql < ./build/sql/create_tables.sql
# service psql stop