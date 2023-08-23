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
import type { HotelInterfaceDetailType } from './HotelInterfaceDetailType';
import {
    HotelInterfaceDetailTypeFromJSON,
    HotelInterfaceDetailTypeFromJSONTyped,
    HotelInterfaceDetailTypeToJSON,
} from './HotelInterfaceDetailType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Request object for creating a new Hotel Interface.
 * @export
 * @interface HotelInterface
 */
export interface HotelInterface {
    /**
     * 
     * @type {HotelInterfaceDetailType}
     * @memberof HotelInterface
     */
    details?: HotelInterfaceDetailType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof HotelInterface
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof HotelInterface
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the HotelInterface interface.
 */
export function instanceOfHotelInterface(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HotelInterfaceFromJSON(json: any): HotelInterface {
    return HotelInterfaceFromJSONTyped(json, false);
}

export function HotelInterfaceFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInterface {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'details': !exists(json, 'details') ? undefined : HotelInterfaceDetailTypeFromJSON(json['details']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function HotelInterfaceToJSON(value?: HotelInterface | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'details': HotelInterfaceDetailTypeToJSON(value.details),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

