/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Information regarding associated reservations to the reservation.
 * @export
 * @interface AssociatedResType
 */
export interface AssociatedResType {
    /**
     * 
     * @type {UniqueIDType}
     * @memberof AssociatedResType
     */
    parentReservation?: UniqueIDType;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof AssociatedResType
     */
    originalMultiRoomRes?: UniqueIDType;
}

/**
 * Check if a given object implements the AssociatedResType interface.
 */
export function instanceOfAssociatedResType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AssociatedResTypeFromJSON(json: any): AssociatedResType {
    return AssociatedResTypeFromJSONTyped(json, false);
}

export function AssociatedResTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AssociatedResType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'parentReservation': !exists(json, 'parentReservation') ? undefined : UniqueIDTypeFromJSON(json['parentReservation']),
        'originalMultiRoomRes': !exists(json, 'originalMultiRoomRes') ? undefined : UniqueIDTypeFromJSON(json['originalMultiRoomRes']),
    };
}

export function AssociatedResTypeToJSON(value?: AssociatedResType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'parentReservation': UniqueIDTypeToJSON(value.parentReservation),
        'originalMultiRoomRes': UniqueIDTypeToJSON(value.originalMultiRoomRes),
    };
}

