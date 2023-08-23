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
import type { HotelFolioTextType } from './HotelFolioTextType';
import {
    HotelFolioTextTypeFromJSON,
    HotelFolioTextTypeFromJSONTyped,
    HotelFolioTextTypeToJSON,
} from './HotelFolioTextType';

/**
 * A type which is used to insert Membership FolioTexts.
 * @export
 * @interface MembershipFolioTextsType
 */
export interface MembershipFolioTextsType {
    /**
     * Membership Type code.
     * @type {string}
     * @memberof MembershipFolioTextsType
     */
    membershipType?: string;
    /**
     * Membership Level code.
     * @type {string}
     * @memberof MembershipFolioTextsType
     */
    membershipLevel?: string;
    /**
     * Folio Texts.
     * @type {Array<Array<HotelFolioTextType>>}
     * @memberof MembershipFolioTextsType
     */
    folioTexts?: Array<Array<HotelFolioTextType>>;
}

/**
 * Check if a given object implements the MembershipFolioTextsType interface.
 */
export function instanceOfMembershipFolioTextsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function MembershipFolioTextsTypeFromJSON(json: any): MembershipFolioTextsType {
    return MembershipFolioTextsTypeFromJSONTyped(json, false);
}

export function MembershipFolioTextsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MembershipFolioTextsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'membershipType': !exists(json, 'membershipType') ? undefined : json['membershipType'],
        'membershipLevel': !exists(json, 'membershipLevel') ? undefined : json['membershipLevel'],
        'folioTexts': !exists(json, 'folioTexts') ? undefined : json['folioTexts'],
    };
}

export function MembershipFolioTextsTypeToJSON(value?: MembershipFolioTextsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'membershipType': value.membershipType,
        'membershipLevel': value.membershipLevel,
        'folioTexts': value.folioTexts,
    };
}

