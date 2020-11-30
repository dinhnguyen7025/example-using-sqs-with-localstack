#!/bin/sh

for filename in ./bin/*; do
    [ -e "$filename" ] || continue

    # gz and upload to s3
    zip ${filename}.zip ${filename}
    awslocal s3 cp ${filename}.zip s3://lambda-bucket

    # delete gz file
    rm ${filename}.zip
done
