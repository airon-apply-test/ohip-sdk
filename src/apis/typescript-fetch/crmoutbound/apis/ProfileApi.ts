/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ExceptionDetailType,
  PostMembershipNumberRequest,
  Status,
} from '../models';
import {
    ExceptionDetailTypeFromJSON,
    ExceptionDetailTypeToJSON,
    PostMembershipNumberRequestFromJSON,
    PostMembershipNumberRequestToJSON,
    StatusFromJSON,
    StatusToJSON,
} from '../models';

export interface PostMembershipNumberOperationRequest {
    extSystemCode?: string;
    membershipType?: string;
    authorization?: string;
    xAppKey?: string;
    membershipNumber?: PostMembershipNumberRequest;
    xTransactionId?: string;
    acceptLanguage?: string;
}

/**
 * 
 */
export class ProfileApi extends runtime.BaseAPI {

    /**
     * This API will generate a new membership number. <p><strong>OperationId:</strong>postMembershipNumber</p>
     * Generate membership number
     */
    async postMembershipNumberRaw(requestParameters: PostMembershipNumberOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Status>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json;charset=UTF-8';

        if (requestParameters.authorization !== undefined && requestParameters.authorization !== null) {
            headerParameters['authorization'] = String(requestParameters.authorization);
        }

        if (requestParameters.xAppKey !== undefined && requestParameters.xAppKey !== null) {
            headerParameters['x-app-key'] = String(requestParameters.xAppKey);
        }

        if (requestParameters.xTransactionId !== undefined && requestParameters.xTransactionId !== null) {
            headerParameters['x-transactionId'] = String(requestParameters.xTransactionId);
        }

        if (requestParameters.acceptLanguage !== undefined && requestParameters.acceptLanguage !== null) {
            headerParameters['Accept-Language'] = String(requestParameters.acceptLanguage);
        }

        const response = await this.request({
            path: `/externalSystems/{extSystemCode}/memberships/{membershipType}/membershipIDs`.replace(`{${"extSystemCode"}}`, encodeURIComponent(String(requestParameters.extSystemCode))).replace(`{${"membershipType"}}`, encodeURIComponent(String(requestParameters.membershipType))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: PostMembershipNumberRequestToJSON(requestParameters.membershipNumber),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => StatusFromJSON(jsonValue));
    }

    /**
     * This API will generate a new membership number. <p><strong>OperationId:</strong>postMembershipNumber</p>
     * Generate membership number
     */
    async postMembershipNumber(requestParameters: PostMembershipNumberOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Status> {
        const response = await this.postMembershipNumberRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
