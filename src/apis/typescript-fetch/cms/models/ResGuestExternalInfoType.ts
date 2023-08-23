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
 * Specifies Company or Travel Agent profile using IATA or Corp. No.
 * @export
 * @interface ResGuestExternalInfoType
 */
export interface ResGuestExternalInfoType {
    /**
     * Given name, first name or names
     * @type {string}
     * @memberof ResGuestExternalInfoType
     */
    givenName?: string;
    /**
     * Family name, last name.
     * @type {string}
     * @memberof ResGuestExternalInfoType
     */
    surname?: string;
}

/**
 * Check if a given object implements the ResGuestExternalInfoType interface.
 */
export function instanceOfResGuestExternalInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResGuestExternalInfoTypeFromJSON(json: any): ResGuestExternalInfoType {
    return ResGuestExternalInfoTypeFromJSONTyped(json, false);
}

export function ResGuestExternalInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResGuestExternalInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'givenName': !exists(json, 'givenName') ? undefined : json['givenName'],
        'surname': !exists(json, 'surname') ? undefined : json['surname'],
    };
}

export function ResGuestExternalInfoTypeToJSON(value?: ResGuestExternalInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'givenName': value.givenName,
        'surname': value.surname,
    };
}

