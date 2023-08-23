/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Rate API
 * APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Threshold for Children in the room.
 * @export
 * @interface OccupantThresholdPricingTypeChildren
 */
export interface OccupantThresholdPricingTypeChildren {
    /**
     * Threshold value, after it is reached the corresponding amount will be charged.
     * @type {number}
     * @memberof OccupantThresholdPricingTypeChildren
     */
    threshold?: number;
    /**
     * Amount to be charged after the threshold is reached.
     * @type {number}
     * @memberof OccupantThresholdPricingTypeChildren
     */
    amount?: number;
    /**
     * When rates are Defined by Age buckets, should the 1st buckets children count be excluded from threshold pricing.
     * @type {boolean}
     * @memberof OccupantThresholdPricingTypeChildren
     */
    excludeBucket1?: boolean;
    /**
     * When rates are Defined by Age buckets, should the 2nd buckets children count be excluded from threshold pricing.
     * @type {boolean}
     * @memberof OccupantThresholdPricingTypeChildren
     */
    excludeBucket2?: boolean;
    /**
     * When rates are Defined by Age buckets, should the 3rd buckets children count be excluded from threshold pricing.
     * @type {boolean}
     * @memberof OccupantThresholdPricingTypeChildren
     */
    excludeBucket3?: boolean;
}

/**
 * Check if a given object implements the OccupantThresholdPricingTypeChildren interface.
 */
export function instanceOfOccupantThresholdPricingTypeChildren(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function OccupantThresholdPricingTypeChildrenFromJSON(json: any): OccupantThresholdPricingTypeChildren {
    return OccupantThresholdPricingTypeChildrenFromJSONTyped(json, false);
}

export function OccupantThresholdPricingTypeChildrenFromJSONTyped(json: any, ignoreDiscriminator: boolean): OccupantThresholdPricingTypeChildren {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'threshold': !exists(json, 'threshold') ? undefined : json['threshold'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'excludeBucket1': !exists(json, 'excludeBucket1') ? undefined : json['excludeBucket1'],
        'excludeBucket2': !exists(json, 'excludeBucket2') ? undefined : json['excludeBucket2'],
        'excludeBucket3': !exists(json, 'excludeBucket3') ? undefined : json['excludeBucket3'],
    };
}

export function OccupantThresholdPricingTypeChildrenToJSON(value?: OccupantThresholdPricingTypeChildren | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'threshold': value.threshold,
        'amount': value.amount,
        'excludeBucket1': value.excludeBucket1,
        'excludeBucket2': value.excludeBucket2,
        'excludeBucket3': value.excludeBucket3,
    };
}

