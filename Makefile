all:
	@gox -os="darwin linux windows" \
	     -arch="arm 386 amd64" \
	     -output="pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"
