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
/**
 * 
 * @export
 * @interface AwardVouchersTypeInner
 */
export interface AwardVouchersTypeInner {
    /**
     * Membership Award code applied on the reservation.
     * @type {string}
     * @memberof AwardVouchersTypeInner
     */
    awardCode?: string;
    /**
     * Membership Award number.
     * @type {string}
     * @memberof AwardVouchersTypeInner
     */
    voucherNo?: string;
}

/**
 * Check if a given object implements the AwardVouchersTypeInner interface.
 */
export function instanceOfAwardVouchersTypeInner(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AwardVouchersTypeInnerFromJSON(json: any): AwardVouchersTypeInner {
    return AwardVouchersTypeInnerFromJSONTyped(json, false);
}

export function AwardVouchersTypeInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): AwardVouchersTypeInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'awardCode': !exists(json, 'awardCode') ? undefined : json['awardCode'],
        'voucherNo': !exists(json, 'voucherNo') ? undefined : json['voucherNo'],
    };
}

export function AwardVouchersTypeInnerToJSON(value?: AwardVouchersTypeInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'awardCode': value.awardCode,
        'voucherNo': value.voucherNo,
    };
}

