/* Copyright (c) 2021, VRAI Labs and/or its affiliates. All rights reserved.
 *
 * This software is licensed under the Apache License, Version 2.0 (the
 * "License") as published by the Apache Software Foundation.
 *
 * You may not use this file except in compliance with the License. You may
 * obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

package api

import (
	"github.com/supertokens/supertokens-golang/recipe/emailverification/evmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func GenerateEmailVerifyToken(apiImplementation evmodels.APIInterface, options evmodels.APIOptions) error {
	if apiImplementation.GenerateEmailVerifyTokenPOST == nil ||
		(*apiImplementation.GenerateEmailVerifyTokenPOST) == nil {
		options.OtherHandler(options.Res, options.Req)
		return nil
	}

	response, err := (*apiImplementation.GenerateEmailVerifyTokenPOST)(options, &map[string]interface{}{})
	if err != nil {
		return err
	}
	if response.EmailAlreadyVerifiedError != nil {
		return supertokens.Send200Response(options.Res, map[string]interface{}{
			"status": "EMAIL_ALREADY_VERIFIED_ERROR",
		})
	} else {
		return supertokens.Send200Response(options.Res, map[string]interface{}{
			"status": "OK",
		})
	}
}
