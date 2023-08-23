/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { DirectBillingType } from './DirectBillingType';
import {
    DirectBillingTypeFromJSON,
    DirectBillingTypeFromJSONTyped,
    DirectBillingTypeToJSON,
} from './DirectBillingType';

/**
 * Contains cashiering related details for the profile
 * @export
 * @interface ProfileCashieringType
 */
export interface ProfileCashieringType {
    /**
     * Contains the direct billing related information for the profile.
     * @type {Array<DirectBillingType>}
     * @memberof ProfileCashieringType
     */
    directBillingList?: Array<DirectBillingType>;
    /**
     * Contains the auto folio settlement type for the profile.
     * @type {string}
     * @memberof ProfileCashieringType
     */
    autoFolioSettlementType?: string;
    /**
     * Ability to define on the account, the number of days after which the invoice should be paid.
     * @type {number}
     * @memberof ProfileCashieringType
     */
    paymentDueDays?: number;
    /**
     * Tax1 Percentage for Collecting Agent.
     * @type {number}
     * @memberof ProfileCashieringType
     */
    taxPercent1?: number;
    /**
     * Tax2 Percentage for Collecting Agent.
     * @type {number}
     * @memberof ProfileCashieringType
     */
    taxPercent2?: number;
    /**
     * Tax3 Percentage for Collecting Agent.
     * @type {number}
     * @memberof ProfileCashieringType
     */
    taxPercent3?: number;
    /**
     * Tax4 Percentage for Collecting Agent.
     * @type {number}
     * @memberof ProfileCashieringType
     */
    taxPercent4?: number;
    /**
     * Tax5 Percentage for Collecting Agent.
     * @type {number}
     * @memberof ProfileCashieringType
     */
    taxPercent5?: number;
    /**
     * Account Receivables Central Number.
     * @type {string}
     * @memberof ProfileCashieringType
     */
    arNoCentral?: string;
    /**
     * Reference Currency.
     * @type {string}
     * @memberof ProfileCashieringType
     */
    referenceCurrency?: string;
    /**
     * VAT Offset Flag.
     * @type {string}
     * @memberof ProfileCashieringType
     */
    vATOffsetYN?: string;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof ProfileCashieringType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof ProfileCashieringType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof ProfileCashieringType
     */
    decimalPlaces?: number;
}

/**
 * Check if a given object implements the ProfileCashieringType interface.
 */
export function instanceOfProfileCashieringType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileCashieringTypeFromJSON(json: any): ProfileCashieringType {
    return ProfileCashieringTypeFromJSONTyped(json, false);
}

export function ProfileCashieringTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileCashieringType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'directBillingList': !exists(json, 'directBillingList') ? undefined : ((json['directBillingList'] as Array<any>).map(DirectBillingTypeFromJSON)),
        'autoFolioSettlementType': !exists(json, 'autoFolioSettlementType') ? undefined : json['autoFolioSettlementType'],
        'paymentDueDays': !exists(json, 'paymentDueDays') ? undefined : json['paymentDueDays'],
        'taxPercent1': !exists(json, 'taxPercent1') ? undefined : json['taxPercent1'],
        'taxPercent2': !exists(json, 'taxPercent2') ? undefined : json['taxPercent2'],
        'taxPercent3': !exists(json, 'taxPercent3') ? undefined : json['taxPercent3'],
        'taxPercent4': !exists(json, 'taxPercent4') ? undefined : json['taxPercent4'],
        'taxPercent5': !exists(json, 'taxPercent5') ? undefined : json['taxPercent5'],
        'arNoCentral': !exists(json, 'arNoCentral') ? undefined : json['arNoCentral'],
        'referenceCurrency': !exists(json, 'referenceCurrency') ? undefined : json['referenceCurrency'],
        'vATOffsetYN': !exists(json, 'vATOffsetYN') ? undefined : json['vATOffsetYN'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
    };
}

export function ProfileCashieringTypeToJSON(value?: ProfileCashieringType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'directBillingList': value.directBillingList === undefined ? undefined : ((value.directBillingList as Array<any>).map(DirectBillingTypeToJSON)),
        'autoFolioSettlementType': value.autoFolioSettlementType,
        'paymentDueDays': value.paymentDueDays,
        'taxPercent1': value.taxPercent1,
        'taxPercent2': value.taxPercent2,
        'taxPercent3': value.taxPercent3,
        'taxPercent4': value.taxPercent4,
        'taxPercent5': value.taxPercent5,
        'arNoCentral': value.arNoCentral,
        'referenceCurrency': value.referenceCurrency,
        'vATOffsetYN': value.vATOffsetYN,
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
    };
}

