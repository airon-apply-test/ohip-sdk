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
import type { MembershipTransactionType } from './MembershipTransactionType';
import {
    MembershipTransactionTypeFromJSON,
    MembershipTransactionTypeFromJSONTyped,
    MembershipTransactionTypeToJSON,
} from './MembershipTransactionType';

/**
 * Response object for fetching a membership transaction.
 * @export
 * @interface MembershipTransaction
 */
export interface MembershipTransaction {
    /**
     * 
     * @type {MembershipTransactionType}
     * @memberof MembershipTransaction
     */
    membershipTransactionDetails?: MembershipTransactionType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof MembershipTransaction
     */
    links?: Array<InstanceLink>;
}

/**
 * Check if a given object implements the MembershipTransaction interface.
 */
export function instanceOfMembershipTransaction(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipTransactionFromJSON(json: any): MembershipTransaction {
    return MembershipTransactionFromJSONTyped(json, false);
}

export function MembershipTransactionFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipTransaction {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipTransactionDetails': !exists(json, 'membershipTransactionDetails') ? undefined : MembershipTransactionTypeFromJSON(json['membershipTransactionDetails']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
    };
}

export function MembershipTransactionToJSON(value?: MembershipTransaction | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipTransactionDetails': MembershipTransactionTypeToJSON(value.membershipTransactionDetails),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
    };
}

