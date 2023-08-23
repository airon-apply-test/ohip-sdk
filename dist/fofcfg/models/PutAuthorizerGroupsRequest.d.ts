/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { AuthorizerGroupsCriteriaType } from './AuthorizerGroupsCriteriaType';
import type { InstanceLink } from './InstanceLink';
import type { WarningType } from './WarningType';
/**
 *
 * @export
 * @interface PutAuthorizerGroupsRequest
 */
export interface PutAuthorizerGroupsRequest {
    /**
     *
     * @type {AuthorizerGroupsCriteriaType}
     * @memberof PutAuthorizerGroupsRequest
     */
    criteria?: AuthorizerGroupsCriteriaType;
    /**
     *
     * @type {Array<InstanceLink>}
     * @memberof PutAuthorizerGroupsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutAuthorizerGroupsRequest
     */
    warnings?: Array<WarningType>;
}
/**
 * Check if a given object implements the PutAuthorizerGroupsRequest interface.
 */
export declare function instanceOfPutAuthorizerGroupsRequest(value: object): boolean;
export declare function PutAuthorizerGroupsRequestFromJSON(json: any): PutAuthorizerGroupsRequest;
export declare function PutAuthorizerGroupsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutAuthorizerGroupsRequest;
export declare function PutAuthorizerGroupsRequestToJSON(value?: PutAuthorizerGroupsRequest | null): any;
