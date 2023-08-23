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
import type { ActivityResultsConfigDetailType } from './ActivityResultsConfigDetailType';
import {
    ActivityResultsConfigDetailTypeFromJSON,
    ActivityResultsConfigDetailTypeFromJSONTyped,
    ActivityResultsConfigDetailTypeToJSON,
} from './ActivityResultsConfigDetailType';
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
 * Response object for fetching configured Activity Results.
 * @export
 * @interface ActivityResultsConfigDetails
 */
export interface ActivityResultsConfigDetails {
    /**
     * List of configured Activity Results.
     * @type {Array<ActivityResultsConfigDetailType>}
     * @memberof ActivityResultsConfigDetails
     */
    activityResultsConfiguration?: Array<ActivityResultsConfigDetailType>;
    /**
     * Evaluated total page count based on the requested max fetch count.
     * @type {number}
     * @memberof ActivityResultsConfigDetails
     */
    totalPages?: number;
    /**
     * Index or initial index of the set(page) being requested. If the index goes out of the bounds of the total set count then no data will be returned.
     * @type {number}
     * @memberof ActivityResultsConfigDetails
     */
    offset?: number;
    /**
     * Indicates maximum number of records a Web Service should return.
     * @type {number}
     * @memberof ActivityResultsConfigDetails
     */
    limit?: number;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof ActivityResultsConfigDetails
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof ActivityResultsConfigDetails
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof ActivityResultsConfigDetails
     */
    count?: number;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ActivityResultsConfigDetails
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ActivityResultsConfigDetails
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ActivityResultsConfigDetails interface.
 */
export function instanceOfActivityResultsConfigDetails(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ActivityResultsConfigDetailsFromJSON(json: any): ActivityResultsConfigDetails {
    return ActivityResultsConfigDetailsFromJSONTyped(json, false);
}

export function ActivityResultsConfigDetailsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityResultsConfigDetails {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activityResultsConfiguration': !exists(json, 'activityResultsConfiguration') ? undefined : ((json['activityResultsConfiguration'] as Array<any>).map(ActivityResultsConfigDetailTypeFromJSON)),
        'totalPages': !exists(json, 'totalPages') ? undefined : json['totalPages'],
        'offset': !exists(json, 'offset') ? undefined : json['offset'],
        'limit': !exists(json, 'limit') ? undefined : json['limit'],
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ActivityResultsConfigDetailsToJSON(value?: ActivityResultsConfigDetails | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activityResultsConfiguration': value.activityResultsConfiguration === undefined ? undefined : ((value.activityResultsConfiguration as Array<any>).map(ActivityResultsConfigDetailTypeToJSON)),
        'totalPages': value.totalPages,
        'offset': value.offset,
        'limit': value.limit,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

