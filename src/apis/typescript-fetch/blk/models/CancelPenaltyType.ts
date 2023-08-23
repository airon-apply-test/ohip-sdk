/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { OffsetUnitType } from './OffsetUnitType';
import {
    OffsetUnitTypeFromJSON,
    OffsetUnitTypeFromJSONTyped,
    OffsetUnitTypeToJSON,
} from './OffsetUnitType';
import type { PolicyAmountPercentType } from './PolicyAmountPercentType';
import {
    PolicyAmountPercentTypeFromJSON,
    PolicyAmountPercentTypeFromJSONTyped,
    PolicyAmountPercentTypeToJSON,
} from './PolicyAmountPercentType';
import type { PolicyDeadlineType } from './PolicyDeadlineType';
import {
    PolicyDeadlineTypeFromJSON,
    PolicyDeadlineTypeFromJSONTyped,
    PolicyDeadlineTypeToJSON,
} from './PolicyDeadlineType';

/**
 * The CancelPenalty class defines the cancellation policy of the hotel facility.
 * @export
 * @interface CancelPenaltyType
 */
export interface CancelPenaltyType {
    /**
     * 
     * @type {PolicyDeadlineType}
     * @memberof CancelPenaltyType
     */
    deadline?: PolicyDeadlineType;
    /**
     * 
     * @type {PolicyAmountPercentType}
     * @memberof CancelPenaltyType
     */
    amountPercent?: PolicyAmountPercentType;
    /**
     * Text description of the Penalty in a given language.
     * @type {string}
     * @memberof CancelPenaltyType
     */
    penaltyDescription?: string;
    /**
     * 
     * @type {OffsetUnitType}
     * @memberof CancelPenaltyType
     */
    offsetUnit?: OffsetUnitType;
    /**
     * Formatted Text Rule of the Cancellation Penalty.
     * @type {string}
     * @memberof CancelPenaltyType
     */
    formattedRule?: string;
    /**
     * Policy Code.
     * @type {string}
     * @memberof CancelPenaltyType
     */
    policyCode?: string;
    /**
     * Flag to indicate if the cancellation policy is manual.
     * @type {boolean}
     * @memberof CancelPenaltyType
     */
    manual?: boolean;
    /**
     * Indicates if the amount is refundable if booking is canceled.
     * @type {boolean}
     * @memberof CancelPenaltyType
     */
    nonRefundable?: boolean;
}

/**
 * Check if a given object implements the CancelPenaltyType interface.
 */
export function instanceOfCancelPenaltyType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CancelPenaltyTypeFromJSON(json: any): CancelPenaltyType {
    return CancelPenaltyTypeFromJSONTyped(json, false);
}

export function CancelPenaltyTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CancelPenaltyType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'deadline': !exists(json, 'deadline') ? undefined : PolicyDeadlineTypeFromJSON(json['deadline']),
        'amountPercent': !exists(json, 'amountPercent') ? undefined : PolicyAmountPercentTypeFromJSON(json['amountPercent']),
        'penaltyDescription': !exists(json, 'penaltyDescription') ? undefined : json['penaltyDescription'],
        'offsetUnit': !exists(json, 'offsetUnit') ? undefined : OffsetUnitTypeFromJSON(json['offsetUnit']),
        'formattedRule': !exists(json, 'formattedRule') ? undefined : json['formattedRule'],
        'policyCode': !exists(json, 'policyCode') ? undefined : json['policyCode'],
        'manual': !exists(json, 'manual') ? undefined : json['manual'],
        'nonRefundable': !exists(json, 'nonRefundable') ? undefined : json['nonRefundable'],
    };
}

export function CancelPenaltyTypeToJSON(value?: CancelPenaltyType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'deadline': PolicyDeadlineTypeToJSON(value.deadline),
        'amountPercent': PolicyAmountPercentTypeToJSON(value.amountPercent),
        'penaltyDescription': value.penaltyDescription,
        'offsetUnit': OffsetUnitTypeToJSON(value.offsetUnit),
        'formattedRule': value.formattedRule,
        'policyCode': value.policyCode,
        'manual': value.manual,
        'nonRefundable': value.nonRefundable,
    };
}

