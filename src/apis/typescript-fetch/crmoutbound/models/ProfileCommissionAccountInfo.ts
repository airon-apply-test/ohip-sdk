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
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { ProfileCommissionAccountInfoType } from './ProfileCommissionAccountInfoType';
import {
    ProfileCommissionAccountInfoTypeFromJSON,
    ProfileCommissionAccountInfoTypeFromJSONTyped,
    ProfileCommissionAccountInfoTypeToJSON,
} from './ProfileCommissionAccountInfoType';

/**
 * Response object for fetching profile commission detail.
 * @export
 * @interface ProfileCommissionAccountInfo
 */
export interface ProfileCommissionAccountInfo {
    /**
     * Profile commission info which contains bank account and commission code details.
     * @type {Array<ProfileCommissionAccountInfoType>}
     * @memberof ProfileCommissionAccountInfo
     */
    profileCommissionAccountInfoList?: Array<ProfileCommissionAccountInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ProfileCommissionAccountInfo
     */
    links?: Array<InstanceLink>;
}

/**
 * Check if a given object implements the ProfileCommissionAccountInfo interface.
 */
export function instanceOfProfileCommissionAccountInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileCommissionAccountInfoFromJSON(json: any): ProfileCommissionAccountInfo {
    return ProfileCommissionAccountInfoFromJSONTyped(json, false);
}

export function ProfileCommissionAccountInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileCommissionAccountInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileCommissionAccountInfoList': !exists(json, 'profileCommissionAccountInfoList') ? undefined : ((json['profileCommissionAccountInfoList'] as Array<any>).map(ProfileCommissionAccountInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
    };
}

export function ProfileCommissionAccountInfoToJSON(value?: ProfileCommissionAccountInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileCommissionAccountInfoList': value.profileCommissionAccountInfoList === undefined ? undefined : ((value.profileCommissionAccountInfoList as Array<any>).map(ProfileCommissionAccountInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
    };
}

