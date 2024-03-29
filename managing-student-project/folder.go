package managingstudentproject

type Folder struct {
	id       string
	name     string
	parentId string
}

func (f *Folder) Id() string {
	return f.id
}

func (f *Folder) Name() string {
	return f.name
}

func (f *Folder) ParentId() string {
	return f.parentId
}

func (f *Folder) SetId(id string) {
	f.id = id
}

func (f *Folder) SetName(name string) {
	f.name = name
}

func (f *Folder) SetParentId(parentId string) {
	f.parentId = parentId
}

func InitFolder(id string, name string, prnt string) Folder {
	return Folder{id: id, name: name, parentId: prnt}
}
