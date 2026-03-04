package uci

import (
	"fmt"
	"strings"
)

// DefaultTreePath points to the default UCI location.
const DefaultTreePath = "/etc/config"

// defaultTree is a convenient accessor to the UCI default location.
var defaultTree = NewTree(DefaultTreePath)

// LoadConfig delegates to the default tree. See Tree for details.
func LoadConfig(name string, forceReload bool) error {
	return defaultTree.LoadConfig(name, forceReload)
}

// Commit delegates to the default tree. See Tree for details.
func Commit(configs ...string) error {
	return defaultTree.Commit(configs...)
}

// Revert delegates to the default tree. See Tree for details.
func Revert(configs ...string) {
	defaultTree.Revert(configs...)
}

// GetSections delegates to the default tree. See Tree for details.
func GetSections(config, secType string) ([]string, error) {
	return defaultTree.GetSections(config, secType)
}

// GetSectionOptions delegates to the default tree. See Tree for details.
func GetSectionOptions(config, section string) ([]SectionOption, error) {
	return defaultTree.GetSectionOptions(config, section)
}

// Get delegates to the default tree. See Tree for details.
func Get(config, section, option string) ([]string, bool) {
	return defaultTree.Get(config, section, option)
}

// GetLast delegates to the default tree. See Tree for details.
func GetLast(config, section, option string) (string, bool) {
	return defaultTree.GetLast(config, section, option)
}

// GetBool delegates to the default tree. See Tree for details.
func GetBool(config, section, option string) (bool, bool) {
	return defaultTree.GetBool(config, section, option)
}

// Del delegates to the default tree. See Tree for details.
func Del(config, section, option string) error {
	return defaultTree.Del(config, section, option)
}

// Set delegates to the default tree. See Tree for details.
func Set(config, section, option string, values ...string) error {
	return defaultTree.SetType(config, section, option, TypeOption, values...)
}

// SetType delegates to the default tree. See Tree for details.
func SetType(config, section, option string, typ OptionType, values ...string) error {
	return defaultTree.SetType(config, section, option, typ, values...)
}

// AddSection delegates to the default tree. See Tree for details.
func AddSection(config, section, typ string) error {
	return defaultTree.AddSection(config, section, typ)
}

// DelSection delegates to the default tree. See Tree for details.
func DelSection(config, section string) error {
	return defaultTree.DelSection(config, section)
}

// AddAnonymousSection delegates to the default tree. See Tree for details.
func AddAnonymousSection(config, typ string) (string, error) {
	return defaultTree.AddAnonymousSection(config, typ)
}

// AddList delegates to the default tree. See Tree for details.
func AddList(config, section, option, value string) error {
	return defaultTree.AddList(config, section, option, value)
}

// DelList delegates to the default tree. See Tree for details.
func DelList(config, section, option, value string) (bool, error) {
	return defaultTree.DelList(config, section, option, value)
}

// RenameSection delegates to the default tree. See Tree for details.
func RenameSection(config, oldName, newName string) error {
	return defaultTree.RenameSection(config, oldName, newName)
}

// Show delegates to the default tree. See Tree for details.
func Show(config, section, option string) (string, error) {
	return defaultTree.Show(config, section, option)
}

// ShowPath shows configuration using a UCI-style path (e.g., "network", "network.lan", "network.lan.ipaddr").
func ShowPath(path string) (string, error) {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	switch len(parts) {
	case UCIConfigOnly:
		return Show(parts[UCIPartIndexConfig], "", "")
	case UCIConfigSection:
		return Show(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], "")
	case UCIConfigSectionOption:
		return Show(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption])
	default:
		return "", fmt.Errorf("invalid UCI path: %s", path)
	}
}

// GetPath gets configuration values using a UCI-style path (e.g., "network.lan.ipaddr").
func GetPath(path string) ([]string, bool) {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	if len(parts) != UCIConfigPartCount {
		return nil, false
	}
	return Get(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption])
}

// GetLastPath gets the last configuration value using a UCI-style path (e.g., "network.lan.ipaddr").
func GetLastPath(path string) (string, bool) {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	if len(parts) != UCIConfigPartCount {
		return "", false
	}
	return GetLast(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption])
}

// GetBoolPath gets a boolean configuration value using a UCI-style path (e.g., "network.lan.enabled").
func GetBoolPath(path string) (bool, bool) {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	if len(parts) != UCIConfigPartCount {
		return false, false
	}
	return GetBool(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption])
}

// SetPath sets a configuration value using a UCI-style path (e.g., "network.lan.ipaddr").
func SetPath(path string, values ...string) error {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	if len(parts) != UCIConfigPartCount {
		return fmt.Errorf("invalid UCI path for set: %s", path)
	}
	return SetType(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption], TypeOption, values...)
}

// DelPath deletes a configuration using a UCI-style path (e.g., "network", "network.lan", "network.lan.ipaddr").
func DelPath(path string) error {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	switch len(parts) {
	case UCIConfigOnly:
		return DelSection(parts[UCIPartIndexConfig], "")
	case UCIConfigSection:
		return DelSection(parts[UCIPartIndexConfig], parts[UCIPartIndexSection])
	case UCIConfigSectionOption:
		return Del(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption])
	default:
		return fmt.Errorf("invalid UCI path: %s", path)
	}
}

// AddListPath adds a value to a list option using a UCI-style path (e.g., "network.lan.dns").
func AddListPath(path string, value string) error {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	if len(parts) != UCIConfigPartCount {
		return fmt.Errorf("invalid UCI path for add_list: %s", path)
	}
	return AddList(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption], value)
}

// DelListPath deletes a value from a list option using a UCI-style path (e.g., "network.lan.dns").
func DelListPath(path string, value string) (bool, error) {
	parts := strings.SplitN(path, ".", UCIConfigPartCount)
	if len(parts) != UCIConfigPartCount {
		return false, fmt.Errorf("invalid UCI path for del_list: %s", path)
	}
	return DelList(parts[UCIPartIndexConfig], parts[UCIPartIndexSection], parts[UCIPartIndexOption], value)
}
