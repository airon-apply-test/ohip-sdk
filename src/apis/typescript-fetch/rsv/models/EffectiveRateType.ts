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
import type { RateByAgeBucketType } from './RateByAgeBucketType';
import {
    RateByAgeBucketTypeFromJSON,
    RateByAgeBucketTypeFromJSONTyped,
    RateByAgeBucketTypeToJSON,
} from './RateByAgeBucketType';

/**
 * Effective rate amount per guest on specific dates.
 * @export
 * @interface EffectiveRateType
 */
export interface EffectiveRateType {
    /**
     * Rate amount for one person.
     * @type {number}
     * @memberof EffectiveRateType
     */
    onePersonRate?: number;
    /**
     * Rate amount for two persons.
     * @type {number}
     * @memberof EffectiveRateType
     */
    twoPersonRate?: number;
    /**
     * Rate amount for three persons.
     * @type {number}
     * @memberof EffectiveRateType
     */
    threePersonRate?: number;
    /**
     * Rate amount for four persons.
     * @type {number}
     * @memberof EffectiveRateType
     */
    fourPersonRate?: number;
    /**
     * Rate amount for five persons.
     * @type {number}
     * @memberof EffectiveRateType
     */
    fivePersonRate?: number;
    /**
     * Rate amount for each extra person.
     * @type {number}
     * @memberof EffectiveRateType
     */
    extraPersonRate?: number;
    /**
     * Rate amount for each extra Child.
     * @type {number}
     * @memberof EffectiveRateType
     */
    extraChildRate?: number;
    /**
     * Rate amount for one Child.
     * @type {number}
     * @memberof EffectiveRateType
     */
    oneChildRate?: number;
    /**
     * Rate amount for two Children.
     * @type {number}
     * @memberof EffectiveRateType
     */
    twoChildrenRate?: number;
    /**
     * Rate amount for three Children.
     * @type {number}
     * @memberof EffectiveRateType
     */
    threeChildrenRate?: number;
    /**
     * Rate amount for four Children.
     * @type {number}
     * @memberof EffectiveRateType
     */
    fourChildrenRate?: number;
    /**
     * Rate amount by age bucket.
     * @type {Array<RateByAgeBucketType>}
     * @memberof EffectiveRateType
     */
    rateByAgeBuckets?: Array<RateByAgeBucketType>;
    /**
     * Minimum number of children needed to get free stay.
     * @type {number}
     * @memberof EffectiveRateType
     */
    minimumChildrenForFreeStay?: number;
    /**
     * The number of award points required for applying this rate plan schedule.
     * @type {number}
     * @memberof EffectiveRateType
     */
    pointsRequired?: number;
    /**
     * true if floor amount needs to be override
     * @type {boolean}
     * @memberof EffectiveRateType
     */
    overrideFloorAmount?: boolean;
    /**
     * 
     * @type {string}
     * @memberof EffectiveRateType
     */
    hotelId?: string;
    /**
     * 
     * @type {string}
     * @memberof EffectiveRateType
     */
    ratePlanCode?: string;
    /**
     * 
     * @type {string}
     * @memberof EffectiveRateType
     */
    roomType?: string;
    /**
     * 
     * @type {number}
     * @memberof EffectiveRateType
     */
    numberOfRooms?: number;
    /**
     * Rate amount Before Tax.
     * @type {number}
     * @memberof EffectiveRateType
     */
    amountBeforeTax?: number;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof EffectiveRateType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof EffectiveRateType
     */
    end?: Date;
}

/**
 * Check if a given object implements the EffectiveRateType interface.
 */
export function instanceOfEffectiveRateType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function EffectiveRateTypeFromJSON(json: any): EffectiveRateType {
    return EffectiveRateTypeFromJSONTyped(json, false);
}

export function EffectiveRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EffectiveRateType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'onePersonRate': !exists(json, 'onePersonRate') ? undefined : json['onePersonRate'],
        'twoPersonRate': !exists(json, 'twoPersonRate') ? undefined : json['twoPersonRate'],
        'threePersonRate': !exists(json, 'threePersonRate') ? undefined : json['threePersonRate'],
        'fourPersonRate': !exists(json, 'fourPersonRate') ? undefined : json['fourPersonRate'],
        'fivePersonRate': !exists(json, 'fivePersonRate') ? undefined : json['fivePersonRate'],
        'extraPersonRate': !exists(json, 'extraPersonRate') ? undefined : json['extraPersonRate'],
        'extraChildRate': !exists(json, 'extraChildRate') ? undefined : json['extraChildRate'],
        'oneChildRate': !exists(json, 'oneChildRate') ? undefined : json['oneChildRate'],
        'twoChildrenRate': !exists(json, 'twoChildrenRate') ? undefined : json['twoChildrenRate'],
        'threeChildrenRate': !exists(json, 'threeChildrenRate') ? undefined : json['threeChildrenRate'],
        'fourChildrenRate': !exists(json, 'fourChildrenRate') ? undefined : json['fourChildrenRate'],
        'rateByAgeBuckets': !exists(json, 'rateByAgeBuckets') ? undefined : ((json['rateByAgeBuckets'] as Array<any>).map(RateByAgeBucketTypeFromJSON)),
        'minimumChildrenForFreeStay': !exists(json, 'minimumChildrenForFreeStay') ? undefined : json['minimumChildrenForFreeStay'],
        'pointsRequired': !exists(json, 'pointsRequired') ? undefined : json['pointsRequired'],
        'overrideFloorAmount': !exists(json, 'overrideFloorAmount') ? undefined : json['overrideFloorAmount'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'ratePlanCode': !exists(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'roomType': !exists(json, 'roomType') ? undefined : json['roomType'],
        'numberOfRooms': !exists(json, 'numberOfRooms') ? undefined : json['numberOfRooms'],
        'amountBeforeTax': !exists(json, 'amountBeforeTax') ? undefined : json['amountBeforeTax'],
        'start': !exists(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !exists(json, 'end') ? undefined : (new Date(json['end'])),
    };
}

export function EffectiveRateTypeToJSON(value?: EffectiveRateType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'onePersonRate': value.onePersonRate,
        'twoPersonRate': value.twoPersonRate,
        'threePersonRate': value.threePersonRate,
        'fourPersonRate': value.fourPersonRate,
        'fivePersonRate': value.fivePersonRate,
        'extraPersonRate': value.extraPersonRate,
        'extraChildRate': value.extraChildRate,
        'oneChildRate': value.oneChildRate,
        'twoChildrenRate': value.twoChildrenRate,
        'threeChildrenRate': value.threeChildrenRate,
        'fourChildrenRate': value.fourChildrenRate,
        'rateByAgeBuckets': value.rateByAgeBuckets === undefined ? undefined : ((value.rateByAgeBuckets as Array<any>).map(RateByAgeBucketTypeToJSON)),
        'minimumChildrenForFreeStay': value.minimumChildrenForFreeStay,
        'pointsRequired': value.pointsRequired,
        'overrideFloorAmount': value.overrideFloorAmount,
        'hotelId': value.hotelId,
        'ratePlanCode': value.ratePlanCode,
        'roomType': value.roomType,
        'numberOfRooms': value.numberOfRooms,
        'amountBeforeTax': value.amountBeforeTax,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0,10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0,10)),
    };
}

