/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * Information related to a member point to expire.
 * @export
 * @interface MemberPointToExpireType
 */
export interface MemberPointToExpireType {
    /**
     * The date when the member future award points will expire.
     * @type {Date}
     * @memberof MemberPointToExpireType
     */
    pointsExpiryDate?: Date;
    /**
     * Displays the total number of points that will expire on the expiration date.
     * @type {number}
     * @memberof MemberPointToExpireType
     */
    pointsToExpire?: number;
    /**
     * Displays the number of points that are relevant for extension. Points that were already extended from the previous year are not considered for extension.
     * @type {number}
     * @memberof MemberPointToExpireType
     */
    previousPointsToExpire?: number;
    /**
     * Indicates if the points is extendable.
     * @type {boolean}
     * @memberof MemberPointToExpireType
     */
    extend?: boolean;
    /**
     * The method of award generation.
     * @type {string}
     * @memberof MemberPointToExpireType
     */
    awardGenerationMethod?: string;
}

/**
 * Check if a given object implements the MemberPointToExpireType interface.
 */
export function instanceOfMemberPointToExpireType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MemberPointToExpireTypeFromJSON(json: any): MemberPointToExpireType {
    return MemberPointToExpireTypeFromJSONTyped(json, false);
}

export function MemberPointToExpireTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MemberPointToExpireType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pointsExpiryDate': !exists(json, 'pointsExpiryDate') ? undefined : (new Date(json['pointsExpiryDate'])),
        'pointsToExpire': !exists(json, 'pointsToExpire') ? undefined : json['pointsToExpire'],
        'previousPointsToExpire': !exists(json, 'previousPointsToExpire') ? undefined : json['previousPointsToExpire'],
        'extend': !exists(json, 'extend') ? undefined : json['extend'],
        'awardGenerationMethod': !exists(json, 'awardGenerationMethod') ? undefined : json['awardGenerationMethod'],
    };
}

export function MemberPointToExpireTypeToJSON(value?: MemberPointToExpireType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pointsExpiryDate': value.pointsExpiryDate === undefined ? undefined : (value.pointsExpiryDate.toISOString().substr(0,10)),
        'pointsToExpire': value.pointsToExpire,
        'previousPointsToExpire': value.previousPointsToExpire,
        'extend': value.extend,
        'awardGenerationMethod': value.awardGenerationMethod,
    };
}

