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
 * Collection of comp postings.
 * @export
 * @interface CompPostingsType
 */
export interface CompPostingsType {
    /**
     * Authorizer name of the Comp Account.
     * @type {string}
     * @memberof CompPostingsType
     */
    authorizer?: string;
    /**
     * Approval status of the comp account.
     * @type {string}
     * @memberof CompPostingsType
     */
    approvalStatus?: string;
}

/**
 * Check if a given object implements the CompPostingsType interface.
 */
export function instanceOfCompPostingsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CompPostingsTypeFromJSON(json: any): CompPostingsType {
    return CompPostingsTypeFromJSONTyped(json, false);
}

export function CompPostingsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompPostingsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'authorizer': !exists(json, 'authorizer') ? undefined : json['authorizer'],
        'approvalStatus': !exists(json, 'approvalStatus') ? undefined : json['approvalStatus'],
    };
}

export function CompPostingsTypeToJSON(value?: CompPostingsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'authorizer': value.authorizer,
        'approvalStatus': value.approvalStatus,
    };
}

