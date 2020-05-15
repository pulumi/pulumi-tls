#!/bin/bash
# publish-plugin.sh builds and publishes a package containing the resource provider to
# s3://rel.pulumi.com/releases/plugins.
set -o nounset -o errexit -o pipefail

ROOT=$(dirname $0)/..
WORK_PATH=$(mktemp -d)
VERSION=$(jq -r '.version' < "${ROOT}/sdk/nodejs/bin/package.json")
PLUGIN_PACKAGE_NAME="pulumi-resource-tls-${VERSION}-$(go env GOOS)-$(go env GOARCH).tar.gz"
PLUGIN_PACKAGE_DIR="$(mktemp -d)"
PLUGIN_PACKAGE_PATH="${PLUGIN_PACKAGE_DIR}/${PLUGIN_PACKAGE_NAME}"

# When crossbuilding, we want to ensure we have .exe for the windows binaries.
BIN_SUFFIX=
if [ "$(go env GOOS)" = "windows" ]; then
    BIN_SUFFIX=".exe"
fi

(cd "${ROOT}/provider" && go build \
   -ldflags "-X github.com/pulumi/pulumi-tls/provider/v2/pkg/version.Version=${VERSION}" \
   -o "${WORK_PATH}/pulumi-resource-tls${BIN_SUFFIX}" \
   "${ROOT}/cmd/pulumi-resource-tls")

# Tar up the plugin
tar -czf ${PLUGIN_PACKAGE_PATH} -C ${WORK_PATH} .

# Store the initial AWS credentials. The AWS creds will get overwritten after
# calling assume_iam_role, and restored via restore_initial_aws_keys.
export INITIAL_AWS_ACCESS_KEY_ID="${AWS_ACCESS_KEY_ID}"
export INITIAL_AWS_SECRET_ACCESS_KEY="${AWS_SECRET_ACCESS_KEY}"

restore_initial_aws_keys() {
    unset {AWS_ACCESS_KEY_ID,AWS_SECRET_ACCESS_KEY,AWS_SECURITY_TOKEN}
    export AWS_ACCESS_KEY_ID="${INITIAL_AWS_ACCESS_KEY_ID}"
    export AWS_SECRET_ACCESS_KEY="${INITIAL_AWS_SECRET_ACCESS_KEY}"
}

assume_iam_role() {
    local CREDS_JSON=$(aws sts assume-role \
                 --role-arn "${1}" \
                 --role-session-name "upload-plugin-pulumi-resource-tls" \
                 --external-id "upload-pulumi-release")
    export AWS_ACCESS_KEY_ID=$(echo ${CREDS_JSON}     | jq ".Credentials.AccessKeyId" --raw-output)
    export AWS_SECRET_ACCESS_KEY=$(echo ${CREDS_JSON} | jq ".Credentials.SecretAccessKey" --raw-output)
    export AWS_SECURITY_TOKEN=$(echo ${CREDS_JSON}    | jq ".Credentials.SessionToken" --raw-output)
}


echo "Uploading ${PLUGIN_PACKAGE_NAME} to s3://rel.pulumi.com..."

assume_iam_role  "arn:aws:iam::058607598222:role/UploadPulumiReleases"
aws s3 cp --only-show-errors "${PLUGIN_PACKAGE_PATH}" "s3://rel.pulumi.com/releases/plugins/${PLUGIN_PACKAGE_NAME}"
restore_initial_aws_keys

# Assume the role to publish plugins to s3://get.pulumi.com. We upload the plugins to two buckets while
# we transition to only publishing/serving them from get.pulumi.com.
echo "Uploading ${PLUGIN_PACKAGE_NAME} to s3://get.pulumi.com..."

assume_iam_role "arn:aws:iam::058607598222:role/PulumiUploadRelease"
aws s3 cp \
    --only-show-errors --acl public-read \
    "${PLUGIN_PACKAGE_PATH}" "s3://get.pulumi.com/releases/plugins/${PLUGIN_PACKAGE_NAME}"
restore_initial_aws_keys

rm -rf "${PLUGIN_PACKAGE_DIR}"
rm -rf "${WORK_PATH}"
