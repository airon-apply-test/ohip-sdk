/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { DiscountType } from './DiscountType';
import {
    DiscountTypeFromJSON,
    DiscountTypeFromJSONTyped,
    DiscountTypeToJSON,
} from './DiscountType';
import type { PointsType } from './PointsType';
import {
    PointsTypeFromJSON,
    PointsTypeFromJSONTyped,
    PointsTypeToJSON,
} from './PointsType';
import type { ShareDistributionInstructionType } from './ShareDistributionInstructionType';
import {
    ShareDistributionInstructionTypeFromJSON,
    ShareDistributionInstructionTypeFromJSONTyped,
    ShareDistributionInstructionTypeToJSON,
} from './ShareDistributionInstructionType';
import type { TotalType } from './TotalType';
import {
    TotalTypeFromJSON,
    TotalTypeFromJSONTyped,
    TotalTypeToJSON,
} from './TotalType';

/**
 * Base charge and additional charges related to a room that includes such things as additional guest amounts, cancel fees, etc. Also includes Discount percentages, total amount, and the rate description.
 * @export
 * @interface AmountType
 */
export interface AmountType {
    /**
     * 
     * @type {TotalType}
     * @memberof AmountType
     */
    base?: TotalType;
    /**
     * 
     * @type {DiscountType}
     * @memberof AmountType
     */
    discount?: DiscountType;
    /**
     * Indicates the share rate percentage for the reservation if set to CUSTOMSPLIT.
     * @type {number}
     * @memberof AmountType
     */
    shareRatePercentage?: number;
    /**
     * 
     * @type {ShareDistributionInstructionType}
     * @memberof AmountType
     */
    shareDistributionInstruction?: ShareDistributionInstructionType;
    /**
     * 
     * @type {TotalType}
     * @memberof AmountType
     */
    total?: TotalType;
    /**
     * 
     * @type {PointsType}
     * @memberof AmountType
     */
    requiredPoints?: PointsType;
    /**
     * 
     * @type {TotalType}
     * @memberof AmountType
     */
    effectiveRate?: TotalType;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof AmountType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof AmountType
     */
    end?: Date;
}

/**
 * Check if a given object implements the AmountType interface.
 */
export function instanceOfAmountType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AmountTypeFromJSON(json: any): AmountType {
    return AmountTypeFromJSONTyped(json, false);
}

export function AmountTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AmountType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'base': !exists(json, 'base') ? undefined : TotalTypeFromJSON(json['base']),
        'discount': !exists(json, 'discount') ? undefined : DiscountTypeFromJSON(json['discount']),
        'shareRatePercentage': !exists(json, 'shareRatePercentage') ? undefined : json['shareRatePercentage'],
        'shareDistributionInstruction': !exists(json, 'shareDistributionInstruction') ? undefined : ShareDistributionInstructionTypeFromJSON(json['shareDistributionInstruction']),
        'total': !exists(json, 'total') ? undefined : TotalTypeFromJSON(json['total']),
        'requiredPoints': !exists(json, 'requiredPoints') ? undefined : PointsTypeFromJSON(json['requiredPoints']),
        'effectiveRate': !exists(json, 'effectiveRate') ? undefined : TotalTypeFromJSON(json['effectiveRate']),
        'start': !exists(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !exists(json, 'end') ? undefined : (new Date(json['end'])),
    };
}

export function AmountTypeToJSON(value?: AmountType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'base': TotalTypeToJSON(value.base),
        'discount': DiscountTypeToJSON(value.discount),
        'shareRatePercentage': value.shareRatePercentage,
        'shareDistributionInstruction': ShareDistributionInstructionTypeToJSON(value.shareDistributionInstruction),
        'total': TotalTypeToJSON(value.total),
        'requiredPoints': PointsTypeToJSON(value.requiredPoints),
        'effectiveRate': TotalTypeToJSON(value.effectiveRate),
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0,10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0,10)),
    };
}

