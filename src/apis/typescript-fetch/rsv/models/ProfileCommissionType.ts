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
import type { CommissionPaymentMethods } from './CommissionPaymentMethods';
import {
    CommissionPaymentMethodsFromJSON,
    CommissionPaymentMethodsFromJSONTyped,
    CommissionPaymentMethodsToJSON,
} from './CommissionPaymentMethods';

/**
 * This is the preconfigured routing instruction type.
 * @export
 * @interface ProfileCommissionType
 */
export interface ProfileCommissionType {
    /**
     * Hotel Code for the commission being used for a profile.
     * @type {string}
     * @memberof ProfileCommissionType
     */
    hotelId?: string;
    /**
     * commission Code for a profile.
     * @type {string}
     * @memberof ProfileCommissionType
     */
    commissionCode?: string;
    /**
     * Bank account used for the commission for a profile.
     * @type {string}
     * @memberof ProfileCommissionType
     */
    bankAccount?: string;
    /**
     * 
     * @type {CommissionPaymentMethods}
     * @memberof ProfileCommissionType
     */
    paymentMethod?: CommissionPaymentMethods;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof ProfileCommissionType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof ProfileCommissionType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof ProfileCommissionType
     */
    decimalPlaces?: number;
}

/**
 * Check if a given object implements the ProfileCommissionType interface.
 */
export function instanceOfProfileCommissionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileCommissionTypeFromJSON(json: any): ProfileCommissionType {
    return ProfileCommissionTypeFromJSONTyped(json, false);
}

export function ProfileCommissionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileCommissionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'commissionCode': !exists(json, 'commissionCode') ? undefined : json['commissionCode'],
        'bankAccount': !exists(json, 'bankAccount') ? undefined : json['bankAccount'],
        'paymentMethod': !exists(json, 'paymentMethod') ? undefined : CommissionPaymentMethodsFromJSON(json['paymentMethod']),
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
    };
}

export function ProfileCommissionTypeToJSON(value?: ProfileCommissionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'commissionCode': value.commissionCode,
        'bankAccount': value.bankAccount,
        'paymentMethod': CommissionPaymentMethodsToJSON(value.paymentMethod),
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
    };
}

