#!/usr/bin/env bash
read -p "Enter latest version: " version
read -p "confirm ${version} ? (Y/N): " confirm && [[ ${confirm} == [yY] || ${confirm} == [yY][eE][sS] ]] || exit 1
echo "latest version -> ${version}"

echo "Yui -> changing source code"
versionFile="common/version.go"
sed -i '' 's/".*"/"'${version}'"/' ${versionFile}

echo "Yui -> commit source code"
git add ${versionFile}
git commit -m"[yui -> tag: ${version}]"

echo "Yui -> create tag and push it"
git tag ${version}
git push origin ${version}

echo "Yui -> tag success!"
