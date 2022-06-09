package twilioService

import (
	"errors"

	"github.com/supertokens/supertokens-golang/ingredients/smsdelivery"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func MakeTwilioService(config smsdelivery.TwilioTypeInput) (smsdelivery.SmsDeliveryInterface, error) {
	config, err := smsdelivery.NormaliseTwilioTypeInput(config)

	if err != nil {
		return smsdelivery.SmsDeliveryInterface{}, err
	}

	serviceImpl := MakeServiceImplementation(config.TwilioSettings)

	if config.Override != nil {
		serviceImpl = config.Override(serviceImpl)
	}

	sendSms := func(input smsdelivery.SmsType, userContext supertokens.UserContext) error {
		if input.PasswordlessLogin != nil {
			content, err := (*serviceImpl.GetContent)(input, userContext)
			if err != nil {
				return err
			}
			return (*serviceImpl.SendRawSms)(content, userContext)

		} else {
			return errors.New("should never come here")
		}
	}

	return smsdelivery.SmsDeliveryInterface{
		SendSms: &sendSms,
	}, nil
}