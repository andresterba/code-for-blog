test:
	pack build buildpacks-test-pack-tiny --path . --buildpack paketo-buildpacks/go  --builder paketobuildpacks/builder:tiny --creation-time now
	pack build buildpacks-test-pack-go --path . --buildpack paketo-buildpacks/go --builder paketobuildpacks/builder:base --creation-time now
	docker build -t buildpacks-test-docker -f Dockerfile .

size: test
	docker image ls | rg buildpacks-test

dive: test
	CI=true dive buildpacks-test-pack-go
	CI=true dive buildpacks-test-pack-tiny
	CI=true dive buildpacks-test-docker

sbom: test
	pack sbom download buildpacks-test-pack-tiny
	tree layers/