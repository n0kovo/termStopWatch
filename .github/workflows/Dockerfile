# Use build arguments for versions
ARG GO_VERSION=1.21
ARG NDK_VERSION=r26b

# DEBUG
RUN echo "NDK_VERSION=${NDK_VERSION}, GO_VERSION=${GO_VERSION}"

# Use an official Go runtime as a parent image
FROM golang:${GO_VERSION} as builder

# Install Android NDK
RUN apt-get update && apt-get install -y curl unzip
RUN curl -O https://dl.google.com/android/repository/android-ndk-${NDK_VERSION}-linux.zip && \
    unzip android-ndk-${NDK_VERSION}-linux-x86_64.zip && \
    mv android-ndk-${NDK_VERSION} /android-ndk

# Set the environment for the Android NDK
ENV ANDROID_NDK_HOME /android-ndk
ENV PATH $PATH:$ANDROID_NDK_HOME
