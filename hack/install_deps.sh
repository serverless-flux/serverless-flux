#!/bin/sh -eu

go get github.com/vektra/mockery/.../
go get github.com/mattn/goveralls

# managing all linters that gometalinter uses with dep is going to take
# a lot of work, so we install all of those from the release tarball
install_gometalinter() {
  version="${1}"
  prefix="https://github.com/alecthomas/gometalinter/releases/download"
  if [ "$(uname)" = "Darwin" ] ; then
    suffix="darwin-amd64"
  else
    suffix="linux-amd64"
  fi
  basename="gometalinter-${version}-${suffix}"
  url="${prefix}/v${version}/${basename}.tar.gz"
  cd "${GOPATH}/bin/"
  curl --silent --location "${url}" | tar xz
  (cd "./${basename}/" ; mv ./* ../)
  rmdir "./${basename}"
  unset version prefix suffix basename url
}

install_gometalinter "2.0.11"
