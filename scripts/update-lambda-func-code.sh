#!/bin/sh

for filename in ./bin/*; do
    [ -e "$filename" ] || continue

    # gz and update lambda function code
    zip ${filename}.zip ${filename}
    awslocal lambda update-function-code --function-name=custom-lambda --zip-file fileb://bin/custom.zip

    # delete gz file
    rm ${filename}.zip
done
