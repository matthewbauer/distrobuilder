package sources

import (
	"fmt"
	"path/filepath"

	lxd "github.com/lxc/lxd/shared"

	"github.com/lxc/distrobuilder/shared"
)

type NixOSHTTP struct{}

func NewNixOSHTTP() *NixOSHTTP {
	return &NixOSHTTP{}
}

// Run downloads a NixOS container tarball.
func (s *NixOSHTTP) Run(definition shared.Definition, rootfsDir string) error {
	// https://hydra.nixos.org/job/nixos/release-%s
	tarball := fmt.Sprintf("%s/nixos.containerTarball.%s-linux/latest/download-by-type/file/system-tarball",
		definition.Source.URL, definition.Image.ArchitectureMapped)

	// Download
	fpath, err := shared.DownloadHash(definition.Image, tarball, "", nil)
	if err != nil {
		return err
	}

	// Unpack
	err = lxd.Unpack(filepath.Join(fpath, "system-tarball"), rootfsDir, false, false, nil)
	if err != nil {
		return err
	}

	return nil
}
