package folder

//easyjson:json
type Folder struct {
	Dir     string   `json:"dir,nocopy"`
	Files   []string `json:"files,nocopy"`
	Folders []Folder `json:"folders,nocopy"`
}
