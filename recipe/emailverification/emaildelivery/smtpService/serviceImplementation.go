package smtpService

import (
	"errors"

	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func MakeServiceImplementation(config emaildelivery.SMTPServiceConfig) emaildelivery.SMTPServiceInterface {
	sendRawEmail := func(input emaildelivery.SMTPGetContentResult, userContext supertokens.UserContext) error {
		return emaildelivery.SendSMTPEmail(config, input)
	}

	getContent := func(input emaildelivery.EmailType, userContext supertokens.UserContext) (emaildelivery.SMTPGetContentResult, error) {
		if input.EmailVerification != nil {
			return getEmailVerifyEmailContent(*input.EmailVerification)
		} else {
			return emaildelivery.SMTPGetContentResult{}, errors.New("should never come here")
		}
	}

	return emaildelivery.SMTPServiceInterface{
		SendRawEmail: &sendRawEmail,
		GetContent:   &getContent,
	}
}