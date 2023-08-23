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


/**
 * This value of the enum is used for Payment methods
 * @export
 */
export const CommissionConfigPaymentMethodType = {
    Check: 'Check',
    CentrallyPaid: 'CentrallyPaid',
    Eft: 'Eft'
} as const;
export type CommissionConfigPaymentMethodType = typeof CommissionConfigPaymentMethodType[keyof typeof CommissionConfigPaymentMethodType];


export function CommissionConfigPaymentMethodTypeFromJSON(json: any): CommissionConfigPaymentMethodType {
    return CommissionConfigPaymentMethodTypeFromJSONTyped(json, false);
}

export function CommissionConfigPaymentMethodTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommissionConfigPaymentMethodType {
    return json as CommissionConfigPaymentMethodType;
}

export function CommissionConfigPaymentMethodTypeToJSON(value?: CommissionConfigPaymentMethodType | null): any {
    return value as any;
}

