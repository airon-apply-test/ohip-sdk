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
import type { RateByAgeBucketType } from './RateByAgeBucketType';
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
export declare function instanceOfEffectiveRateType(value: object): boolean;
export declare function EffectiveRateTypeFromJSON(json: any): EffectiveRateType;
export declare function EffectiveRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EffectiveRateType;
export declare function EffectiveRateTypeToJSON(value?: EffectiveRateType | null): any;
