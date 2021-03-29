// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qcloud

import (
	"context"
	"database/sql"
	"fmt"

	"yunion.io/x/pkg/errors"

	api "yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudid/models"
	"yunion.io/x/onecloud/pkg/httperrors"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/util/samlutils"
	"yunion.io/x/onecloud/pkg/util/samlutils/idp"
)

func (d *SQcloudSAMLDriver) GetIdpInitiatedLoginData(ctx context.Context, userCred mcclient.TokenCredential, cloudAccountId string, sp *idp.SSAMLServiceProvider) (samlutils.SSAMLIdpInitiatedLoginData, error) {
	data := samlutils.SSAMLIdpInitiatedLoginData{}

	_account, err := models.CloudaccountManager.FetchById(cloudAccountId)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return data, httperrors.NewResourceNotFoundError2("cloudaccount", cloudAccountId)
		}
		return data, httperrors.NewGeneralError(err)
	}
	account := _account.(*models.SCloudaccount)
	if account.Provider != api.CLOUD_PROVIDER_QCLOUD {
		return data, httperrors.NewClientError("cloudaccount %s is %s not %s", account.Id, account.Provider, api.CLOUD_PROVIDER_QCLOUD)
	}
	if account.SAMLAuth.IsFalse() {
		return data, httperrors.NewNotSupportedError("cloudaccount %s not open saml auth", account.Id)
	}

	SAMLProvider, valid := account.IsSAMLProviderValid()
	if !valid {
		return data, httperrors.NewResourceNotReadyError("SAMLProvider for account %s not ready", account.Id)
	}

	roles, err := account.SyncRoles(userCred.GetUserId(), true)
	if err != nil {
		return data, httperrors.NewGeneralError(errors.Wrapf(err, "SyncRole"))
	}

	roleStr := fmt.Sprintf("qcs::cam::uin/%s:roleName/%s,qcs::cam::uin/%s:saml-provider/%s", account.AccountId, roles[0].ExternalId, account.AccountId, SAMLProvider.ExternalId)

	data.NameId = roles[0].Name
	data.NameIdFormat = samlutils.NAME_ID_FORMAT_TRANSIENT
	data.AudienceRestriction = "https://cloud.tencent.com"
	for _, v := range []struct {
		name         string
		friendlyName string
		value        string
	}{
		{
			name:         "https://cloud.tencent.com/SAML/Attributes/Role",
			friendlyName: "RoleEntitlement",
			value:        roleStr,
		},
		{
			name:         "https://cloud.tencent.com/SAML/Attributes/RoleSessionName",
			friendlyName: "RoleSessionName",
			value:        roles[0].Name,
		},
	} {
		data.Attributes = append(data.Attributes, samlutils.SSAMLResponseAttribute{
			Name:         v.name,
			FriendlyName: v.friendlyName,
			Values:       []string{v.value},
		})
	}
	data.RelayState = "https://console.cloud.tencent.com/"

	return data, nil
}

func (d *SQcloudSAMLDriver) GetSpInitiatedLoginData(ctx context.Context, userCred mcclient.TokenCredential, cloudAccountId string, sp *idp.SSAMLServiceProvider) (samlutils.SSAMLSpInitiatedLoginData, error) {
	data := samlutils.SSAMLSpInitiatedLoginData{}
	_account, err := models.CloudaccountManager.FetchById(cloudAccountId)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return data, httperrors.NewResourceNotFoundError2("cloudaccount", cloudAccountId)
		}
		return data, httperrors.NewGeneralError(err)
	}
	account := _account.(*models.SCloudaccount)
	if account.Provider != api.CLOUD_PROVIDER_QCLOUD {
		return data, httperrors.NewClientError("cloudaccount %s is %s not %s", account.Id, account.Provider, api.CLOUD_PROVIDER_QCLOUD)
	}
	if account.SAMLAuth.IsFalse() {
		return data, httperrors.NewNotSupportedError("cloudaccount %s not open saml auth", account.Id)
	}

	SAMLProvider, valid := account.IsSAMLProviderValid()
	if !valid {
		return data, httperrors.NewResourceNotReadyError("SAMLProvider for account %s not ready", account.Id)
	}

	roles, err := account.SyncRoles(userCred.GetUserId(), true)
	if err != nil {
		return data, httperrors.NewGeneralError(errors.Wrapf(err, "SyncRole"))
	}

	roleStr := fmt.Sprintf("qcs::cam::uin/%s:roleName/%s,qcs::cam::uin/%s:saml-provider/%s", account.AccountId, roles[0].ExternalId, account.AccountId, SAMLProvider.ExternalId)

	data.NameId = roles[0].Name
	data.NameIdFormat = samlutils.NAME_ID_FORMAT_TRANSIENT
	data.AudienceRestriction = "https://cloud.tencent.com"
	for _, v := range []struct {
		name         string
		friendlyName string
		value        string
	}{
		{
			name:         "https://cloud.tencent.com/SAML/Attributes/Role",
			friendlyName: "RoleEntitlement",
			value:        roleStr,
		},
		{
			name:         "https://cloud.tencent.com/SAML/Attributes/RoleSessionName",
			friendlyName: "RoleSessionName",
			value:        roles[0].Name,
		},
	} {
		data.Attributes = append(data.Attributes, samlutils.SSAMLResponseAttribute{
			Name:         v.name,
			FriendlyName: v.friendlyName,
			Values:       []string{v.value},
		})
	}

	return data, nil
}