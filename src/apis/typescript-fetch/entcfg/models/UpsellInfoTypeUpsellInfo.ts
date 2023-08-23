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
 * Information of a reservation during the process of upgrade.
 * @export
 * @interface UpsellInfoTypeUpsellInfo
 */
export interface UpsellInfoTypeUpsellInfo {
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    totalUpsellCharge?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    totalUpsellMoneyAmount?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    firstNightUpsellAmount?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    firstNightUpsellCharge?: CurrencyAmountType;
    /**
     * Total Upsell Points.
     * @type {number}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    totalUpsellPoints?: number;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    totalActualRateAmount?: CurrencyAmountType;
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    roomType?: CodeDescriptionType;
    /**
     * Detail description of the Room Type.
     * @type {string}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    roomLongDescription?: string;
    /**
     * Upsell Rule Id.
     * @type {number}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    ruleId?: number;
    /**
     * Upsell rule code
     * @type {string}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    ruleCode?: string;
    /**
     * Upsell rule description
     * @type {string}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    ruleDescription?: string;
    /**
     * Percentage savings resulting from accepting upsell offer. This will be calculated as the ratio of the amount saved by accepting the upsell offer to the actual room rate (i.e. room rate when upsell is not offered)
     * @type {number}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    percentageSavings?: number;
    /**
     * The date on which reservation is upgraded.
     * @type {Date}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    upsellDate?: Date;
    /**
     * User who upgraded the reservation.
     * @type {string}
     * @memberof UpsellInfoTypeUpsellInfo
     */
    upsellUser?: string;
}

/**
 * Check if a given object implements the UpsellInfoTypeUpsellInfo interface.
 */
export function instanceOfUpsellInfoTypeUpsellInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UpsellInfoTypeUpsellInfoFromJSON(json: any): UpsellInfoTypeUpsellInfo {
    return UpsellInfoTypeUpsellInfoFromJSONTyped(json, false);
}

export function UpsellInfoTypeUpsellInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpsellInfoTypeUpsellInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'totalUpsellCharge': !exists(json, 'totalUpsellCharge') ? undefined : CurrencyAmountTypeFromJSON(json['totalUpsellCharge']),
        'totalUpsellMoneyAmount': !exists(json, 'totalUpsellMoneyAmount') ? undefined : CurrencyAmountTypeFromJSON(json['totalUpsellMoneyAmount']),
        'firstNightUpsellAmount': !exists(json, 'firstNightUpsellAmount') ? undefined : CurrencyAmountTypeFromJSON(json['firstNightUpsellAmount']),
        'firstNightUpsellCharge': !exists(json, 'firstNightUpsellCharge') ? undefined : CurrencyAmountTypeFromJSON(json['firstNightUpsellCharge']),
        'totalUpsellPoints': !exists(json, 'totalUpsellPoints') ? undefined : json['totalUpsellPoints'],
        'totalActualRateAmount': !exists(json, 'totalActualRateAmount') ? undefined : CurrencyAmountTypeFromJSON(json['totalActualRateAmount']),
        'roomType': !exists(json, 'roomType') ? undefined : CodeDescriptionTypeFromJSON(json['roomType']),
        'roomLongDescription': !exists(json, 'roomLongDescription') ? undefined : json['roomLongDescription'],
        'ruleId': !exists(json, 'ruleId') ? undefined : json['ruleId'],
        'ruleCode': !exists(json, 'ruleCode') ? undefined : json['ruleCode'],
        'ruleDescription': !exists(json, 'ruleDescription') ? undefined : json['ruleDescription'],
        'percentageSavings': !exists(json, 'percentageSavings') ? undefined : json['percentageSavings'],
        'upsellDate': !exists(json, 'upsellDate') ? undefined : (new Date(json['upsellDate'])),
        'upsellUser': !exists(json, 'upsellUser') ? undefined : json['upsellUser'],
    };
}

export function UpsellInfoTypeUpsellInfoToJSON(value?: UpsellInfoTypeUpsellInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'totalUpsellCharge': CurrencyAmountTypeToJSON(value.totalUpsellCharge),
        'totalUpsellMoneyAmount': CurrencyAmountTypeToJSON(value.totalUpsellMoneyAmount),
        'firstNightUpsellAmount': CurrencyAmountTypeToJSON(value.firstNightUpsellAmount),
        'firstNightUpsellCharge': CurrencyAmountTypeToJSON(value.firstNightUpsellCharge),
        'totalUpsellPoints': value.totalUpsellPoints,
        'totalActualRateAmount': CurrencyAmountTypeToJSON(value.totalActualRateAmount),
        'roomType': CodeDescriptionTypeToJSON(value.roomType),
        'roomLongDescription': value.roomLongDescription,
        'ruleId': value.ruleId,
        'ruleCode': value.ruleCode,
        'ruleDescription': value.ruleDescription,
        'percentageSavings': value.percentageSavings,
        'upsellDate': value.upsellDate === undefined ? undefined : (value.upsellDate.toISOString().substr(0,10)),
        'upsellUser': value.upsellUser,
    };
}

