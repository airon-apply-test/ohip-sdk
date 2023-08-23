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
import type { MembershipTypeChangeInstructionType } from './MembershipTypeChangeInstructionType';
import {
    MembershipTypeChangeInstructionTypeFromJSON,
    MembershipTypeChangeInstructionTypeFromJSONTyped,
    MembershipTypeChangeInstructionTypeToJSON,
} from './MembershipTypeChangeInstructionType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for changing membership type.
 * @export
 * @interface MembershipTypeChangeInstructions
 */
export interface MembershipTypeChangeInstructions {
    /**
     * A collection of MembershipTypes with information that needs to be changed.
     * @type {Array<MembershipTypeChangeInstructionType>}
     * @memberof MembershipTypeChangeInstructions
     */
    membershipTypeChangeInstructions?: Array<MembershipTypeChangeInstructionType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof MembershipTypeChangeInstructions
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof MembershipTypeChangeInstructions
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the MembershipTypeChangeInstructions interface.
 */
export function instanceOfMembershipTypeChangeInstructions(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipTypeChangeInstructionsFromJSON(json: any): MembershipTypeChangeInstructions {
    return MembershipTypeChangeInstructionsFromJSONTyped(json, false);
}

export function MembershipTypeChangeInstructionsFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipTypeChangeInstructions {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipTypeChangeInstructions': !exists(json, 'membershipTypeChangeInstructions') ? undefined : ((json['membershipTypeChangeInstructions'] as Array<any>).map(MembershipTypeChangeInstructionTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function MembershipTypeChangeInstructionsToJSON(value?: MembershipTypeChangeInstructions | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipTypeChangeInstructions': value.membershipTypeChangeInstructions === undefined ? undefined : ((value.membershipTypeChangeInstructions as Array<any>).map(MembershipTypeChangeInstructionTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

