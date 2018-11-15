package tree

//Folder structure
type Folder struct {
	path           string
	dev            string
	projects       string
	scripts        string
	metadataFolder string
}

//CreateFolder create Folder instance
func CreateFolder(_path string, isMetadataNeeded bool) *Folder {
	f := &Folder{
		path:           _path,
		dev:            "dev",
		projects:       "projects",
		scripts:        "scripts",
		metadataFolder: "",
	}
	if isMetadataNeeded {
		f.metadataFolder = ".devenv"
	}
	return f
}

// Dev path
func (f Folder) Dev() string {
	return f.path + f.dev
}

// Scripts path
func (f Folder) Scripts() string {
	return f.Dev() + "/" + f.scripts
}

// Projects path
func (f Folder) Projects() string {
	return f.Dev() + "/" + f.projects
}

// MetadataFolder path
func (f Folder) MetadataFolder() (string, bool) {
	var isMetadataExist bool
	var path string

	if f.metadataFolder != "" {
		isMetadataExist = true
		path = f.Dev() + "/" + f.metadataFolder
	} else {
		isMetadataExist = false
		path = ""
	}
	return path, isMetadataExist
}
