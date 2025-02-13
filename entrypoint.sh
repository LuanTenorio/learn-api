#!/bin/sh

echo "Running migrations ..."
cd migration && tern migrate && cd ..

echo "Starting the application ..."
exec air
