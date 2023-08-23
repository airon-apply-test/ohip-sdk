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
 * Information which uniquely identifies SFTP Configuration
 * @export
 * @interface SftpConfigurationType
 */
export interface SftpConfigurationType {
    /**
     * Unique id associated with this configuration
     * @type {number}
     * @memberof SftpConfigurationType
     */
    configurationId?: number;
    /**
     * SFTP destination
     * @type {string}
     * @memberof SftpConfigurationType
     */
    destination?: string;
    /**
     * Description of the destination, such as Shift Reports.
     * @type {string}
     * @memberof SftpConfigurationType
     */
    description?: string;
    /**
     * Indicates whether the configuration is inactive or not.
     * @type {boolean}
     * @memberof SftpConfigurationType
     */
    inactive?: boolean;
    /**
     * Hotel code
     * @type {string}
     * @memberof SftpConfigurationType
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the SftpConfigurationType interface.
 */
export function instanceOfSftpConfigurationType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SftpConfigurationTypeFromJSON(json: any): SftpConfigurationType {
    return SftpConfigurationTypeFromJSONTyped(json, false);
}

export function SftpConfigurationTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SftpConfigurationType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'configurationId': !exists(json, 'configurationId') ? undefined : json['configurationId'],
        'destination': !exists(json, 'destination') ? undefined : json['destination'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'inactive': !exists(json, 'inactive') ? undefined : json['inactive'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function SftpConfigurationTypeToJSON(value?: SftpConfigurationType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'configurationId': value.configurationId,
        'destination': value.destination,
        'description': value.description,
        'inactive': value.inactive,
        'hotelId': value.hotelId,
    };
}

