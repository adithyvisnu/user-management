#!/bin/bash

echo ">>>> Deleting old apps"
if test -f "../app"; then
    rm ../app
fi