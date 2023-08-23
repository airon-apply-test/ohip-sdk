/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Integration Configuration API
 * APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Generic self-contained request object that is used when updating and fetching UDFs.
 * @export
 * @interface UDFMappingType
 */
export interface UDFMappingType {
    /**
     * Name of the user-defined function (UDF) field.
     * @type {string}
     * @memberof UDFMappingType
     */
    uDFName?: string;
    /**
     * Label the user-defined function (UDF) field shared with external system which will be send in the request message for mapping of actual UDF Name.
     * @type {string}
     * @memberof UDFMappingType
     */
    uDFLabel?: string;
    /**
     * A flag to indicate whether the UDF mapping is active in the Interface system or not.
     * @type {boolean}
     * @memberof UDFMappingType
     */
    active?: boolean;
    /**
     * A flag to indicate whether the update is allowed in the Interface system or not.
     * @type {boolean}
     * @memberof UDFMappingType
     */
    updateAllowed?: boolean;
    /**
     * A reference to the type of object defined by the Type for Reservation and Profile Type.
     * @type {string}
     * @memberof UDFMappingType
     */
    type?: string;
}

/**
 * Check if a given object implements the UDFMappingType interface.
 */
export function instanceOfUDFMappingType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UDFMappingTypeFromJSON(json: any): UDFMappingType {
    return UDFMappingTypeFromJSONTyped(json, false);
}

export function UDFMappingTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UDFMappingType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'uDFName': !exists(json, 'uDFName') ? undefined : json['uDFName'],
        'uDFLabel': !exists(json, 'uDFLabel') ? undefined : json['uDFLabel'],
        'active': !exists(json, 'active') ? undefined : json['active'],
        'updateAllowed': !exists(json, 'updateAllowed') ? undefined : json['updateAllowed'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function UDFMappingTypeToJSON(value?: UDFMappingType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'uDFName': value.uDFName,
        'uDFLabel': value.uDFLabel,
        'active': value.active,
        'updateAllowed': value.updateAllowed,
        'type': value.type,
    };
}

