#!/bin/bash

echo ">>>> Deleting old apps"
if test -f "../user-management"; then
    rm ../user-management
fi