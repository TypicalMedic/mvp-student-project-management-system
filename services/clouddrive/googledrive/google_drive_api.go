package googledrive

import (
	"log"
	"mvp-student-project-management-system/services/googleapi"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type googleDriveApi struct {
	API *drive.Service
}

func InitDrive(googleAPI googleapi.GoogleAPI) googleDriveApi {
	api, err := drive.NewService(googleAPI.Context, option.WithHTTPClient(googleAPI.Client))
	c := googleDriveApi{api}
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return c
}

func (d *googleDriveApi) CreateFolder(parentFolderId string, foldername string) (*drive.File, error) {
	fileMetadata := &drive.File{
		Name:     foldername,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentFolderId},
	}
	file, err := d.API.Files.Create(fileMetadata).Fields("id").Do()
	if err == nil {
		return file, nil
	}
	return nil, err
}

func (d *googleDriveApi) DeleteFolder(folderId string) error {
	err := d.API.Files.Delete(folderId).Do()
	if err == nil {
		return nil
	}
	return err
}

func (d *googleDriveApi) AddTextFileToFolder(parentFolderId string, filename string, text string) (*drive.File, error) {
	fileMetadata := &drive.File{
		Name:     filename,
		MimeType: "application/vnd.google-apps.document",
		Parents:  []string{parentFolderId},
	}

	r := strings.NewReader(text)
	file, err := d.API.Files.Create(fileMetadata).Media(r).Fields("id").Do()
	if err == nil {
		return file, nil
	}
	return nil, err
}

func (d *googleDriveApi) AddCommentToFile(fileId, commentText string) error {
	comment := drive.Comment{
		Content: commentText,
	}
	_, err := d.API.Comments.Create(fileId, &comment).Fields("id").Do()
	if err == nil {
		return nil
	}
	return err
}
