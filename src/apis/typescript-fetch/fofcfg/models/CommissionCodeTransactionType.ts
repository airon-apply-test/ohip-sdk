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
import type { CommissionPaidDetailsType } from './CommissionPaidDetailsType';
import {
    CommissionPaidDetailsTypeFromJSON,
    CommissionPaidDetailsTypeFromJSONTyped,
    CommissionPaidDetailsTypeToJSON,
} from './CommissionPaidDetailsType';

/**
 * Revenue based commission details.
 * @export
 * @interface CommissionCodeTransactionType
 */
export interface CommissionCodeTransactionType {
    /**
     * 
     * @type {string}
     * @memberof CommissionCodeTransactionType
     */
    transactionCode?: string;
    /**
     * 
     * @type {string}
     * @memberof CommissionCodeTransactionType
     */
    transactionCodeDesc?: string;
    /**
     * 
     * @type {boolean}
     * @memberof CommissionCodeTransactionType
     */
    basedOnNetAmount?: boolean;
    /**
     * 
     * @type {CommissionPaidDetailsType}
     * @memberof CommissionCodeTransactionType
     */
    commissionPaidDetails?: CommissionPaidDetailsType;
}

/**
 * Check if a given object implements the CommissionCodeTransactionType interface.
 */
export function instanceOfCommissionCodeTransactionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CommissionCodeTransactionTypeFromJSON(json: any): CommissionCodeTransactionType {
    return CommissionCodeTransactionTypeFromJSONTyped(json, false);
}

export function CommissionCodeTransactionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommissionCodeTransactionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'transactionCode': !exists(json, 'transactionCode') ? undefined : json['transactionCode'],
        'transactionCodeDesc': !exists(json, 'transactionCodeDesc') ? undefined : json['transactionCodeDesc'],
        'basedOnNetAmount': !exists(json, 'basedOnNetAmount') ? undefined : json['basedOnNetAmount'],
        'commissionPaidDetails': !exists(json, 'commissionPaidDetails') ? undefined : CommissionPaidDetailsTypeFromJSON(json['commissionPaidDetails']),
    };
}

export function CommissionCodeTransactionTypeToJSON(value?: CommissionCodeTransactionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'transactionCode': value.transactionCode,
        'transactionCodeDesc': value.transactionCodeDesc,
        'basedOnNetAmount': value.basedOnNetAmount,
        'commissionPaidDetails': CommissionPaidDetailsTypeToJSON(value.commissionPaidDetails),
    };
}

