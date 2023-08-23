/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud API for Customer Management Service
 * This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Defines descriptive name, unique identification, and basic information combination.
 * @export
 * @interface UniqueNameIDType
 */
export interface UniqueNameIDType {
    /**
     * Display name.
     * @type {string}
     * @memberof UniqueNameIDType
     */
    name?: string;
    /**
     * Phone number
     * @type {string}
     * @memberof UniqueNameIDType
     */
    phoneNumber?: string;
    /**
     * Email address
     * @type {string}
     * @memberof UniqueNameIDType
     */
    email?: string;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof UniqueNameIDType
     */
    primaryInd?: boolean;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof UniqueNameIDType
     */
    id?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof UniqueNameIDType
     */
    type?: string;
}

/**
 * Check if a given object implements the UniqueNameIDType interface.
 */
export function instanceOfUniqueNameIDType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UniqueNameIDTypeFromJSON(json: any): UniqueNameIDType {
    return UniqueNameIDTypeFromJSONTyped(json, false);
}

export function UniqueNameIDTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UniqueNameIDType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'phoneNumber': !exists(json, 'phoneNumber') ? undefined : json['phoneNumber'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'primaryInd': !exists(json, 'primaryInd') ? undefined : json['primaryInd'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function UniqueNameIDTypeToJSON(value?: UniqueNameIDType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'phoneNumber': value.phoneNumber,
        'email': value.email,
        'primaryInd': value.primaryInd,
        'id': value.id,
        'type': value.type,
    };
}

