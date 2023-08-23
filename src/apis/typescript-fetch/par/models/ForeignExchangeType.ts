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
/**
 * Foreign Exchange info.
 * @export
 * @interface ForeignExchangeType
 */
export interface ForeignExchangeType {
    /**
     * Exchange Type for the postings.
     * @type {string}
     * @memberof ForeignExchangeType
     */
    type?: string;
    /**
     * Effective Exchange date for the foreign currency posting.
     * @type {Date}
     * @memberof ForeignExchangeType
     */
    effectiveDate?: Date;
}

/**
 * Check if a given object implements the ForeignExchangeType interface.
 */
export function instanceOfForeignExchangeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ForeignExchangeTypeFromJSON(json: any): ForeignExchangeType {
    return ForeignExchangeTypeFromJSONTyped(json, false);
}

export function ForeignExchangeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ForeignExchangeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'type': !exists(json, 'type') ? undefined : json['type'],
        'effectiveDate': !exists(json, 'effectiveDate') ? undefined : (new Date(json['effectiveDate'])),
    };
}

export function ForeignExchangeTypeToJSON(value?: ForeignExchangeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'type': value.type,
        'effectiveDate': value.effectiveDate === undefined ? undefined : (value.effectiveDate.toISOString().substr(0,10)),
    };
}

