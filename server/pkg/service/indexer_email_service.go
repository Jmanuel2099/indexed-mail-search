package service

import (
	"fmt"
	"indexed-mail-search/server/pkg/domain"
	"indexed-mail-search/server/pkg/handlers/contracts"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const emailFolderPath = "../enron_mail_20110402/maildir/"

// const emailFolderPath = "../index_test/maildir/"

// IndexerEmailService is the struc that will communicate with the datasource
type IndexerEmailService struct {
	datasource contracts.IEmail
}

// NewIndexerService works as the conntrucutor of the IndexerEmailService struc
func NewIndexerService(ds contracts.IEmail) *IndexerEmailService {
	return &IndexerEmailService{
		datasource: ds,
	}
}

// GetMailUsers gets the users in the email database and then reads the emails
// per user with the ProcessMailsByUser method
func (ies *IndexerEmailService) GetMailUsers() ([]string, error) {
	var mailUsers []string

	dirs, err := os.ReadDir(emailFolderPath)
	if err != nil {

		return nil, err
	}

	for _, dir := range dirs {
		mailUsers = append(mailUsers, dir.Name())
	}

	return mailUsers, nil
}

// ProcessMailsByUser reads and processes the mails that a user has to return
// them as Emails struts
func (ies *IndexerEmailService) ProcessMailsByUser(user string) ([]domain.Email, error) {
	var emails []domain.Email
	path := emailFolderPath + "/" + user

	err := filepath.Walk(path, readEmails(&emails))
	if err != nil {
		return nil, err
	}

	return emails, nil
}

// IndexEmails is the function that indexes the emails in a datasource
func (ies *IndexerEmailService) IndexEmails(records []domain.Email) error {
	reponse, err := ies.datasource.CreateEmails(records)
	if err != nil {
		return err
	}
	fmt.Printf("Message" + reponse.Message + "Count" + strconv.Itoa(reponse.RecordCount))

	return nil
}

// readEmail it goes through all the files that a user has to process all the mail
// files that he has
func readEmails(emails *[]domain.Email) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			email, err := processEmailFile(path)
			if err != nil {
				return nil
			}
			*emails = append(*emails, *email)
		}
		return nil
	}

}

// processEmailFile reads the content of a email containing the information of an
// email and then processes it
func processEmailFile(emailFilepath string) (*domain.Email, error) {
	emailContent, err := os.ReadFile(emailFilepath)
	if err != nil {
		return nil, err
	}
	return mapStringToEmail(string(emailContent)), nil
}

// mapStringToEmail maps the content of an email that is in a string to an Email struc
func mapStringToEmail(emailString string) *domain.Email {
	detailsAndContent := strings.SplitN(string(emailString), "\r\n\r\n", 2)
	details := strings.Split(detailsAndContent[0], "\r\n")

	newEmail := &domain.Email{}
	for _, detail := range details {
		detailValue := strings.SplitN(detail, ": ", 2)
		switch detailValue[0] {
		case "Message-ID":
			newEmail.MessageID = detailValue[1]
		case "Date":
			newEmail.Date = detailValue[1]
		case "From":
			newEmail.From = detailValue[1]
		case "To":
			newEmail.To = detailValue[1]
		case "Subject":
			newEmail.Subject = detailValue[1]
		default:
			continue
		}
	}
	newEmail.Content = detailsAndContent[1]

	return newEmail
}
