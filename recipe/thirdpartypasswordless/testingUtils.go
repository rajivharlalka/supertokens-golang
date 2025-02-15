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

package thirdpartypasswordless

import (
	"errors"

	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"github.com/supertokens/supertokens-golang/test/unittesting"
)

func resetAll() {
	supertokens.ResetForTest()
	ResetForTest()
	session.ResetForTest()
}

func BeforeEach() {
	unittesting.KillAllST()
	resetAll()
	unittesting.SetUpST()
}

func AfterEach() {
	unittesting.KillAllST()
	resetAll()
	unittesting.CleanST()
}

var customProvider1 = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext supertokens.UserContext) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
				Params: map[string]interface{}{
					"scope":     "test",
					"client_id": "supertokens",
				},
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext supertokens.UserContext) (tpmodels.UserInfo, error) {
				return tpmodels.UserInfo{
					ID: "user",
					Email: &tpmodels.EmailStruct{
						ID:         "email@test.com",
						IsVerified: true,
					},
				}, nil
			},
			GetClientId: func(userContext supertokens.UserContext) string {
				return "supertokens"
			},
		}
	},
}

var mockThirdPartyProvider1 = tpmodels.TypeProvider{
	ID: "mock1",
}

var mockThirdPartyProvider2 = tpmodels.TypeProvider{
	ID: "mock2",
}

var signinupCustomProvider1 = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext supertokens.UserContext) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext supertokens.UserContext) (tpmodels.UserInfo, error) {
				return tpmodels.UserInfo{
					ID: "user",
					Email: &tpmodels.EmailStruct{
						ID:         "email@test.com",
						IsVerified: true,
					},
				}, nil
			},
			GetClientId: func(userContext supertokens.UserContext) string {
				return "supertokens"
			},
		}
	},
}

var signinupCustomProvider3 = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext supertokens.UserContext) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext supertokens.UserContext) (tpmodels.UserInfo, error) {
				return tpmodels.UserInfo{
					ID: "user",
				}, nil
			},
			GetClientId: func(userContext supertokens.UserContext) string {
				return "supertokens"
			},
		}
	},
}

var signinupCustomProvider4 = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext supertokens.UserContext) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext supertokens.UserContext) (tpmodels.UserInfo, error) {
				return tpmodels.UserInfo{}, errors.New("error from getProfileInfo")
			},
			GetClientId: func(userContext supertokens.UserContext) string {
				return "supertokens"
			},
		}
	},
}

var signinupCustomProvider5 = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext supertokens.UserContext) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext supertokens.UserContext) (tpmodels.UserInfo, error) {
				return tpmodels.UserInfo{
					ID: "user",
					Email: &tpmodels.EmailStruct{
						ID:         "email@test.com",
						IsVerified: false,
					},
				}, nil
			},
			GetClientId: func(userContext supertokens.UserContext) string {
				return "supertokens"
			},
		}
	},
}

var signinupCustomProvider6 = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext *map[string]interface{}) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext *map[string]interface{}) (tpmodels.UserInfo, error) {
				if authCodeResponse.(map[string]interface{})["access_token"] == nil {
					return tpmodels.UserInfo{}, nil
				}
				return tpmodels.UserInfo{
					ID: "user",
					Email: &tpmodels.EmailStruct{
						ID:         "email@test.com",
						IsVerified: true,
					},
				}, nil
			},
			GetClientId: func(userContext *map[string]interface{}) string {
				return "supertokens"
			},
		}
	},
}

var userTestCustomProvider = tpmodels.TypeProvider{
	ID: "custom",
	Get: func(redirectURI, authCodeFromRequest *string, userContext *map[string]interface{}) tpmodels.TypeProviderGetResponse {
		return tpmodels.TypeProviderGetResponse{
			AccessTokenAPI: tpmodels.AccessTokenAPI{
				URL: "https://test.com/oauth/token",
			},
			AuthorisationRedirect: tpmodels.AuthorisationRedirect{
				URL: "https://test.com/oauth/auth",
			},
			GetProfileInfo: func(authCodeResponse interface{}, userContext *map[string]interface{}) (tpmodels.UserInfo, error) {
				return tpmodels.UserInfo{
					ID: authCodeResponse.(map[string]interface{})["id"].(string),
					Email: &tpmodels.EmailStruct{
						ID:         authCodeResponse.(map[string]interface{})["email"].(string),
						IsVerified: true,
					},
				}, nil
			},
			GetClientId: func(userContext *map[string]interface{}) string {
				return "supertokens"
			},
		}
	},
}
