/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Configuration API
 * APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';

/**
 * Commission paid details.
 * @export
 * @interface CommissionPaidDetailsType
 */
export interface CommissionPaidDetailsType {
    /**
     * 
     * @type {boolean}
     * @memberof CommissionPaidDetailsType
     */
    perNight?: boolean;
    /**
     * 
     * @type {boolean}
     * @memberof CommissionPaidDetailsType
     */
    perStay?: boolean;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CommissionPaidDetailsType
     */
    commissionAmount?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CommissionPaidDetailsType
     */
    weekEndCommissionAmount?: CurrencyAmountType;
    /**
     * Indicates the type of commission paid to agent.
     * @type {string}
     * @memberof CommissionPaidDetailsType
     */
    commissionType?: string;
}

/**
 * Check if a given object implements the CommissionPaidDetailsType interface.
 */
export function instanceOfCommissionPaidDetailsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CommissionPaidDetailsTypeFromJSON(json: any): CommissionPaidDetailsType {
    return CommissionPaidDetailsTypeFromJSONTyped(json, false);
}

export function CommissionPaidDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommissionPaidDetailsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'perNight': !exists(json, 'perNight') ? undefined : json['perNight'],
        'perStay': !exists(json, 'perStay') ? undefined : json['perStay'],
        'commissionAmount': !exists(json, 'commissionAmount') ? undefined : CurrencyAmountTypeFromJSON(json['commissionAmount']),
        'weekEndCommissionAmount': !exists(json, 'weekEndCommissionAmount') ? undefined : CurrencyAmountTypeFromJSON(json['weekEndCommissionAmount']),
        'commissionType': !exists(json, 'commissionType') ? undefined : json['commissionType'],
    };
}

export function CommissionPaidDetailsTypeToJSON(value?: CommissionPaidDetailsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'perNight': value.perNight,
        'perStay': value.perStay,
        'commissionAmount': CurrencyAmountTypeToJSON(value.commissionAmount),
        'weekEndCommissionAmount': CurrencyAmountTypeToJSON(value.weekEndCommissionAmount),
        'commissionType': value.commissionType,
    };
}

