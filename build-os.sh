#!/bin/sh

# compiling the app
package_name="versioncompare"
platforms=("windows/amd64" "linux/amd64" "darwin/amd64")
for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=${package_name}'-'${GOOS}'-'${GOARCH}
	zip_package_name=${output_name}
	if [ ${GOOS} = "windows" ]; then
		output_name+='.exe'
	fi
	env GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${output_name}
	zip -m build/${zip_package_name}.zip ${output_name}
done