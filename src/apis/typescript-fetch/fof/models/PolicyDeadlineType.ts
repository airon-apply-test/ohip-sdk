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
 * Cancellation deadline, absolute or relative.
 * @export
 * @interface PolicyDeadlineType
 */
export interface PolicyDeadlineType {
    /**
     * Defines the absolute deadline. Either this or the offset attributes may be used.
     * @type {Date}
     * @memberof PolicyDeadlineType
     */
    absoluteDeadline?: Date;
    /**
     * The number of days before arrival that allows cancellation without penalties.
     * @type {number}
     * @memberof PolicyDeadlineType
     */
    offsetFromArrival?: number;
    /**
     * Time on offset day the cancellation penalties applies.
     * @type {Date}
     * @memberof PolicyDeadlineType
     */
    offsetDropTime?: Date;
    /**
     * The number of days after booking deposit must be paid.
     * @type {number}
     * @memberof PolicyDeadlineType
     */
    offsetFromBookingDate?: number;
}

/**
 * Check if a given object implements the PolicyDeadlineType interface.
 */
export function instanceOfPolicyDeadlineType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PolicyDeadlineTypeFromJSON(json: any): PolicyDeadlineType {
    return PolicyDeadlineTypeFromJSONTyped(json, false);
}

export function PolicyDeadlineTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicyDeadlineType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'absoluteDeadline': !exists(json, 'absoluteDeadline') ? undefined : (new Date(json['absoluteDeadline'])),
        'offsetFromArrival': !exists(json, 'offsetFromArrival') ? undefined : json['offsetFromArrival'],
        'offsetDropTime': !exists(json, 'offsetDropTime') ? undefined : (new Date(json['offsetDropTime'])),
        'offsetFromBookingDate': !exists(json, 'offsetFromBookingDate') ? undefined : json['offsetFromBookingDate'],
    };
}

export function PolicyDeadlineTypeToJSON(value?: PolicyDeadlineType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'absoluteDeadline': value.absoluteDeadline === undefined ? undefined : (value.absoluteDeadline.toISOString()),
        'offsetFromArrival': value.offsetFromArrival,
        'offsetDropTime': value.offsetDropTime === undefined ? undefined : (value.offsetDropTime.toISOString()),
        'offsetFromBookingDate': value.offsetFromBookingDate,
    };
}

