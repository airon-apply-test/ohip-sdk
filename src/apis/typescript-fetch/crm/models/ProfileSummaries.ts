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
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { MasterInfoType } from './MasterInfoType';
import {
    MasterInfoTypeFromJSON,
    MasterInfoTypeFromJSONTyped,
    MasterInfoTypeToJSON,
} from './MasterInfoType';
import type { ProfileSummariesType } from './ProfileSummariesType';
import {
    ProfileSummariesTypeFromJSON,
    ProfileSummariesTypeFromJSONTyped,
    ProfileSummariesTypeToJSON,
} from './ProfileSummariesType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Response object for fetching profiles. This object contains profile primary details,profile summary details,master info list containing code and description used as lookup for description.
 * @export
 * @interface ProfileSummaries
 */
export interface ProfileSummaries {
    /**
     * 
     * @type {ProfileSummariesType}
     * @memberof ProfileSummaries
     */
    profileSummaries?: ProfileSummariesType;
    /**
     * Refer to Generic common types document.
     * @type {Array<MasterInfoType>}
     * @memberof ProfileSummaries
     */
    masterInfoList?: Array<MasterInfoType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ProfileSummaries
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ProfileSummaries
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ProfileSummaries interface.
 */
export function instanceOfProfileSummaries(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileSummariesFromJSON(json: any): ProfileSummaries {
    return ProfileSummariesFromJSONTyped(json, false);
}

export function ProfileSummariesFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileSummaries {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileSummaries': !exists(json, 'profileSummaries') ? undefined : ProfileSummariesTypeFromJSON(json['profileSummaries']),
        'masterInfoList': !exists(json, 'masterInfoList') ? undefined : ((json['masterInfoList'] as Array<any>).map(MasterInfoTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ProfileSummariesToJSON(value?: ProfileSummaries | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileSummaries': ProfileSummariesTypeToJSON(value.profileSummaries),
        'masterInfoList': value.masterInfoList === undefined ? undefined : ((value.masterInfoList as Array<any>).map(MasterInfoTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

