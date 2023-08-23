/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Xchange Interface OXI API
 * APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 23.0.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CommunicationMethodType } from './CommunicationMethodType';
import {
    CommunicationMethodTypeFromJSON,
    CommunicationMethodTypeFromJSONTyped,
    CommunicationMethodTypeToJSON,
} from './CommunicationMethodType';
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
 * @interface CommunicationMethods
 */
export interface CommunicationMethods {
    /**
     * 
     * @type {Array<CommunicationMethodType>}
     * @memberof CommunicationMethods
     */
    communicationDetails?: Array<CommunicationMethodType>;
    /**
     * Indicates the index of the next applicable set(page).
     * @type {number}
     * @memberof CommunicationMethods
     */
    offset?: number;
    /**
     * Indicates number of records the API can return as per the API request limit sent. A maximum of 200 records can be only returned at a time.
     * @type {number}
     * @memberof CommunicationMethods
     */
    limit?: number;
    /**
     * Indicates number of records the API has returned actually as per the API request criteria.
     * @type {number}
     * @memberof CommunicationMethods
     */
    count?: number;
    /**
     * Indicates whether there are more records available to be returned as per the API request criteria or not.
     * @type {boolean}
     * @memberof CommunicationMethods
     */
    hasMore?: boolean;
    /**
     * Indicates total number of records available that can be returned as per the API request criteria.
     * @type {number}
     * @memberof CommunicationMethods
     */
    totalResults?: number;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CommunicationMethods
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CommunicationMethods
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CommunicationMethods interface.
 */
export function instanceOfCommunicationMethods(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CommunicationMethodsFromJSON(json: any): CommunicationMethods {
    return CommunicationMethodsFromJSONTyped(json, false);
}

export function CommunicationMethodsFromJSONTyped(json: any, ignoreDiscriminator: boolean): CommunicationMethods {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'communicationDetails': !exists(json, 'communicationDetails') ? undefined : ((json['communicationDetails'] as Array<any>).map(CommunicationMethodTypeFromJSON)),
        'offset': !exists(json, 'offset') ? undefined : json['offset'],
        'limit': !exists(json, 'limit') ? undefined : json['limit'],
        'count': !exists(json, 'count') ? undefined : json['count'],
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CommunicationMethodsToJSON(value?: CommunicationMethods | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'communicationDetails': value.communicationDetails === undefined ? undefined : ((value.communicationDetails as Array<any>).map(CommunicationMethodTypeToJSON)),
        'offset': value.offset,
        'limit': value.limit,
        'count': value.count,
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

