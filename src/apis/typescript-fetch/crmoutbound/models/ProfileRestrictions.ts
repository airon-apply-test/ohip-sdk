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
 * 
 * @export
 * @interface ProfileRestrictions
 */
export interface ProfileRestrictions {
    /**
     * Restriction reason associated with the current profile.
     * @type {string}
     * @memberof ProfileRestrictions
     */
    reason?: string;
    /**
     * Description of restriction reason.
     * @type {string}
     * @memberof ProfileRestrictions
     */
    reasonDescription?: string;
    /**
     * True indicates there are restrictions associated with the current profile.
     * @type {boolean}
     * @memberof ProfileRestrictions
     */
    restricted?: boolean;
}

/**
 * Check if a given object implements the ProfileRestrictions interface.
 */
export function instanceOfProfileRestrictions(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileRestrictionsFromJSON(json: any): ProfileRestrictions {
    return ProfileRestrictionsFromJSONTyped(json, false);
}

export function ProfileRestrictionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileRestrictions {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reason': !exists(json, 'reason') ? undefined : json['reason'],
        'reasonDescription': !exists(json, 'reasonDescription') ? undefined : json['reasonDescription'],
        'restricted': !exists(json, 'restricted') ? undefined : json['restricted'],
    };
}

export function ProfileRestrictionsToJSON(value?: ProfileRestrictions | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reason': value.reason,
        'reasonDescription': value.reasonDescription,
        'restricted': value.restricted,
    };
}

