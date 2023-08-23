/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
import type { LinkedActivitiesType } from './LinkedActivitiesType';
import {
    LinkedActivitiesTypeFromJSON,
    LinkedActivitiesTypeFromJSONTyped,
    LinkedActivitiesTypeToJSON,
} from './LinkedActivitiesType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PostLinkedActivitiesRequest
 */
export interface PostLinkedActivitiesRequest {
    /**
     * 
     * @type {LinkedActivitiesType}
     * @memberof PostLinkedActivitiesRequest
     */
    linkedActivityDetails?: LinkedActivitiesType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostLinkedActivitiesRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostLinkedActivitiesRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostLinkedActivitiesRequest interface.
 */
export function instanceOfPostLinkedActivitiesRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostLinkedActivitiesRequestFromJSON(json: any): PostLinkedActivitiesRequest {
    return PostLinkedActivitiesRequestFromJSONTyped(json, false);
}

export function PostLinkedActivitiesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostLinkedActivitiesRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'linkedActivityDetails': !exists(json, 'linkedActivityDetails') ? undefined : LinkedActivitiesTypeFromJSON(json['linkedActivityDetails']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostLinkedActivitiesRequestToJSON(value?: PostLinkedActivitiesRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'linkedActivityDetails': LinkedActivitiesTypeToJSON(value.linkedActivityDetails),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

