package ssh

func AddPubkeys(pubkeys ...Pubkey) func() (string, error) {
	return func() (string, error) {
		return "", AddPubkeysDirect(pubkeys...)
	}
}

func AddPubkeysFromFile(file, user string) func() (string, error) {
	return func() (string, error) {
		return "", AddPubkeysFromFileDirect(file, user)
	}
}