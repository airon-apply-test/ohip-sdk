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
import type { ProfileSharedLevelType } from './ProfileSharedLevelType';
import {
    ProfileSharedLevelTypeFromJSON,
    ProfileSharedLevelTypeFromJSONTyped,
    ProfileSharedLevelTypeToJSON,
} from './ProfileSharedLevelType';

/**
 * 
 * @export
 * @interface ProfileAccessType
 */
export interface ProfileAccessType {
    /**
     * Indicates the Chain code where the profile was created
     * @type {string}
     * @memberof ProfileAccessType
     */
    chainCode?: string;
    /**
     * Indicates the CRO code where the profile was created
     * @type {string}
     * @memberof ProfileAccessType
     */
    croCode?: string;
    /**
     * Indicates the Hotel code where the profile was created
     * @type {string}
     * @memberof ProfileAccessType
     */
    hotelId?: string;
    /**
     * 
     * @type {ProfileSharedLevelType}
     * @memberof ProfileAccessType
     */
    sharedLevel?: ProfileSharedLevelType;
}

/**
 * Check if a given object implements the ProfileAccessType interface.
 */
export function instanceOfProfileAccessType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileAccessTypeFromJSON(json: any): ProfileAccessType {
    return ProfileAccessTypeFromJSONTyped(json, false);
}

export function ProfileAccessTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileAccessType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'chainCode': !exists(json, 'chainCode') ? undefined : json['chainCode'],
        'croCode': !exists(json, 'croCode') ? undefined : json['croCode'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'sharedLevel': !exists(json, 'sharedLevel') ? undefined : ProfileSharedLevelTypeFromJSON(json['sharedLevel']),
    };
}

export function ProfileAccessTypeToJSON(value?: ProfileAccessType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'chainCode': value.chainCode,
        'croCode': value.croCode,
        'hotelId': value.hotelId,
        'sharedLevel': ProfileSharedLevelTypeToJSON(value.sharedLevel),
    };
}

