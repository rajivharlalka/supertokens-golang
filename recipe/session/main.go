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

package session

import (
	"context"
	"errors"
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/jwt/jwtmodels"
	"github.com/supertokens/supertokens-golang/recipe/openid/openidmodels"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Init(config *sessmodels.TypeInput) supertokens.Recipe {
	return recipeInit(config)
}

func CreateNewSessionWithContext(res http.ResponseWriter, userID string, accessTokenPayload map[string]interface{}, sessionData map[string]interface{}, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return sessmodels.SessionContainer{}, err
	}
	return (*instance.RecipeImpl.CreateNewSession)(res, userID, accessTokenPayload, sessionData, userContext)
}

func GetSessionWithContext(req *http.Request, res http.ResponseWriter, options *sessmodels.VerifySessionOptions, userContext supertokens.UserContext) (*sessmodels.SessionContainer, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.GetSession)(req, res, options, userContext)
}

func GetSessionInformationWithContext(sessionHandle string, userContext supertokens.UserContext) (sessmodels.SessionInformation, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return sessmodels.SessionInformation{}, err
	}
	return (*instance.RecipeImpl.GetSessionInformation)(sessionHandle, userContext)
}

func RefreshSessionWithContext(req *http.Request, res http.ResponseWriter, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return sessmodels.SessionContainer{}, err
	}
	return (*instance.RecipeImpl.RefreshSession)(req, res, userContext)
}

func RevokeAllSessionsForUserWithContext(userID string, userContext supertokens.UserContext) ([]string, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.RevokeAllSessionsForUser)(userID, userContext)
}

func GetAllSessionHandlesForUserWithContext(userID string, userContext supertokens.UserContext) ([]string, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.GetAllSessionHandlesForUser)(userID, userContext)
}

func RevokeSessionWithContext(sessionHandle string, userContext supertokens.UserContext) (bool, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return false, err
	}
	return (*instance.RecipeImpl.RevokeSession)(sessionHandle, userContext)
}

func RevokeMultipleSessionsWithContext(sessionHandles []string, userContext supertokens.UserContext) ([]string, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return nil, err
	}
	return (*instance.RecipeImpl.RevokeMultipleSessions)(sessionHandles, userContext)
}

func UpdateSessionDataWithContext(sessionHandle string, newSessionData map[string]interface{}, userContext supertokens.UserContext) error {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return err
	}
	return (*instance.RecipeImpl.UpdateSessionData)(sessionHandle, newSessionData, userContext)
}

func UpdateAccessTokenPayloadWithContext(sessionHandle string, newAccessTokenPayload map[string]interface{}, userContext supertokens.UserContext) error {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return err
	}
	return (*instance.RecipeImpl.UpdateAccessTokenPayload)(sessionHandle, newAccessTokenPayload, userContext)
}

func VerifySession(options *sessmodels.VerifySessionOptions, otherHandler http.HandlerFunc) http.HandlerFunc {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		panic("can't fetch supertokens instance. You should call the supertokens.Init function before using the VerifySession function.")
	}
	return VerifySessionHelper(*instance, options, otherHandler)
}

func GetSessionFromRequestContext(ctx context.Context) *sessmodels.SessionContainer {
	value := ctx.Value(sessmodels.SessionContext)
	if value == nil {
		return nil
	}
	temp := value.(*sessmodels.SessionContainer)
	return temp
}

func CreateJWTWithContext(payload map[string]interface{}, validitySecondsPointer *uint64, userContext supertokens.UserContext) (jwtmodels.CreateJWTResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return jwtmodels.CreateJWTResponse{}, err
	}
	if instance.OpenIdRecipe == nil {
		return jwtmodels.CreateJWTResponse{}, errors.New("CreateJWT cannot be used without enabling the Jwt feature")
	}
	return (*instance.OpenIdRecipe.RecipeImpl.CreateJWT)(payload, validitySecondsPointer, userContext)
}

func GetJWKSWithContext(userContext supertokens.UserContext) (jwtmodels.GetJWKSResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return jwtmodels.GetJWKSResponse{}, err
	}
	if instance.OpenIdRecipe == nil {
		return jwtmodels.GetJWKSResponse{}, errors.New("GetJWKS cannot be used without enabling the Jwt feature")
	}
	return (*instance.OpenIdRecipe.RecipeImpl.GetJWKS)(userContext)
}

func GetOpenIdDiscoveryConfigurationWithContext(userContext supertokens.UserContext) (openidmodels.GetOpenIdDiscoveryConfigurationResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return openidmodels.GetOpenIdDiscoveryConfigurationResponse{}, err
	}
	if instance.OpenIdRecipe == nil {
		return openidmodels.GetOpenIdDiscoveryConfigurationResponse{}, errors.New("GetOpenIdDiscoveryConfiguration cannot be used without enabling the Jwt feature")
	}
	return (*instance.OpenIdRecipe.RecipeImpl.GetOpenIdDiscoveryConfiguration)(userContext)
}

func RegenerateAccessTokenWithContext(accessToken string, newAccessTokenPayload *map[string]interface{}, sessionHandle string, userContext supertokens.UserContext) (sessmodels.RegenerateAccessTokenResponse, error) {
	instance, err := getRecipeInstanceOrThrowError()
	if err != nil {
		return sessmodels.RegenerateAccessTokenResponse{}, err
	}
	return (*instance.RecipeImpl.RegenerateAccessToken)(accessToken, newAccessTokenPayload, userContext)
}

func CreateNewSession(res http.ResponseWriter, userID string, accessTokenPayload map[string]interface{}, sessionData map[string]interface{}) (sessmodels.SessionContainer, error) {
	return CreateNewSessionWithContext(res, userID, accessTokenPayload, sessionData, &map[string]interface{}{})
}

func GetSession(req *http.Request, res http.ResponseWriter, options *sessmodels.VerifySessionOptions) (*sessmodels.SessionContainer, error) {
	return GetSessionWithContext(req, res, options, &map[string]interface{}{})
}

func GetSessionInformation(sessionHandle string) (sessmodels.SessionInformation, error) {
	return GetSessionInformationWithContext(sessionHandle, &map[string]interface{}{})
}

func RefreshSession(req *http.Request, res http.ResponseWriter) (sessmodels.SessionContainer, error) {
	return RefreshSessionWithContext(req, res, &map[string]interface{}{})
}

func RevokeAllSessionsForUser(userID string) ([]string, error) {
	return RevokeAllSessionsForUserWithContext(userID, &map[string]interface{}{})
}

func GetAllSessionHandlesForUser(userID string) ([]string, error) {
	return GetAllSessionHandlesForUserWithContext(userID, &map[string]interface{}{})
}

func RevokeSession(sessionHandle string) (bool, error) {
	return RevokeSessionWithContext(sessionHandle, &map[string]interface{}{})
}

func RevokeMultipleSessions(sessionHandles []string) ([]string, error) {
	return RevokeMultipleSessionsWithContext(sessionHandles, &map[string]interface{}{})
}

func UpdateSessionData(sessionHandle string, newSessionData map[string]interface{}) error {
	return UpdateSessionDataWithContext(sessionHandle, newSessionData, &map[string]interface{}{})
}

func UpdateAccessTokenPayload(sessionHandle string, newAccessTokenPayload map[string]interface{}) error {
	return UpdateAccessTokenPayloadWithContext(sessionHandle, newAccessTokenPayload, &map[string]interface{}{})
}

func CreateJWT(payload map[string]interface{}, validitySecondsPointer *uint64) (jwtmodels.CreateJWTResponse, error) {
	return CreateJWTWithContext(payload, validitySecondsPointer, &map[string]interface{}{})
}

func GetJWKS() (jwtmodels.GetJWKSResponse, error) {
	return GetJWKSWithContext(&map[string]interface{}{})
}

func GetOpenIdDiscoveryConfiguration() (openidmodels.GetOpenIdDiscoveryConfigurationResponse, error) {
	return GetOpenIdDiscoveryConfigurationWithContext(&map[string]interface{}{})
}

func RegenerateAccessToken(accessToken string, newAccessTokenPayload *map[string]interface{}, sessionHandle string) (sessmodels.RegenerateAccessTokenResponse, error) {
	return RegenerateAccessTokenWithContext(accessToken, newAccessTokenPayload, sessionHandle, &map[string]interface{}{})
}
