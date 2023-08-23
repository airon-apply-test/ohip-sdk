/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Room Configuration API
 * APIs to cater for room configuration, such as configuring room types, room Classes, creating new room features, or updating housekeeping room maintenance reasons.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
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
import type { RoomPotentialType } from './RoomPotentialType';
import {
    RoomPotentialTypeFromJSON,
    RoomPotentialTypeFromJSONTyped,
    RoomPotentialTypeToJSON,
} from './RoomPotentialType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PutRoomPotentialsRequest
 */
export interface PutRoomPotentialsRequest {
    /**
     * List of Room Potentials.
     * @type {Array<RoomPotentialType>}
     * @memberof PutRoomPotentialsRequest
     */
    roomPotentials?: Array<RoomPotentialType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PutRoomPotentialsRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PutRoomPotentialsRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PutRoomPotentialsRequest interface.
 */
export function instanceOfPutRoomPotentialsRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PutRoomPotentialsRequestFromJSON(json: any): PutRoomPotentialsRequest {
    return PutRoomPotentialsRequestFromJSONTyped(json, false);
}

export function PutRoomPotentialsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PutRoomPotentialsRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomPotentials': !exists(json, 'roomPotentials') ? undefined : ((json['roomPotentials'] as Array<any>).map(RoomPotentialTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PutRoomPotentialsRequestToJSON(value?: PutRoomPotentialsRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomPotentials': value.roomPotentials === undefined ? undefined : ((value.roomPotentials as Array<any>).map(RoomPotentialTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

