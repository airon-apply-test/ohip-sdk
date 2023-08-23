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
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';
import type { DepositCancelRevenueType } from './DepositCancelRevenueType';
import {
    DepositCancelRevenueTypeFromJSON,
    DepositCancelRevenueTypeFromJSONTyped,
    DepositCancelRevenueTypeToJSON,
} from './DepositCancelRevenueType';
import type { DepositPolicyType } from './DepositPolicyType';
import {
    DepositPolicyTypeFromJSON,
    DepositPolicyTypeFromJSONTyped,
    DepositPolicyTypeToJSON,
} from './DepositPolicyType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * A deposit policy attached with the reservation.
 * @export
 * @interface GuaranteePolicyTypeDepositPolicy
 */
export interface GuaranteePolicyTypeDepositPolicy {
    /**
     * 
     * @type {DepositCancelRevenueType}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    revenueType?: DepositCancelRevenueType;
    /**
     * 
     * @type {DepositPolicyType}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    policy?: DepositPolicyType;
    /**
     * Comments attached with a deposit.
     * @type {string}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    comments?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    amountPaid?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    amountDue?: CurrencyAmountType;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    policyId?: UniqueIDType;
    /**
     * Message.
     * @type {number}
     * @memberof GuaranteePolicyTypeDepositPolicy
     */
    estimatedAmount?: number;
}

/**
 * Check if a given object implements the GuaranteePolicyTypeDepositPolicy interface.
 */
export function instanceOfGuaranteePolicyTypeDepositPolicy(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function GuaranteePolicyTypeDepositPolicyFromJSON(json: any): GuaranteePolicyTypeDepositPolicy {
    return GuaranteePolicyTypeDepositPolicyFromJSONTyped(json, false);
}

export function GuaranteePolicyTypeDepositPolicyFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuaranteePolicyTypeDepositPolicy {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'revenueType': !exists(json, 'revenueType') ? undefined : DepositCancelRevenueTypeFromJSON(json['revenueType']),
        'policy': !exists(json, 'policy') ? undefined : DepositPolicyTypeFromJSON(json['policy']),
        'comments': !exists(json, 'comments') ? undefined : json['comments'],
        'amountPaid': !exists(json, 'amountPaid') ? undefined : CurrencyAmountTypeFromJSON(json['amountPaid']),
        'amountDue': !exists(json, 'amountDue') ? undefined : CurrencyAmountTypeFromJSON(json['amountDue']),
        'policyId': !exists(json, 'policyId') ? undefined : UniqueIDTypeFromJSON(json['policyId']),
        'estimatedAmount': !exists(json, 'estimatedAmount') ? undefined : json['estimatedAmount'],
    };
}

export function GuaranteePolicyTypeDepositPolicyToJSON(value?: GuaranteePolicyTypeDepositPolicy | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'revenueType': DepositCancelRevenueTypeToJSON(value.revenueType),
        'policy': DepositPolicyTypeToJSON(value.policy),
        'comments': value.comments,
        'amountPaid': CurrencyAmountTypeToJSON(value.amountPaid),
        'amountDue': CurrencyAmountTypeToJSON(value.amountDue),
        'policyId': UniqueIDTypeToJSON(value.policyId),
        'estimatedAmount': value.estimatedAmount,
    };
}

