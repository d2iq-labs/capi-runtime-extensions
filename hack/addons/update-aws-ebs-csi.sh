#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
readonly SCRIPT_DIR

# shellcheck source=hack/common.sh
source "${SCRIPT_DIR}/../common.sh"

if [ -z "${AWS_EBS_CSI_VERSION:-}" ]; then
  echo "Missing environment variable: AWS_EBS_CSI_VERSION"
  exit 1
fi

readonly ASSETS_DIR="${GIT_REPO_ROOT}/.local/csi/aws-ebs/${AWS_EBS_CSI_VERSION}"
mkdir -p "${ASSETS_DIR}"

FILENAME="aws-ebs-csi.yaml"

kubectl create -k \
  "github.com/kubernetes-sigs/aws-ebs-csi-driver/deploy/kubernetes/overlays/stable/?ref=${AWS_EBS_CSI_VERSION}" \
  --dry-run=client -o yaml |
  sed 's|imagePullPolicy: Always|imagePullPolicy: IfNotPresent|g' \
    >"${ASSETS_DIR}/${FILENAME}"

# Add default storage class here.
cat <<EOF >"${ASSETS_DIR}/aws-ebs-csi-temp.yaml"
---
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: ebs-sc
provisioner: ebs.csi.aws.com
volumeBindingMode: WaitForFirstConsumer
parameters:
  csi.storage.k8s.io/fstype: ext4
  type: gp3
---
$(cat "${ASSETS_DIR}/${FILENAME}")
EOF

mv "${ASSETS_DIR}/aws-ebs-csi-temp.yaml" "${ASSETS_DIR}/${FILENAME}"

# add CSI external-snapshotter CRDs and Deployment
FILES=(
  snapshot.storage.k8s.io_volumesnapshotclasses.yaml
  snapshot.storage.k8s.io_volumesnapshotcontents.yaml
  snapshot.storage.k8s.io_volumesnapshots.yaml
)

for file in "${FILES[@]}"; do
  echo "---" >>"${ASSETS_DIR}/${FILENAME}"

  kubectl create -f \
    "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${AWS_CSI_SNAPSHOT_CONTROLLER_VERSION}/client/config/crd/$file" \
    --dry-run=client -o yaml >>"${ASSETS_DIR}/${FILENAME}"
done

#shellcheck disable=SC2129 # want to add breaks in a specific order
echo "---" >>"${ASSETS_DIR}/${FILENAME}"
kubectl create -f \
  "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${AWS_CSI_SNAPSHOT_CONTROLLER_VERSION}/deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml" \
  --dry-run=client -o yaml >>"${ASSETS_DIR}/${FILENAME}"

echo "---" >>"${ASSETS_DIR}/${FILENAME}"
kubectl create -f \
  "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${AWS_CSI_SNAPSHOT_CONTROLLER_VERSION}/deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml" \
  --dry-run=client -o yaml >>"${ASSETS_DIR}/${FILENAME}"
cat <<EOF >>"${ASSETS_DIR}/${FILENAME}"
      nodeSelector:
        kubernetes.io/os: linux
        node-role.kubernetes.io/control-plane: ""
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
        tolerationSeconds: 300
      - key: node-role.kubernetes.io/master
        effect: NoSchedule
        operator: Exists
EOF

readonly KUSTOMIZE_BASE_DIR="${SCRIPT_DIR}/kustomize/aws-ebs/"
mv "${ASSETS_DIR}/${FILENAME}" "${KUSTOMIZE_BASE_DIR}"/"${FILENAME}"
kustomize --load-restrictor=LoadRestrictionsNone build --reorder=none "${KUSTOMIZE_BASE_DIR}" >>"${ASSETS_DIR}/${FILENAME}"
rm "${KUSTOMIZE_BASE_DIR}"/"${FILENAME}"

readonly OUTPUT_FILE=aws-ebs-configmap

kubectl create configmap aws-ebs-csi-crs-cm --dry-run=client --output yaml \
  --from-file "${ASSETS_DIR}/${FILENAME}" \
  >"${GIT_REPO_ROOT}/charts/capi-runtime-extensions/templates/csi/aws-ebs/${OUTPUT_FILE}.yaml"

cat <<EOF >"${GIT_REPO_ROOT}/charts/capi-runtime-extensions/templates/csi/aws-ebs/${OUTPUT_FILE}-temp.yaml"
#=================================================================
#                 DO NOT EDIT THIS FILE
#  IT HAS BEEN GENERATED BY /hack/addons/update-aws-ebs-csi.sh
#=================================================================
$(cat "${GIT_REPO_ROOT}/charts/capi-runtime-extensions/templates/csi/aws-ebs/${OUTPUT_FILE}.yaml")
EOF

mv "${GIT_REPO_ROOT}/charts/capi-runtime-extensions/templates/csi/aws-ebs/${OUTPUT_FILE}-temp.yaml" "${GIT_REPO_ROOT}/charts/capi-runtime-extensions/templates/csi/aws-ebs/${OUTPUT_FILE}.yaml"
