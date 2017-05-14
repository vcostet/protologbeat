
### Version 0.2.0
- Added support for receiving GELF format messages by setting `enable_gelf: true`

### Version 0.1.1
- Added Dockerfile and seperate `protologbeat-docker.yml` config file to be used by docker image
- Updated default `protologbeat.yml` to have bare-minimum config values
- Added `build-bin.sh` build script to simplify compiling the binary for the most common platforms