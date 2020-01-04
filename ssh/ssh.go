package ssh

import (
	"errors"
	"github.com/ui-kreinhard/go-setuper/files"
	"github.com/ui-kreinhard/go-setuper/setuper"
	"github.com/ui-kreinhard/go-setuper/utils"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkIfPubkeyFileExists(user string) bool {
	if _, err := os.Stat(getPubkeyFilename(user)); os.IsNotExist(err) {
		return false
	}
	return true
}

func checkIfPubkeyExists(newPubkey Pubkey) (bool, error) {
	pubkeys, err := readAndParsePubKeyFile(newPubkey.User)
	if err != nil {
		return false, err
	}
	for _, pubkey := range pubkeys {
		if pubkey.Pubkey == newPubkey.Pubkey {
			return true, nil
		}
	}
	return false, nil
}

func getSSHDirectory(username string) string {
	if username == "root" {
		return "/" + username + "/.ssh/"
	}
	return "/home/" + username + "/.ssh/"
}

func getPubkeyFilename(username string) string {
	return getSSHDirectory(username) + "authorized_keys"
}

func parsePubKeyFile(rawBytes []byte, user string) ([]Pubkey, error){
	pubkeys := []Pubkey{}

	pubkeyContentRaw := string(rawBytes)
	rawLines := strings.Split(pubkeyContentRaw, "\n")
	for _, rawLine := range rawLines {
		rawPubkeyData := strings.Fields(rawLine)
		if len(rawPubkeyData) <=0 {
			continue
		}
		if len(rawPubkeyData) < 3 {
			return pubkeys, errors.New("Line length doesn't match format. It's " + strconv.FormatInt(int64(len(rawPubkeyData)), 10))
		}
		algorithm := rawPubkeyData[0]
		pubkey := rawPubkeyData[1]
		name := rawPubkeyData[2]
		pubkeys = append(pubkeys, Pubkey{user, pubkey, algorithm, name})
	}
	return pubkeys, nil
}

func readAndParsePubKeyFile(user string) ([]Pubkey, error) {
	pubkeys := []Pubkey{}

	pubkeyContentRawByte, err := ioutil.ReadFile(getPubkeyFilename(user))
	if err != nil {
		return pubkeys, err
	}
	return parsePubKeyFile(pubkeyContentRawByte, user)
}

type Pubkey struct {
	User      string
	Pubkey    string
	Algorithm string
	Name      string
}

func createSSHAndPubkeyFile(user string) error {
	getSSHDirectory(user)
	err := os.MkdirAll(getSSHDirectory(user), 0700)
	if err != nil {
		return err
	}
	err = files.CreateEmptyFileDirect(getPubkeyFilename(user))
	if err != nil {
		return err
	}
	err = files.ChownDirect(getSSHDirectory(user), user, user)
	if err != nil {
		return err
	}
	err = files.ChownDirect(getPubkeyFilename(user), user, user)
	if err != nil {
		return err
	}
	err = files.ChmodDirect(getPubkeyFilename(user), 0700)
	if err != nil {
		return err
	}
	err = files.ChmodDirect(getSSHDirectory(user), 0700)
	return err
	
}

func AddSinglePubkey(pubkey Pubkey) error {
	if !checkIfPubkeyFileExists(pubkey.User) {
		err := createSSHAndPubkeyFile(pubkey.User)	
		if err != nil {
			return err
		}
	}
	pubkeyExists, err := checkIfPubkeyExists(pubkey)
	if err != nil {
		return err
	}
	if !pubkeyExists {
		return utils.AppendToFile(getPubkeyFilename(pubkey.User), ""+pubkey.Algorithm+" "+pubkey.Pubkey+" "+pubkey.Name)
	}
	return nil
}

func AddPubkeysDirect(pubkeys ...Pubkey) error {
	for _, pubkey := range pubkeys {
		err := AddSinglePubkey(pubkey)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddPubkeysFromFileDirect(file, user string) error {
	setuperInstance := setuper.NewSetuper()
	sourceContent, err := setuperInstance.FilesBox.FindString(file)
	if err != nil {
		return err
	}
	pubkeys, err := parsePubKeyFile([]byte(sourceContent), user)
	for _, pubkey := range pubkeys{
		err = AddSinglePubkey(pubkey)
		if err != nil {
			return err
		}
	}
	return nil
}
