/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { AwardVouchersTypeInner } from './AwardVouchersTypeInner';
import {
    AwardVouchersTypeInnerFromJSON,
    AwardVouchersTypeInnerFromJSONTyped,
    AwardVouchersTypeInnerToJSON,
} from './AwardVouchersTypeInner';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Membership Awards code applied on the reservation.
 * @export
 * @interface ResAwardsType
 */
export interface ResAwardsType {
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ResAwardsType
     */
    membershipNo?: UniqueIDType;
    /**
     * This stores the Membership Awards code applied on the reservation.
     * @type {Array<AwardVouchersTypeInner>}
     * @memberof ResAwardsType
     */
    awardVouchers?: Array<AwardVouchersTypeInner>;
    /**
     * Room Type before the Upgrade Award.
     * @type {string}
     * @memberof ResAwardsType
     */
    originalRoomType?: string;
    /**
     * Room Type after the Upgrade Award.
     * @type {string}
     * @memberof ResAwardsType
     */
    upgradeRoomType?: string;
}

/**
 * Check if a given object implements the ResAwardsType interface.
 */
export function instanceOfResAwardsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResAwardsTypeFromJSON(json: any): ResAwardsType {
    return ResAwardsTypeFromJSONTyped(json, false);
}

export function ResAwardsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResAwardsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipNo': !exists(json, 'membershipNo') ? undefined : UniqueIDTypeFromJSON(json['membershipNo']),
        'awardVouchers': !exists(json, 'awardVouchers') ? undefined : ((json['awardVouchers'] as Array<any>).map(AwardVouchersTypeInnerFromJSON)),
        'originalRoomType': !exists(json, 'originalRoomType') ? undefined : json['originalRoomType'],
        'upgradeRoomType': !exists(json, 'upgradeRoomType') ? undefined : json['upgradeRoomType'],
    };
}

export function ResAwardsTypeToJSON(value?: ResAwardsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipNo': UniqueIDTypeToJSON(value.membershipNo),
        'awardVouchers': value.awardVouchers === undefined ? undefined : ((value.awardVouchers as Array<any>).map(AwardVouchersTypeInnerToJSON)),
        'originalRoomType': value.originalRoomType,
        'upgradeRoomType': value.upgradeRoomType,
    };
}

