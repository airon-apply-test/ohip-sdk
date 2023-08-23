/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ProfileTypeType } from './ProfileTypeType';
import {
    ProfileTypeTypeFromJSON,
    ProfileTypeTypeFromJSONTyped,
    ProfileTypeTypeToJSON,
} from './ProfileTypeType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * 
 * @export
 * @interface RelationshipProfileType
 */
export interface RelationshipProfileType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof RelationshipProfileType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     * 
     * @type {ProfileTypeType}
     * @memberof RelationshipProfileType
     */
    profileType?: ProfileTypeType;
}

/**
 * Check if a given object implements the RelationshipProfileType interface.
 */
export function instanceOfRelationshipProfileType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RelationshipProfileTypeFromJSON(json: any): RelationshipProfileType {
    return RelationshipProfileTypeFromJSONTyped(json, false);
}

export function RelationshipProfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RelationshipProfileType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profileIdList': !exists(json, 'profileIdList') ? undefined : ((json['profileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'profileType': !exists(json, 'profileType') ? undefined : ProfileTypeTypeFromJSON(json['profileType']),
    };
}

export function RelationshipProfileTypeToJSON(value?: RelationshipProfileType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profileIdList': value.profileIdList === undefined ? undefined : ((value.profileIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'profileType': ProfileTypeTypeToJSON(value.profileType),
    };
}

