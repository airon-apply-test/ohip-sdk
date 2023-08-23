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
import type { TrackItId } from './TrackItId';
import {
    TrackItIdFromJSON,
    TrackItIdFromJSONTyped,
    TrackItIdToJSON,
} from './TrackItId';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Individual activity log related to the Track It ticket.
 * @export
 * @interface TrackItLogType
 */
export interface TrackItLogType {
    /**
     * 
     * @type {string}
     * @memberof TrackItLogType
     */
    hotelId?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof TrackItLogType
     */
    trackItLogId?: UniqueIDType;
    /**
     * 
     * @type {TrackItId}
     * @memberof TrackItLogType
     */
    trackItId?: TrackItId;
    /**
     * 
     * @type {string}
     * @memberof TrackItLogType
     */
    type?: string;
    /**
     * 
     * @type {string}
     * @memberof TrackItLogType
     */
    description?: string;
    /**
     * 
     * @type {Date}
     * @memberof TrackItLogType
     */
    businessDate?: Date;
    /**
     * 
     * @type {number}
     * @memberof TrackItLogType
     */
    logUserId?: number;
    /**
     * 
     * @type {string}
     * @memberof TrackItLogType
     */
    logUserName?: string;
    /**
     * 
     * @type {Date}
     * @memberof TrackItLogType
     */
    logDate?: Date;
}

/**
 * Check if a given object implements the TrackItLogType interface.
 */
export function instanceOfTrackItLogType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TrackItLogTypeFromJSON(json: any): TrackItLogType {
    return TrackItLogTypeFromJSONTyped(json, false);
}

export function TrackItLogTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrackItLogType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'trackItLogId': !exists(json, 'trackItLogId') ? undefined : UniqueIDTypeFromJSON(json['trackItLogId']),
        'trackItId': !exists(json, 'trackItId') ? undefined : TrackItIdFromJSON(json['trackItId']),
        'type': !exists(json, 'type') ? undefined : json['type'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'businessDate': !exists(json, 'businessDate') ? undefined : (new Date(json['businessDate'])),
        'logUserId': !exists(json, 'logUserId') ? undefined : json['logUserId'],
        'logUserName': !exists(json, 'logUserName') ? undefined : json['logUserName'],
        'logDate': !exists(json, 'logDate') ? undefined : (new Date(json['logDate'])),
    };
}

export function TrackItLogTypeToJSON(value?: TrackItLogType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'trackItLogId': UniqueIDTypeToJSON(value.trackItLogId),
        'trackItId': TrackItIdToJSON(value.trackItId),
        'type': value.type,
        'description': value.description,
        'businessDate': value.businessDate === undefined ? undefined : (value.businessDate.toISOString().substr(0,10)),
        'logUserId': value.logUserId,
        'logUserName': value.logUserName,
        'logDate': value.logDate === undefined ? undefined : (value.logDate.toISOString()),
    };
}

