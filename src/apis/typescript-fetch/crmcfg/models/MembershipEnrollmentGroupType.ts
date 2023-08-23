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
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';

/**
 * Base type provides information about Membership Market/Property Groups Configuration.
 * @export
 * @interface MembershipEnrollmentGroupType
 */
export interface MembershipEnrollmentGroupType {
    /**
     * Code is used to identify a Membership Market/Resort Group.
     * @type {string}
     * @memberof MembershipEnrollmentGroupType
     */
    code?: string;
    /**
     * Description of the Membership Market/Propety Group.
     * @type {string}
     * @memberof MembershipEnrollmentGroupType
     */
    description?: string;
    /**
     * Membership Market/Property Groups display sequence Number
     * @type {number}
     * @memberof MembershipEnrollmentGroupType
     */
    displaySequence?: number;
    /**
     * Membership enrollment code code and description.
     * @type {Array<CodeDescriptionType>}
     * @memberof MembershipEnrollmentGroupType
     */
    enrollmentCodes?: Array<CodeDescriptionType>;
}

/**
 * Check if a given object implements the MembershipEnrollmentGroupType interface.
 */
export function instanceOfMembershipEnrollmentGroupType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipEnrollmentGroupTypeFromJSON(json: any): MembershipEnrollmentGroupType {
    return MembershipEnrollmentGroupTypeFromJSONTyped(json, false);
}

export function MembershipEnrollmentGroupTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipEnrollmentGroupType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'displaySequence': !exists(json, 'displaySequence') ? undefined : json['displaySequence'],
        'enrollmentCodes': !exists(json, 'enrollmentCodes') ? undefined : ((json['enrollmentCodes'] as Array<any>).map(CodeDescriptionTypeFromJSON)),
    };
}

export function MembershipEnrollmentGroupTypeToJSON(value?: MembershipEnrollmentGroupType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'description': value.description,
        'displaySequence': value.displaySequence,
        'enrollmentCodes': value.enrollmentCodes === undefined ? undefined : ((value.enrollmentCodes as Array<any>).map(CodeDescriptionTypeToJSON)),
    };
}

