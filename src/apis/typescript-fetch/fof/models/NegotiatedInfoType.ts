/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * This provides information for a profile negotiated rate.
 * @export
 * @interface NegotiatedInfoType
 */
export interface NegotiatedInfoType {
    /**
     * The master identifier for multiple offices/locations under the same company profile. This is optional
     * @type {string}
     * @memberof NegotiatedInfoType
     */
    corporateAgreementId?: string;
    /**
     * Informational purposes only in numeric format.
     * @type {string}
     * @memberof NegotiatedInfoType
     */
    comissionCode?: string;
    /**
     * The sell order.
     * @type {number}
     * @memberof NegotiatedInfoType
     */
    order?: number;
    /**
     * Negotiated Rate is inactive or not
     * @type {boolean}
     * @memberof NegotiatedInfoType
     */
    inactive?: boolean;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof NegotiatedInfoType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof NegotiatedInfoType
     */
    end?: Date;
}

/**
 * Check if a given object implements the NegotiatedInfoType interface.
 */
export function instanceOfNegotiatedInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function NegotiatedInfoTypeFromJSON(json: any): NegotiatedInfoType {
    return NegotiatedInfoTypeFromJSONTyped(json, false);
}

export function NegotiatedInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): NegotiatedInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'corporateAgreementId': !exists(json, 'corporateAgreementId') ? undefined : json['corporateAgreementId'],
        'comissionCode': !exists(json, 'comissionCode') ? undefined : json['comissionCode'],
        'order': !exists(json, 'order') ? undefined : json['order'],
        'inactive': !exists(json, 'inactive') ? undefined : json['inactive'],
        'start': !exists(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !exists(json, 'end') ? undefined : (new Date(json['end'])),
    };
}

export function NegotiatedInfoTypeToJSON(value?: NegotiatedInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'corporateAgreementId': value.corporateAgreementId,
        'comissionCode': value.comissionCode,
        'order': value.order,
        'inactive': value.inactive,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0,10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0,10)),
    };
}

