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
import type { AccumulatedBusinessEventStatusType } from './AccumulatedBusinessEventStatusType';
import {
    AccumulatedBusinessEventStatusTypeFromJSON,
    AccumulatedBusinessEventStatusTypeFromJSONTyped,
    AccumulatedBusinessEventStatusTypeToJSON,
} from './AccumulatedBusinessEventStatusType';

/**
 * Business Event message details
 * @export
 * @interface AccumulatedBusinessEventType
 */
export interface AccumulatedBusinessEventType {
    /**
     * Primary Key identifier assigned to the business event.
     * @type {string}
     * @memberof AccumulatedBusinessEventType
     */
    primaryKey?: string;
    /**
     * 
     * @type {AccumulatedBusinessEventStatusType}
     * @memberof AccumulatedBusinessEventType
     */
    status?: AccumulatedBusinessEventStatusType;
    /**
     * Data when the message was created by external systems
     * @type {Date}
     * @memberof AccumulatedBusinessEventType
     */
    createDate?: Date;
    /**
     * The interface this message was created for
     * @type {string}
     * @memberof AccumulatedBusinessEventType
     */
    _interface?: string;
    /**
     * Data module of outgoing message. This identifies the kind of message sent (e.g., reservation, profile, rate, block, rate restriction, inventory, and result)
     * @type {string}
     * @memberof AccumulatedBusinessEventType
     */
    module?: string;
    /**
     * Property this message was sent from.
     * @type {string}
     * @memberof AccumulatedBusinessEventType
     */
    hotelId?: string;
}

/**
 * Check if a given object implements the AccumulatedBusinessEventType interface.
 */
export function instanceOfAccumulatedBusinessEventType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AccumulatedBusinessEventTypeFromJSON(json: any): AccumulatedBusinessEventType {
    return AccumulatedBusinessEventTypeFromJSONTyped(json, false);
}

export function AccumulatedBusinessEventTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AccumulatedBusinessEventType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'primaryKey': !exists(json, 'primaryKey') ? undefined : json['primaryKey'],
        'status': !exists(json, 'status') ? undefined : AccumulatedBusinessEventStatusTypeFromJSON(json['status']),
        'createDate': !exists(json, 'createDate') ? undefined : (new Date(json['createDate'])),
        '_interface': !exists(json, 'interface') ? undefined : json['interface'],
        'module': !exists(json, 'module') ? undefined : json['module'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
    };
}

export function AccumulatedBusinessEventTypeToJSON(value?: AccumulatedBusinessEventType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'primaryKey': value.primaryKey,
        'status': AccumulatedBusinessEventStatusTypeToJSON(value.status),
        'createDate': value.createDate === undefined ? undefined : (value.createDate.toISOString()),
        'interface': value._interface,
        'module': value.module,
        'hotelId': value.hotelId,
    };
}

