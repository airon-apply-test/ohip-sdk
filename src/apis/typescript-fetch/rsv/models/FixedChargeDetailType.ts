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
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';

/**
 * Fixed charge amount could be specified by flat fee or be a percentage of the rate amount.
 * @export
 * @interface FixedChargeDetailType
 */
export interface FixedChargeDetailType {
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof FixedChargeDetailType
     */
    transaction?: CodeDescriptionType;
    /**
     * Quantity of the product.
     * @type {number}
     * @memberof FixedChargeDetailType
     */
    quantity?: number;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof FixedChargeDetailType
     */
    chargeAmount?: CurrencyAmountType;
    /**
     * Percentage of the rate amount.
     * @type {number}
     * @memberof FixedChargeDetailType
     */
    percent?: number;
    /**
     * Additional information regarding the fixed charge.
     * @type {string}
     * @memberof FixedChargeDetailType
     */
    supplement?: string;
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof FixedChargeDetailType
     */
    article?: CodeDescriptionType;
    /**
     * Holds number of comp or cash room night to allocate.
     * @type {number}
     * @memberof FixedChargeDetailType
     */
    roomNights?: number;
}

/**
 * Check if a given object implements the FixedChargeDetailType interface.
 */
export function instanceOfFixedChargeDetailType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FixedChargeDetailTypeFromJSON(json: any): FixedChargeDetailType {
    return FixedChargeDetailTypeFromJSONTyped(json, false);
}

export function FixedChargeDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FixedChargeDetailType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transaction': !exists(json, 'transaction') ? undefined : CodeDescriptionTypeFromJSON(json['transaction']),
        'quantity': !exists(json, 'quantity') ? undefined : json['quantity'],
        'chargeAmount': !exists(json, 'chargeAmount') ? undefined : CurrencyAmountTypeFromJSON(json['chargeAmount']),
        'percent': !exists(json, 'percent') ? undefined : json['percent'],
        'supplement': !exists(json, 'supplement') ? undefined : json['supplement'],
        'article': !exists(json, 'article') ? undefined : CodeDescriptionTypeFromJSON(json['article']),
        'roomNights': !exists(json, 'roomNights') ? undefined : json['roomNights'],
    };
}

export function FixedChargeDetailTypeToJSON(value?: FixedChargeDetailType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transaction': CodeDescriptionTypeToJSON(value.transaction),
        'quantity': value.quantity,
        'chargeAmount': CurrencyAmountTypeToJSON(value.chargeAmount),
        'percent': value.percent,
        'supplement': value.supplement,
        'article': CodeDescriptionTypeToJSON(value.article),
        'roomNights': value.roomNights,
    };
}

