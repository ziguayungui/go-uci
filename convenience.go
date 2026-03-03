package uci

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
