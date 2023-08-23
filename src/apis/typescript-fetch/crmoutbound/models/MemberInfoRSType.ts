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
import type { MemberInfoDetailsType } from './MemberInfoDetailsType';
import {
    MemberInfoDetailsTypeFromJSON,
    MemberInfoDetailsTypeFromJSONTyped,
    MemberInfoDetailsTypeToJSON,
} from './MemberInfoDetailsType';

/**
 * 
 * @export
 * @interface MemberInfoRSType
 */
export interface MemberInfoRSType {
    /**
     * 
     * @type {MemberInfoDetailsType}
     * @memberof MemberInfoRSType
     */
    memberInfoDetails?: MemberInfoDetailsType;
    /**
     * 
     * @type {MemberInfoDetailsType}
     * @memberof MemberInfoRSType
     */
    memberInfoRangeDetails?: MemberInfoDetailsType;
    /**
     * Points required for next tier of membership.
     * @type {number}
     * @memberof MemberInfoRSType
     */
    nextTierPoints?: number;
    /**
     * Points available for the member.
     * @type {number}
     * @memberof MemberInfoRSType
     */
    pointsAvailable?: number;
}

/**
 * Check if a given object implements the MemberInfoRSType interface.
 */
export function instanceOfMemberInfoRSType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MemberInfoRSTypeFromJSON(json: any): MemberInfoRSType {
    return MemberInfoRSTypeFromJSONTyped(json, false);
}

export function MemberInfoRSTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MemberInfoRSType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'memberInfoDetails': !exists(json, 'memberInfoDetails') ? undefined : MemberInfoDetailsTypeFromJSON(json['memberInfoDetails']),
        'memberInfoRangeDetails': !exists(json, 'memberInfoRangeDetails') ? undefined : MemberInfoDetailsTypeFromJSON(json['memberInfoRangeDetails']),
        'nextTierPoints': !exists(json, 'nextTierPoints') ? undefined : json['nextTierPoints'],
        'pointsAvailable': !exists(json, 'pointsAvailable') ? undefined : json['pointsAvailable'],
    };
}

export function MemberInfoRSTypeToJSON(value?: MemberInfoRSType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'memberInfoDetails': MemberInfoDetailsTypeToJSON(value.memberInfoDetails),
        'memberInfoRangeDetails': MemberInfoDetailsTypeToJSON(value.memberInfoRangeDetails),
        'nextTierPoints': value.nextTierPoints,
        'pointsAvailable': value.pointsAvailable,
    };
}

