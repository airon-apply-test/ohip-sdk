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
 * This contains a generic code and description information.
 * @export
 * @interface CodeDescriptionType
 */
export interface CodeDescriptionType {
    /**
     * Code.
     * @type {string}
     * @memberof CodeDescriptionType
     */
    code?: string;
    /**
     * description.
     * @type {string}
     * @memberof CodeDescriptionType
     */
    description?: string;
}

/**
 * Check if a given object implements the CodeDescriptionType interface.
 */
export function instanceOfCodeDescriptionType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CodeDescriptionTypeFromJSON(json: any): CodeDescriptionType {
    return CodeDescriptionTypeFromJSONTyped(json, false);
}

export function CodeDescriptionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CodeDescriptionType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'code': !exists(json, 'code') ? undefined : json['code'],
        'description': !exists(json, 'description') ? undefined : json['description'],
    };
}

export function CodeDescriptionTypeToJSON(value?: CodeDescriptionType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'code': value.code,
        'description': value.description,
    };
}

