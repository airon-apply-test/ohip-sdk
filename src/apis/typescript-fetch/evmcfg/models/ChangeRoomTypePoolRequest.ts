/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Event Configuration API
 * This API caters for Event Configuration in OPERA Cloud. <br /><There are operations to post, update, fetch and delete codes such as item inventory, function spaces, menu items and many more.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
import type { RoomTypePoolType } from './RoomTypePoolType';
import {
    RoomTypePoolTypeFromJSON,
    RoomTypePoolTypeFromJSONTyped,
    RoomTypePoolTypeToJSON,
} from './RoomTypePoolType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface ChangeRoomTypePoolRequest
 */
export interface ChangeRoomTypePoolRequest {
    /**
     * Collection of room type pool and associated room type(s).
     * @type {Array<RoomTypePoolType>}
     * @memberof ChangeRoomTypePoolRequest
     */
    roomPoolTypes?: Array<RoomTypePoolType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ChangeRoomTypePoolRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ChangeRoomTypePoolRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ChangeRoomTypePoolRequest interface.
 */
export function instanceOfChangeRoomTypePoolRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ChangeRoomTypePoolRequestFromJSON(json: any): ChangeRoomTypePoolRequest {
    return ChangeRoomTypePoolRequestFromJSONTyped(json, false);
}

export function ChangeRoomTypePoolRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChangeRoomTypePoolRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomPoolTypes': !exists(json, 'roomPoolTypes') ? undefined : ((json['roomPoolTypes'] as Array<any>).map(RoomTypePoolTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ChangeRoomTypePoolRequestToJSON(value?: ChangeRoomTypePoolRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomPoolTypes': value.roomPoolTypes === undefined ? undefined : ((value.roomPoolTypes as Array<any>).map(RoomTypePoolTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

