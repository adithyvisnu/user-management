#!/bin/bash

echo ">>>> Deleting old apps"
if test -f "../user-service"; then
    rm ../user-service
fi