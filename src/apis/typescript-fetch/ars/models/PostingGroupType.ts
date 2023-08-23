/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Transaction group information. Ever posting could represent a. Regular transaction b. Group header - has Group Id, Count, Descripton and Includes Group Memebers elements c. Group transaction - has SubGroupId element pointing to Group Header Id
 * @export
 * @interface PostingGroupType
 */
export interface PostingGroupType {
    /**
     * Group Id.
     * @type {string}
     * @memberof PostingGroupType
     */
    groupId?: string;
    /**
     * Number of transactions in group.
     * @type {number}
     * @memberof PostingGroupType
     */
    groupCount?: number;
    /**
     * Group Description.
     * @type {string}
     * @memberof PostingGroupType
     */
    groupDescription?: string;
    /**
     * Indicates to which group transaction belongs.
     * @type {string}
     * @memberof PostingGroupType
     */
    subGroupId?: string;
}

/**
 * Check if a given object implements the PostingGroupType interface.
 */
export function instanceOfPostingGroupType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostingGroupTypeFromJSON(json: any): PostingGroupType {
    return PostingGroupTypeFromJSONTyped(json, false);
}

export function PostingGroupTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostingGroupType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'groupId': !exists(json, 'groupId') ? undefined : json['groupId'],
        'groupCount': !exists(json, 'groupCount') ? undefined : json['groupCount'],
        'groupDescription': !exists(json, 'groupDescription') ? undefined : json['groupDescription'],
        'subGroupId': !exists(json, 'subGroupId') ? undefined : json['subGroupId'],
    };
}

export function PostingGroupTypeToJSON(value?: PostingGroupType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'groupId': value.groupId,
        'groupCount': value.groupCount,
        'groupDescription': value.groupDescription,
        'subGroupId': value.subGroupId,
    };
}

