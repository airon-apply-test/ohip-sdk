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
 * The Membership object identifies the frequent customer reward program.
 * @export
 * @interface MembershipInfoType
 */
export interface MembershipInfoType {
    /**
     * Membership ID (Unique ID from the memberships table).
     * @type {number}
     * @memberof MembershipInfoType
     */
    membershipId?: number;
    /**
     * The code or name of the membership program ('Hertz', 'AAdvantage', etc.).
     * @type {string}
     * @memberof MembershipInfoType
     */
    programCode?: string;
    /**
     * The code or name of the bonus program. BonusCode can be used to indicate the level of membership (Gold Club, Platinum member, etc.)
     * @type {string}
     * @memberof MembershipInfoType
     */
    bonusCode?: string;
    /**
     * The description of the ProgramCode.(Delta Previlige for code DP)
     * @type {string}
     * @memberof MembershipInfoType
     */
    membershipTypeDesc?: string;
    /**
     * The description of the Bonus Code.(Platinum for code P)
     * @type {string}
     * @memberof MembershipInfoType
     */
    membershipLevelDesc?: string;
    /**
     * The account identification number for this particular member in this particular program.
     * @type {string}
     * @memberof MembershipInfoType
     */
    accountId?: string;
    /**
     * The code or name of the membership level and indicates the level of membership (Gold Club, Platinum member, etc.). This is same as the BonusCode.
     * @type {string}
     * @memberof MembershipInfoType
     */
    membershipLevel?: string;
    /**
     * Ranking assigned to the Player Profile by the Gaming system.
     * @type {number}
     * @memberof MembershipInfoType
     */
    playerRanking?: number;
}

/**
 * Check if a given object implements the MembershipInfoType interface.
 */
export function instanceOfMembershipInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipInfoTypeFromJSON(json: any): MembershipInfoType {
    return MembershipInfoTypeFromJSONTyped(json, false);
}

export function MembershipInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipId': !exists(json, 'membershipId') ? undefined : json['membershipId'],
        'programCode': !exists(json, 'programCode') ? undefined : json['programCode'],
        'bonusCode': !exists(json, 'bonusCode') ? undefined : json['bonusCode'],
        'membershipTypeDesc': !exists(json, 'membershipTypeDesc') ? undefined : json['membershipTypeDesc'],
        'membershipLevelDesc': !exists(json, 'membershipLevelDesc') ? undefined : json['membershipLevelDesc'],
        'accountId': !exists(json, 'accountId') ? undefined : json['accountId'],
        'membershipLevel': !exists(json, 'membershipLevel') ? undefined : json['membershipLevel'],
        'playerRanking': !exists(json, 'playerRanking') ? undefined : json['playerRanking'],
    };
}

export function MembershipInfoTypeToJSON(value?: MembershipInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipId': value.membershipId,
        'programCode': value.programCode,
        'bonusCode': value.bonusCode,
        'membershipTypeDesc': value.membershipTypeDesc,
        'membershipLevelDesc': value.membershipLevelDesc,
        'accountId': value.accountId,
        'membershipLevel': value.membershipLevel,
        'playerRanking': value.playerRanking,
    };
}

