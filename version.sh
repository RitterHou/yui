#!/usr/bin/env bash
read -p "Enter latest version: " version
read -p "confirm ${version} ? (Y/N): " confirm && [[ ${confirm} == [yY] || ${confirm} == [yY][eE][sS] ]] || exit 1
echo "latest version -> ${version}"

echo "changing source code"
versionFile="common/version.go"
versionFull="\"${version}\""
sed -i -e "s/\(const Version = \).*/\1$versionFull/" ${versionFile}

echo "commit source code"
git add ${versionFile}
git commit -m"[yui]tag: ${version}"

echo "create tag and push it"
git tag ${version}
git push origin ${version}

echo "tag success!"
