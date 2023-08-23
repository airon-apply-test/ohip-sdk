/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { PolicyBasisTypeType } from './PolicyBasisTypeType';
import {
    PolicyBasisTypeTypeFromJSON,
    PolicyBasisTypeTypeFromJSONTyped,
    PolicyBasisTypeTypeToJSON,
} from './PolicyBasisTypeType';

/**
 * Defines the percentage basis for calculating the fee amount or the amount.
 * @export
 * @interface PolicyAmountPercentType
 */
export interface PolicyAmountPercentType {
    /**
     * 
     * @type {PolicyBasisTypeType}
     * @memberof PolicyAmountPercentType
     */
    basisType?: PolicyBasisTypeType;
    /**
     * The number of nights of the hotel stay that are used to calculate the fee amount.
     * @type {number}
     * @memberof PolicyAmountPercentType
     */
    nights?: number;
    /**
     * The percentage used to calculate the amount.
     * @type {number}
     * @memberof PolicyAmountPercentType
     */
    percent?: number;
    /**
     * A monetary amount.
     * @type {number}
     * @memberof PolicyAmountPercentType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof PolicyAmountPercentType
     */
    currencyCode?: string;
}

/**
 * Check if a given object implements the PolicyAmountPercentType interface.
 */
export function instanceOfPolicyAmountPercentType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PolicyAmountPercentTypeFromJSON(json: any): PolicyAmountPercentType {
    return PolicyAmountPercentTypeFromJSONTyped(json, false);
}

export function PolicyAmountPercentTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicyAmountPercentType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'basisType': !exists(json, 'basisType') ? undefined : PolicyBasisTypeTypeFromJSON(json['basisType']),
        'nights': !exists(json, 'nights') ? undefined : json['nights'],
        'percent': !exists(json, 'percent') ? undefined : json['percent'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
    };
}

export function PolicyAmountPercentTypeToJSON(value?: PolicyAmountPercentType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'basisType': PolicyBasisTypeTypeToJSON(value.basisType),
        'nights': value.nights,
        'percent': value.percent,
        'amount': value.amount,
        'currencyCode': value.currencyCode,
    };
}

