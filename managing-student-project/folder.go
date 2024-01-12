package managingstudentproject

type Folder struct {
	id       string
	name     string
	parentId string
}

func InitFolder(id string, name string, prnt string) Folder {
	return Folder{id: id, name: name, parentId: prnt}
}
