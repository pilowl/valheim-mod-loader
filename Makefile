pack_mods:
	@go run cmd/packer/packer.go

run_installer:
	@go run cmd/installer/installer.go

build_installer:
	@go build -o bin/vAlGayM_LC_mOdDEr_v1.exe -ldflags="-w -s" -gcflags=all=-l -ldflags -H=windowsgui cmd/installer/installer.go