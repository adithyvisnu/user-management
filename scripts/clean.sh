#!/bin/bash

echo ">>>> Deleting old apps"
if test -f "../application"; then
    rm ../application
fi