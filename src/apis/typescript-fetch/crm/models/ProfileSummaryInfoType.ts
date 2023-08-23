/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management API
 * APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ProfileSummaryType } from './ProfileSummaryType';
import {
    ProfileSummaryTypeFromJSON,
    ProfileSummaryTypeFromJSONTyped,
    ProfileSummaryTypeToJSON,
} from './ProfileSummaryType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Summary information about a profile and the associated Unique IDs.
 * @export
 * @interface ProfileSummaryInfoType
 */
export interface ProfileSummaryInfoType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof ProfileSummaryInfoType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     * 
     * @type {ProfileSummaryType}
     * @memberof ProfileSummaryInfoType
     */
    profile?: ProfileSummaryType;
}

/**
 * Check if a given object implements the ProfileSummaryInfoType interface.
 */
export function instanceOfProfileSummaryInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileSummaryInfoTypeFromJSON(json: any): ProfileSummaryInfoType {
    return ProfileSummaryInfoTypeFromJSONTyped(json, false);
}

export function ProfileSummaryInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileSummaryInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileIdList': !exists(json, 'profileIdList') ? undefined : ((json['profileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'profile': !exists(json, 'profile') ? undefined : ProfileSummaryTypeFromJSON(json['profile']),
    };
}

export function ProfileSummaryInfoTypeToJSON(value?: ProfileSummaryInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileIdList': value.profileIdList === undefined ? undefined : ((value.profileIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'profile': ProfileSummaryTypeToJSON(value.profile),
    };
}

