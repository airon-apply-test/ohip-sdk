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
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { TemplateDeviceLocationType } from './TemplateDeviceLocationType';
import {
    TemplateDeviceLocationTypeFromJSON,
    TemplateDeviceLocationTypeFromJSONTyped,
    TemplateDeviceLocationTypeToJSON,
} from './TemplateDeviceLocationType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PostTemplateDeviceLocationsRequest
 */
export interface PostTemplateDeviceLocationsRequest {
    /**
     * List of Device locations at template level.
     * @type {Array<TemplateDeviceLocationType>}
     * @memberof PostTemplateDeviceLocationsRequest
     */
    templateDeviceLocations?: Array<TemplateDeviceLocationType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostTemplateDeviceLocationsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostTemplateDeviceLocationsRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostTemplateDeviceLocationsRequest interface.
 */
export function instanceOfPostTemplateDeviceLocationsRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostTemplateDeviceLocationsRequestFromJSON(json: any): PostTemplateDeviceLocationsRequest {
    return PostTemplateDeviceLocationsRequestFromJSONTyped(json, false);
}

export function PostTemplateDeviceLocationsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostTemplateDeviceLocationsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'templateDeviceLocations': !exists(json, 'templateDeviceLocations') ? undefined : ((json['templateDeviceLocations'] as Array<any>).map(TemplateDeviceLocationTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostTemplateDeviceLocationsRequestToJSON(value?: PostTemplateDeviceLocationsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'templateDeviceLocations': value.templateDeviceLocations === undefined ? undefined : ((value.templateDeviceLocations as Array<any>).map(TemplateDeviceLocationTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

