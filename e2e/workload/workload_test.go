# Binaries for programs and docgen
*.exe
*.exe~
*.dll
*.so
*.dylib
bin
_bin
e2e/vela
vela

# Test binary, build with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
coverage.txt

# Kubernetes Generated files - skip generated files, except for vendored files

!vendor/**/zz_generated.*

# editor and IDE paraphernalia
.idea
*.swp
*.swo
*~
.DS_Store
_.yaml

# Dependency directories (remove the comment below to include it)
vendor/

# Vscode files
.vscode
.history

pkg/test/vela
config/crd/bases
_tmp/

references/cmd/cli/fake/source.go
references/cmd/cli/fake/chart_source.go
references/vela-sdk-gen/*
charts/vela-core/crds/_.yaml
.test_vela
tmp/

.vela/

# check docs
git-page/

vela.json

dist/
