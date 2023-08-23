/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ARAccountStatusType } from './ARAccountStatusType';
import {
    ARAccountStatusTypeFromJSON,
    ARAccountStatusTypeFromJSONTyped,
    ARAccountStatusTypeToJSON,
} from './ARAccountStatusType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Accounts Receivabales Account details type charged for batched posting.
 * @export
 * @interface ARAccountShortInfoType
 */
export interface ARAccountShortInfoType {
    /**
     * Name of the AR Account.
     * @type {string}
     * @memberof ARAccountShortInfoType
     */
    accountName?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ARAccountShortInfoType
     */
    accountId?: UniqueIDType;
    /**
     * The Account Number for the Account.
     * @type {string}
     * @memberof ARAccountShortInfoType
     */
    accountNo?: string;
    /**
     * 
     * @type {ARAccountStatusType}
     * @memberof ARAccountShortInfoType
     */
    status?: ARAccountStatusType;
}

/**
 * Check if a given object implements the ARAccountShortInfoType interface.
 */
export function instanceOfARAccountShortInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ARAccountShortInfoTypeFromJSON(json: any): ARAccountShortInfoType {
    return ARAccountShortInfoTypeFromJSONTyped(json, false);
}

export function ARAccountShortInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARAccountShortInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'accountName': !exists(json, 'accountName') ? undefined : json['accountName'],
        'accountId': !exists(json, 'accountId') ? undefined : UniqueIDTypeFromJSON(json['accountId']),
        'accountNo': !exists(json, 'accountNo') ? undefined : json['accountNo'],
        'status': !exists(json, 'status') ? undefined : ARAccountStatusTypeFromJSON(json['status']),
    };
}

export function ARAccountShortInfoTypeToJSON(value?: ARAccountShortInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'accountName': value.accountName,
        'accountId': UniqueIDTypeToJSON(value.accountId),
        'accountNo': value.accountNo,
        'status': ARAccountStatusTypeToJSON(value.status),
    };
}

