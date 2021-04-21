#!/bin/bash

BASEDIR=$(dirname $0)
config=${BASEDIR}/config.yaml

usage() { 
  echo "Usage assumes setup for Apigee X. You may need to edit this for hybrid and definitely for legacy/opdk" 1>&2
  echo "Usage: $0 [-h] [-a <path-to-ax-sa-json>] [-u <user-email: an email address for the account>] [-o <organiztion> ] [-e <environment>] [-t <token>] [-r <remote url>] [-c <config-file-path]" 1>&2
  exit 1;
}

while getopts "a::e:o:t:r:c:u:h" o; do
    case "${o}" in
        h)
          usage
        ;;
        a)
            analytics=${OPTARG}
        ;;
        u)
            email=${OPTARG}
        ;;
        e)
            env=${OPTARG}
        ;;
        o)
            org=${OPTARG}
        ;;
        t)
            token=${OPTARG}
        ;;
        r)
            remoteurl=${OPTARG}
        ;;
        c)
            config=${OPTARG}
        ;;
        *)
            usage
        ;;
    esac
done
shift $((OPTIND-1))

if [ -z "${env}" ] || [ -z "${org}" ] || [ -t "${token}" ] || [ -z "${remoteurl}" ] || [ -z "${email}" ]; then
    usage
fi

BASEURL=https://apigee.googleapis.com/v1/organizations/${org}

OS=$(uname -s)

case $OS in
  "Linux")
    arch="linux"
  ;;
  "Darwin")
    arch="macOS"
  ;;
esac

wget "https://github.com/apigee/apigee-remote-service-cli/releases/download/v2.0.0/apigee-remote-service-cli_2.0.0_${arch}_64-bit.tar.gz" -O apigee-remote-service-cli.tar.gz
wget "https://github.com/apigee/apigee-remote-service-envoy/releases/download/v2.0.0/apigee-remote-service-envoy_2.0.0_${arch}_64-bit.tar.gz" -O apigee-remote-service-envoy.tar.gz



tar xf apigee-remote-service-cli.tar.gz apigee-remote-service-cli
tar xf apigee-remote-service-envoy.tar.gz apigee-remote-service-envoy
cp apigee-remote-service-envoy envoy_adapter


${BASEDIR}/apigee-remote-service-cli provision -o $org -e $env -t $token --runtime $remoteurl --analytics-sa ${analytics} > $config

# Our api product
curl -X POST "${BASEURL}/apiproducts" -H "Authorization: Bearer $token" -H "Content-Type: application/json" -d @${BASEDIR}/apigee-jsons/apiproduct.json > ${BASEDIR}/my_apiproduct.json

# Our developer
user=$(echo gsjurseth@google.com | sed -e 's/\(.*\)@.*/\1/')
cat $BASEDIR/apigee-jsons/developer.json | sed -e "s/@@EMAIL@@/${email}/" -e "s/@@USER@@/${user}/" | curl -X POST "${BASEURL}/developers" -H "Authorization: Bearer $token" -H "Content-Type: application/json" -d @- > ${BASEDIR}/my_developer.json

# Our app
curl -X POST "${BASEURL}/developers/${email}/apps" -H "Authorization: Bearer $token" -H "Content-Type: application/json" -d @${BASEDIR}/apigee-jsons/app.json > ${BASEDIR}/my_app.json


rm apigee-remote-service-*
