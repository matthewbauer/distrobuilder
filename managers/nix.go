package managers

// NewNix creates a new Manager instance
func NewNix() *Manager {
	return &Manager{
		commands: ManagerCommands{
			clean:   "nix-collect-garbage",
			install: "nix-env",
			refresh: "nix-channel",
			remove:  "nix-env",
			update:  "nix-env",
		},
		flags: ManagerFlags{
			global: []string{},
			clean: []string{
				"-d",
			},
			install: []string{
				"-f", "<nixpkgs>", "-i", "-A",
			},
			remove: []string{
				"-e",
			},
			refresh: []string{
				"--update",
			},
			update: []string{
				"-u",
			},
		},
	}
}
