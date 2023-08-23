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
import type { MembershipClaimDetailsType } from './MembershipClaimDetailsType';
import {
    MembershipClaimDetailsTypeFromJSON,
    MembershipClaimDetailsTypeFromJSONTyped,
    MembershipClaimDetailsTypeToJSON,
} from './MembershipClaimDetailsType';

/**
 * Request object for creating membership claim.
 * @export
 * @interface MembershipClaim
 */
export interface MembershipClaim {
    /**
     * 
     * @type {MembershipClaimDetailsType}
     * @memberof MembershipClaim
     */
    membershipClaimDetails?: MembershipClaimDetailsType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof MembershipClaim
     */
    links?: Array<InstanceLink>;
}

/**
 * Check if a given object implements the MembershipClaim interface.
 */
export function instanceOfMembershipClaim(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipClaimFromJSON(json: any): MembershipClaim {
    return MembershipClaimFromJSONTyped(json, false);
}

export function MembershipClaimFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipClaim {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipClaimDetails': !exists(json, 'membershipClaimDetails') ? undefined : MembershipClaimDetailsTypeFromJSON(json['membershipClaimDetails']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
    };
}

export function MembershipClaimToJSON(value?: MembershipClaim | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipClaimDetails': MembershipClaimDetailsTypeToJSON(value.membershipClaimDetails),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
    };
}

