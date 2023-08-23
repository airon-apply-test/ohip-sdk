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
import type { TranslationTextType2000 } from './TranslationTextType2000';
import {
    TranslationTextType2000FromJSON,
    TranslationTextType2000FromJSONTyped,
    TranslationTextType2000ToJSON,
} from './TranslationTextType2000';

/**
 * Contains Common Master configuration detail.
 * @export
 * @interface ProfileRestrictionReasonType
 */
export interface ProfileRestrictionReasonType {
    /**
     * Common Master unique code.
     * @type {string}
     * @memberof ProfileRestrictionReasonType
     */
    code?: string;
    /**
     * 
     * @type {TranslationTextType2000}
     * @memberof ProfileRestrictionReasonType
     */
    description?: TranslationTextType2000;
    /**
     * Common Master record sequence number.
     * @type {number}
     * @memberof ProfileRestrictionReasonType
     */
    displayOrder?: number;
}

/**
 * Check if a given object implements the ProfileRestrictionReasonType interface.
 */
export function instanceOfProfileRestrictionReasonType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ProfileRestrictionReasonTypeFromJSON(json: any): ProfileRestrictionReasonType {
    return ProfileRestrictionReasonTypeFromJSONTyped(json, false);
}

export function ProfileRestrictionReasonTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileRestrictionReasonType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'description': !exists(json, 'description') ? undefined : TranslationTextType2000FromJSON(json['description']),
        'displayOrder': !exists(json, 'displayOrder') ? undefined : json['displayOrder'],
    };
}

export function ProfileRestrictionReasonTypeToJSON(value?: ProfileRestrictionReasonType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'description': TranslationTextType2000ToJSON(value.description),
        'displayOrder': value.displayOrder,
    };
}

