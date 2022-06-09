package twilioService

import (
	"strings"

	"github.com/supertokens/supertokens-golang/ingredients/smsdelivery"
	"github.com/supertokens/supertokens-golang/supertokens"
)

const magicLinkLoginTemplate = `${appname} - Login to your account

Click on this link: ${magicLink}

This is valid for ${time}.`
const otpLoginTemplate = `${appname} - Login to your account

Your OTP to login: ${otp}

This is valid for ${time}.`
const magicLinkAndOtpLoginTemplate = `${appname} - Login to your account

Your OTP to login: ${otp}

OR

Click on this link: ${magicLink}

This is valid for ${time}.`

func getPasswordlessLoginSmsContent(input smsdelivery.PasswordlessLoginType) smsdelivery.TwilioGetContentResult {
	stInstance, err := supertokens.GetInstanceOrThrowError()
	if err != nil {
		panic("Please call supertokens.Init function before using the Middleware")
	}
	return smsdelivery.TwilioGetContentResult{
		Body:          getPasswordlessLoginSmsBody(stInstance.AppInfo.AppName, input.CodeLifetime, input.UrlWithLinkCode, input.UserInputCode),
		ToPhoneNumber: input.PhoneNumber,
	}
}

func getPasswordlessLoginSmsBody(appName string, codeLifetime uint64, urlWithLinkCode *string, userInputCode *string) string {
	var smsBody string

	if urlWithLinkCode != nil && userInputCode != nil {
		smsBody = magicLinkAndOtpLoginTemplate
	} else if urlWithLinkCode != nil {
		smsBody = magicLinkLoginTemplate
	} else if userInputCode != nil {
		smsBody = otpLoginTemplate
	} else {
		// Should never come here
	}

	humanisedCodeLifetime := supertokens.HumaniseMilliseconds(codeLifetime)

	smsBody = strings.Replace(smsBody, "*|MC:SUBJECT|*", "Login to your account", -1)
	smsBody = strings.Replace(smsBody, "${appname}", appName, -1)
	smsBody = strings.Replace(smsBody, "${time}", humanisedCodeLifetime, -1)
	if urlWithLinkCode != nil {
		smsBody = strings.Replace(smsBody, "${magicLink}", *urlWithLinkCode, -1)
	}
	if userInputCode != nil {
		smsBody = strings.Replace(smsBody, "${otp}", *userInputCode, -1)
	}

	return smsBody
}