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
 * Configured Billing Instruction which represents a set of Transaction Codes.
 * @export
 * @interface BillingInstructionType
 */
export interface BillingInstructionType {
    /**
     * Billing Instruction code description.
     * @type {string}
     * @memberof BillingInstructionType
     */
    desc?: string;
    /**
     * This is the Routing Instruction Id attached with Reservation. It is only used for internal purpose. It should not be used by external vendor or consumer.
     * @type {number}
     * @memberof BillingInstructionType
     */
    routingInstructionsId?: number;
    /**
     * Unique identifier for the Billing Instruction.
     * @type {string}
     * @memberof BillingInstructionType
     */
    billingCode?: string;
    /**
     * Hotel context of the Billing Instruction.
     * @type {string}
     * @memberof BillingInstructionType
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the BillingInstructionType interface.
 */
export function instanceOfBillingInstructionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BillingInstructionTypeFromJSON(json: any): BillingInstructionType {
    return BillingInstructionTypeFromJSONTyped(json, false);
}

export function BillingInstructionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BillingInstructionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'desc': !exists(json, 'desc') ? undefined : json['desc'],
        'routingInstructionsId': !exists(json, 'routingInstructionsId') ? undefined : json['routingInstructionsId'],
        'billingCode': !exists(json, 'billingCode') ? undefined : json['billingCode'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function BillingInstructionTypeToJSON(value?: BillingInstructionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'desc': value.desc,
        'routingInstructionsId': value.routingInstructionsId,
        'billingCode': value.billingCode,
        'hotelId': value.hotelId,
    };
}

