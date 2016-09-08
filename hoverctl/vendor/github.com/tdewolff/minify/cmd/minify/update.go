package main

import (
	"fmt"

	"github.com/equinox-io/equinox"
)

const appID = "app_bDLQozP627p"

// public portion of signing key generated by `equinox genkey`
var publicKey = []byte(`
-----BEGIN ECDSA PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEHtVEOqnnA3qX5ZZF3p22hqLpfMtgU+wv
eBhQ7XUAulc+Is8qzr8Z63VGvgHQPuTdXyHRnJUQWixv/tRXqlCiIGULqK4W0QFn
v5XxiWrJJCB8n5DGC4e4RPqiVbaknDOi
-----END ECDSA PUBLIC KEY-----
`)

func equinoxUpdate() error {
	var opts equinox.Options
	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
		return err
	}

	// check for the update
	resp, err := equinox.Check(appID, opts)
	switch {
	case err == equinox.NotAvailableErr:
		fmt.Println("No update available, already at the latest version!")
		return nil
	case err != nil:
		fmt.Println("Update failed:", err)
		return err
	}

	// fetch the update and apply it
	err = resp.Apply()
	if err != nil {
		return err
	}

	fmt.Printf("Updated to new version: %s!\n", resp.ReleaseVersion)
	return nil
}