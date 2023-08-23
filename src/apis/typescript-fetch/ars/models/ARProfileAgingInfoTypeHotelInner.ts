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
import type { ARAccountType } from './ARAccountType';
import {
    ARAccountTypeFromJSON,
    ARAccountTypeFromJSONTyped,
    ARAccountTypeToJSON,
} from './ARAccountType';
import type { ARAgingInfoType } from './ARAgingInfoType';
import {
    ARAgingInfoTypeFromJSON,
    ARAgingInfoTypeFromJSONTyped,
    ARAgingInfoTypeToJSON,
} from './ARAgingInfoType';

/**
 * 
 * @export
 * @interface ARProfileAgingInfoTypeHotelInner
 */
export interface ARProfileAgingInfoTypeHotelInner {
    /**
     * 
     * @type {ARAgingInfoType}
     * @memberof ARProfileAgingInfoTypeHotelInner
     */
    totalHotelAging?: ARAgingInfoType;
    /**
     * Information regarding the AR Account.
     * @type {Array<ARAccountType>}
     * @memberof ARProfileAgingInfoTypeHotelInner
     */
    accountAging?: Array<ARAccountType>;
    /**
     * Property Code for the Aging information.
     * @type {string}
     * @memberof ARProfileAgingInfoTypeHotelInner
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the ARProfileAgingInfoTypeHotelInner interface.
 */
export function instanceOfARProfileAgingInfoTypeHotelInner(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ARProfileAgingInfoTypeHotelInnerFromJSON(json: any): ARProfileAgingInfoTypeHotelInner {
    return ARProfileAgingInfoTypeHotelInnerFromJSONTyped(json, false);
}

export function ARProfileAgingInfoTypeHotelInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ARProfileAgingInfoTypeHotelInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'totalHotelAging': !exists(json, 'totalHotelAging') ? undefined : ARAgingInfoTypeFromJSON(json['totalHotelAging']),
        'accountAging': !exists(json, 'accountAging') ? undefined : ((json['accountAging'] as Array<any>).map(ARAccountTypeFromJSON)),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function ARProfileAgingInfoTypeHotelInnerToJSON(value?: ARProfileAgingInfoTypeHotelInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'totalHotelAging': ARAgingInfoTypeToJSON(value.totalHotelAging),
        'accountAging': value.accountAging === undefined ? undefined : ((value.accountAging as Array<any>).map(ARAccountTypeToJSON)),
        'hotelId': value.hotelId,
    };
}

