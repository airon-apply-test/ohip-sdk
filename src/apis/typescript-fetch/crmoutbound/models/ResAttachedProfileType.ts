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
import type { ResProfileTypeType } from './ResProfileTypeType';
import {
    ResProfileTypeTypeFromJSON,
    ResProfileTypeTypeFromJSONTyped,
    ResProfileTypeTypeToJSON,
} from './ResProfileTypeType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * 
 * @export
 * @interface ResAttachedProfileType
 */
export interface ResAttachedProfileType {
    /**
     * Attached profile name
     * @type {string}
     * @memberof ResAttachedProfileType
     */
    name?: string;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof ResAttachedProfileType
     */
    profileIdList?: Array<UniqueIDType>;
    /**
     * 
     * @type {ResProfileTypeType}
     * @memberof ResAttachedProfileType
     */
    resProfileType?: ResProfileTypeType;
}

/**
 * Check if a given object implements the ResAttachedProfileType interface.
 */
export function instanceOfResAttachedProfileType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResAttachedProfileTypeFromJSON(json: any): ResAttachedProfileType {
    return ResAttachedProfileTypeFromJSONTyped(json, false);
}

export function ResAttachedProfileTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResAttachedProfileType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'profileIdList': !exists(json, 'profileIdList') ? undefined : ((json['profileIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'resProfileType': !exists(json, 'resProfileType') ? undefined : ResProfileTypeTypeFromJSON(json['resProfileType']),
    };
}

export function ResAttachedProfileTypeToJSON(value?: ResAttachedProfileType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'profileIdList': value.profileIdList === undefined ? undefined : ((value.profileIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'resProfileType': ResProfileTypeTypeToJSON(value.resProfileType),
    };
}

