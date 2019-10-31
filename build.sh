#!/bin/sh

CURRENT_VERSION="$(cat version)"
./increment_version.sh -m "$CURRENT_VERSION" > version

echo "Building new version $CURRENT_VERSION"

docker build -t koloooo/mbb:"$(cat version)" .
echo "Docker image build"
echo "Pushing.."

docker push koloooo/mbb:"$(cat version)"
echo "Docker image pushed"
