/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Leisure Management API
 * APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ActivityLocationType } from './ActivityLocationType';
import {
    ActivityLocationTypeFromJSON,
    ActivityLocationTypeFromJSONTyped,
    ActivityLocationTypeToJSON,
} from './ActivityLocationType';
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
 * 
 * @export
 * @interface PostActivityLocationsRequest
 */
export interface PostActivityLocationsRequest {
    /**
     * Collection of Activity Locations.
     * @type {Array<ActivityLocationType>}
     * @memberof PostActivityLocationsRequest
     */
    activityLocations?: Array<ActivityLocationType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostActivityLocationsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostActivityLocationsRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostActivityLocationsRequest interface.
 */
export function instanceOfPostActivityLocationsRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostActivityLocationsRequestFromJSON(json: any): PostActivityLocationsRequest {
    return PostActivityLocationsRequestFromJSONTyped(json, false);
}

export function PostActivityLocationsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostActivityLocationsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activityLocations': !exists(json, 'activityLocations') ? undefined : ((json['activityLocations'] as Array<any>).map(ActivityLocationTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostActivityLocationsRequestToJSON(value?: PostActivityLocationsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activityLocations': value.activityLocations === undefined ? undefined : ((value.activityLocations as Array<any>).map(ActivityLocationTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

