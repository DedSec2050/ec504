#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Application variables
APP_NAME="my-go-backend"
BUILD_DIR="./build"
OUTPUT_DIR="./output"
BINARY_NAME="$APP_NAME"
SRC_DIR=$(pwd) # Automatically get the current directory

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}Building the Go application from: ${SRC_DIR}${NC}"

# Step 1: Clean up previous builds
if [ -d "$BUILD_DIR" ]; then
    echo "Cleaning up old build directory..."
    rm -rf "$BUILD_DIR"
fi

if [ -d "$OUTPUT_DIR" ]; then
    echo "Cleaning up old output directory..."
    rm -rf "$OUTPUT_DIR"
fi

# Step 2: Create new build directory
mkdir -p "$BUILD_DIR"
mkdir -p "$OUTPUT_DIR"

# Step 3: Set up Go environment for cross-compilation (Linux only)
export GOOS=linux
export GOARCH=amd64

# Step 4: Build the application binary
echo "Building the Go binary..."
go build -o "$BUILD_DIR/$BINARY_NAME" "$SRC_DIR"

echo -e "${GREEN}Build successful! Binary located at $BUILD_DIR/$BINARY_NAME${NC}"

# Step 5: Package the binary and assets
echo "Packaging the application..."
tar -czvf "$OUTPUT_DIR/$APP_NAME.tar.gz" -C "$BUILD_DIR" "$BINARY_NAME" -C "$SRC_DIR" configs static

echo -e "${GREEN}Packaging complete! Tarball located at $OUTPUT_DIR/$APP_NAME.tar.gz${NC}"

# Step 6: Optional - Restart the application (if already running)
if pgrep "$BINARY_NAME" > /dev/null; then
    echo -e "${GREEN}Stopping the existing application...${NC}"
    pkill "$BINARY_NAME"
fi

echo -e "${GREEN}Starting the new version of the application...${NC}"
nohup "$BUILD_DIR/$BINARY_NAME" > "$OUTPUT_DIR/app.log" 2>&1 &

echo -e "${GREEN}Deployment complete!${NC}"
