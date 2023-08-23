/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity Management API
 * APIs to cater for Activity Configuration functionality in OPERA Cloud. In this module you can retrieve, create, update Activity configuration codes, for example create a new Activity Type.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ActivityConfigTypeDetailType } from './ActivityConfigTypeDetailType';
import {
    ActivityConfigTypeDetailTypeFromJSON,
    ActivityConfigTypeDetailTypeFromJSONTyped,
    ActivityConfigTypeDetailTypeToJSON,
} from './ActivityConfigTypeDetailType';
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
 * @interface PostActivityTypesRequest
 */
export interface PostActivityTypesRequest {
    /**
     * Detailed information of activity type.
     * @type {Array<ActivityConfigTypeDetailType>}
     * @memberof PostActivityTypesRequest
     */
    activityConfigTypes?: Array<ActivityConfigTypeDetailType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostActivityTypesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostActivityTypesRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostActivityTypesRequest interface.
 */
export function instanceOfPostActivityTypesRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostActivityTypesRequestFromJSON(json: any): PostActivityTypesRequest {
    return PostActivityTypesRequestFromJSONTyped(json, false);
}

export function PostActivityTypesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostActivityTypesRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activityConfigTypes': !exists(json, 'activityConfigTypes') ? undefined : ((json['activityConfigTypes'] as Array<any>).map(ActivityConfigTypeDetailTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostActivityTypesRequestToJSON(value?: PostActivityTypesRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activityConfigTypes': value.activityConfigTypes === undefined ? undefined : ((value.activityConfigTypes as Array<any>).map(ActivityConfigTypeDetailTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

