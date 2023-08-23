/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { FeeType } from './FeeType';
import {
    FeeTypeFromJSON,
    FeeTypeFromJSONTyped,
    FeeTypeToJSON,
} from './FeeType';

/**
 * A collection of fees or service charges.
 * @export
 * @interface FeesType
 */
export interface FeesType {
    /**
     * An individual fee or service charge.
     * @type {Array<FeeType>}
     * @memberof FeesType
     */
    fee?: Array<FeeType>;
    /**
     * A monetary amount.
     * @type {number}
     * @memberof FeesType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof FeesType
     */
    currencyCode?: string;
}

/**
 * Check if a given object implements the FeesType interface.
 */
export function instanceOfFeesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FeesTypeFromJSON(json: any): FeesType {
    return FeesTypeFromJSONTyped(json, false);
}

export function FeesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FeesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'fee': !exists(json, 'fee') ? undefined : ((json['fee'] as Array<any>).map(FeeTypeFromJSON)),
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
    };
}

export function FeesTypeToJSON(value?: FeesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'fee': value.fee === undefined ? undefined : ((value.fee as Array<any>).map(FeeTypeToJSON)),
        'amount': value.amount,
        'currencyCode': value.currencyCode,
    };
}

