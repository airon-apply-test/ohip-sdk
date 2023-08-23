/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud CRM Configuration API
 * APIs for Customer Relationship Management (profile) configuration, such as creating preferences, or address types.  It also includes Membership Configuration, where you can retrieve membership levels that are configured for a property, or create new membership enrollment codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
import type { ProfileRestrictionReasonType } from './ProfileRestrictionReasonType';
import {
    ProfileRestrictionReasonTypeFromJSON,
    ProfileRestrictionReasonTypeFromJSONTyped,
    ProfileRestrictionReasonTypeToJSON,
} from './ProfileRestrictionReasonType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for changing Profile Restriction Reasons.
 * @export
 * @interface ProfileRestrictionReasonsToBeChanged
 */
export interface ProfileRestrictionReasonsToBeChanged {
    /**
     * List of Profile Restriction Reasons.
     * @type {Array<ProfileRestrictionReasonType>}
     * @memberof ProfileRestrictionReasonsToBeChanged
     */
    profileRestrictionReasons?: Array<ProfileRestrictionReasonType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ProfileRestrictionReasonsToBeChanged
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ProfileRestrictionReasonsToBeChanged
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ProfileRestrictionReasonsToBeChanged interface.
 */
export function instanceOfProfileRestrictionReasonsToBeChanged(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileRestrictionReasonsToBeChangedFromJSON(json: any): ProfileRestrictionReasonsToBeChanged {
    return ProfileRestrictionReasonsToBeChangedFromJSONTyped(json, false);
}

export function ProfileRestrictionReasonsToBeChangedFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileRestrictionReasonsToBeChanged {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileRestrictionReasons': !exists(json, 'profileRestrictionReasons') ? undefined : ((json['profileRestrictionReasons'] as Array<any>).map(ProfileRestrictionReasonTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ProfileRestrictionReasonsToBeChangedToJSON(value?: ProfileRestrictionReasonsToBeChanged | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileRestrictionReasons': value.profileRestrictionReasons === undefined ? undefined : ((value.profileRestrictionReasons as Array<any>).map(ProfileRestrictionReasonTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

