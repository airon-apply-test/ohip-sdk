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
/**
 * Membership type rule ID.
 * @export
 * @interface MembershipTypeRuleIDType
 */
export interface MembershipTypeRuleIDType {
    /**
     * Membership type of the Rule.
     * @type {string}
     * @memberof MembershipTypeRuleIDType
     */
    membershipType?: string;
    /**
     * Sequence number of the membership type rule.
     * @type {number}
     * @memberof MembershipTypeRuleIDType
     */
    membershipPointsSequence?: number;
}

/**
 * Check if a given object implements the MembershipTypeRuleIDType interface.
 */
export function instanceOfMembershipTypeRuleIDType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipTypeRuleIDTypeFromJSON(json: any): MembershipTypeRuleIDType {
    return MembershipTypeRuleIDTypeFromJSONTyped(json, false);
}

export function MembershipTypeRuleIDTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipTypeRuleIDType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipType': !exists(json, 'membershipType') ? undefined : json['membershipType'],
        'membershipPointsSequence': !exists(json, 'membershipPointsSequence') ? undefined : json['membershipPointsSequence'],
    };
}

export function MembershipTypeRuleIDTypeToJSON(value?: MembershipTypeRuleIDType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipType': value.membershipType,
        'membershipPointsSequence': value.membershipPointsSequence,
    };
}

