# go.mod needs Go 1.23. Go 1.27.3 also fixes CVE-2026-111943 in crypto/x509.
FROM golang:1.27.3 AS build
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# CVE-2026-114406: a distroless build from after the Debian ca-certificates fix that drops
# emSign ECC Root CA - G3 and Comodo AAA
FROM gcr.io/distroless/static:nonroot@sha256:25ac8f0c3b2334a9eb05452266bd4697864a8c8ef0ef71e5e6480d7f92868329
WORKDIR /
COPY --from=build /workspace/manager .
USER nonroot:nonroot

ENTRYPOINT ["/manager"]
