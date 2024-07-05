package deployment

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"mytexas42-compose/system"
)

func updateGit(environment string) error {
	println(fmt.Sprintf("Updating %s from latest version.", environment))

	backendPath, frontendPath := system.GetCodePaths(environment)

	// Function to handle git pull and possible dubious ownership error
	handleGitPull := func(path string) error {
		//return system.Run("git", "-C", path, "pull")
		repo, err := git.PlainOpen(path)
		if err != nil {
			return err
		}

		workDir, err := repo.Worktree()
		if err != nil {
			return err
		}

		// Load private SSH key
		sshKey, err := ssh.NewPublicKeysFromFile("git", system.GetSSHKeyPath(), system.GetSSHPassphrase())
		if err != nil {
			return fmt.Errorf("failed to load SSH key: %w", err)
		}

		// Configure PullOptions with SSH authentication
		err = workDir.Pull(&git.PullOptions{
			RemoteName: "origin",
			Auth:       sshKey,
		})

		return nil
	}

	// Update backend repository
	if err := handleGitPull(backendPath); err != nil {
		return err
	}

	// Update frontend repository
	if err := handleGitPull(frontendPath); err != nil {
		return err
	}

	return nil
}
