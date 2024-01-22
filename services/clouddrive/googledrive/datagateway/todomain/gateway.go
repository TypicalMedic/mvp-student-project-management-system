package todomain

import (
	entities "mvp-student-project-management-system/managing-student-project"

	"google.golang.org/api/drive/v3"
)

func Folder(folder drive.File) entities.Folder {
	f := entities.InitFolder(
		folder.Id,
		folder.Name,
		folder.Parents[0],
	)
	return f
}
