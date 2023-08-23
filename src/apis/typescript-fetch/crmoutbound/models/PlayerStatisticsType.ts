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
 * Contains Player Statistics information
 * @export
 * @interface PlayerStatisticsType
 */
export interface PlayerStatisticsType {
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    enrollmentDate?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    compDollars?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    actualTableWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    actualOtherWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    tableTimePlayed?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    otherTimePlayed?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    averageSlotBet?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    skillRating?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    playerTableWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    playerOtherWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    theoriticalTableWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    theoriticalOtherWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    expectedProfit?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    grossMarkers?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    compPoints?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    totalComps?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    actualSlotWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    creditLimit?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    slotTimePlayed?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    avgTableBet?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    avgOtherBet?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    preferredGame?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    playerSlotWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    noOfRatings?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    theoSlotWins?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    theoProfit?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    bettingLimit?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    netMarkers?: string;
    /**
     * 
     * @type {string}
     * @memberof PlayerStatisticsType
     */
    notes?: string;
}

/**
 * Check if a given object implements the PlayerStatisticsType interface.
 */
export function instanceOfPlayerStatisticsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PlayerStatisticsTypeFromJSON(json: any): PlayerStatisticsType {
    return PlayerStatisticsTypeFromJSONTyped(json, false);
}

export function PlayerStatisticsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PlayerStatisticsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'enrollmentDate': !exists(json, 'enrollmentDate') ? undefined : json['enrollmentDate'],
        'compDollars': !exists(json, 'compDollars') ? undefined : json['compDollars'],
        'actualTableWins': !exists(json, 'actualTableWins') ? undefined : json['actualTableWins'],
        'actualOtherWins': !exists(json, 'actualOtherWins') ? undefined : json['actualOtherWins'],
        'tableTimePlayed': !exists(json, 'tableTimePlayed') ? undefined : json['tableTimePlayed'],
        'otherTimePlayed': !exists(json, 'otherTimePlayed') ? undefined : json['otherTimePlayed'],
        'averageSlotBet': !exists(json, 'averageSlotBet') ? undefined : json['averageSlotBet'],
        'skillRating': !exists(json, 'skillRating') ? undefined : json['skillRating'],
        'playerTableWins': !exists(json, 'playerTableWins') ? undefined : json['playerTableWins'],
        'playerOtherWins': !exists(json, 'playerOtherWins') ? undefined : json['playerOtherWins'],
        'theoriticalTableWins': !exists(json, 'theoriticalTableWins') ? undefined : json['theoriticalTableWins'],
        'theoriticalOtherWins': !exists(json, 'theoriticalOtherWins') ? undefined : json['theoriticalOtherWins'],
        'expectedProfit': !exists(json, 'expectedProfit') ? undefined : json['expectedProfit'],
        'grossMarkers': !exists(json, 'grossMarkers') ? undefined : json['grossMarkers'],
        'compPoints': !exists(json, 'compPoints') ? undefined : json['compPoints'],
        'totalComps': !exists(json, 'totalComps') ? undefined : json['totalComps'],
        'actualSlotWins': !exists(json, 'actualSlotWins') ? undefined : json['actualSlotWins'],
        'creditLimit': !exists(json, 'creditLimit') ? undefined : json['creditLimit'],
        'slotTimePlayed': !exists(json, 'slotTimePlayed') ? undefined : json['slotTimePlayed'],
        'avgTableBet': !exists(json, 'avgTableBet') ? undefined : json['avgTableBet'],
        'avgOtherBet': !exists(json, 'avgOtherBet') ? undefined : json['avgOtherBet'],
        'preferredGame': !exists(json, 'preferredGame') ? undefined : json['preferredGame'],
        'playerSlotWins': !exists(json, 'playerSlotWins') ? undefined : json['playerSlotWins'],
        'noOfRatings': !exists(json, 'noOfRatings') ? undefined : json['noOfRatings'],
        'theoSlotWins': !exists(json, 'theoSlotWins') ? undefined : json['theoSlotWins'],
        'theoProfit': !exists(json, 'theoProfit') ? undefined : json['theoProfit'],
        'bettingLimit': !exists(json, 'bettingLimit') ? undefined : json['bettingLimit'],
        'netMarkers': !exists(json, 'netMarkers') ? undefined : json['netMarkers'],
        'notes': !exists(json, 'notes') ? undefined : json['notes'],
    };
}

export function PlayerStatisticsTypeToJSON(value?: PlayerStatisticsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'enrollmentDate': value.enrollmentDate,
        'compDollars': value.compDollars,
        'actualTableWins': value.actualTableWins,
        'actualOtherWins': value.actualOtherWins,
        'tableTimePlayed': value.tableTimePlayed,
        'otherTimePlayed': value.otherTimePlayed,
        'averageSlotBet': value.averageSlotBet,
        'skillRating': value.skillRating,
        'playerTableWins': value.playerTableWins,
        'playerOtherWins': value.playerOtherWins,
        'theoriticalTableWins': value.theoriticalTableWins,
        'theoriticalOtherWins': value.theoriticalOtherWins,
        'expectedProfit': value.expectedProfit,
        'grossMarkers': value.grossMarkers,
        'compPoints': value.compPoints,
        'totalComps': value.totalComps,
        'actualSlotWins': value.actualSlotWins,
        'creditLimit': value.creditLimit,
        'slotTimePlayed': value.slotTimePlayed,
        'avgTableBet': value.avgTableBet,
        'avgOtherBet': value.avgOtherBet,
        'preferredGame': value.preferredGame,
        'playerSlotWins': value.playerSlotWins,
        'noOfRatings': value.noOfRatings,
        'theoSlotWins': value.theoSlotWins,
        'theoProfit': value.theoProfit,
        'bettingLimit': value.bettingLimit,
        'netMarkers': value.netMarkers,
        'notes': value.notes,
    };
}

