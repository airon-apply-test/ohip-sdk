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
import type { FixedChargeDetailType } from './FixedChargeDetailType';
import {
    FixedChargeDetailTypeFromJSON,
    FixedChargeDetailTypeFromJSONTyped,
    FixedChargeDetailTypeToJSON,
} from './FixedChargeDetailType';
import type { FixedChargeScheduleType } from './FixedChargeScheduleType';
import {
    FixedChargeScheduleTypeFromJSON,
    FixedChargeScheduleTypeFromJSONTyped,
    FixedChargeScheduleTypeToJSON,
} from './FixedChargeScheduleType';

/**
 * Holds fixed charge information.
 * @export
 * @interface FixedChargeType
 */
export interface FixedChargeType {
    /**
     * 
     * @type {FixedChargeScheduleType}
     * @memberof FixedChargeType
     */
    schedule?: FixedChargeScheduleType;
    /**
     * 
     * @type {FixedChargeDetailType}
     * @memberof FixedChargeType
     */
    charge?: FixedChargeDetailType;
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof FixedChargeType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof FixedChargeType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof FixedChargeType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof FixedChargeType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof FixedChargeType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof FixedChargeType
     */
    idExtension?: number;
}

/**
 * Check if a given object implements the FixedChargeType interface.
 */
export function instanceOfFixedChargeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function FixedChargeTypeFromJSON(json: any): FixedChargeType {
    return FixedChargeTypeFromJSONTyped(json, false);
}

export function FixedChargeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): FixedChargeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'schedule': !exists(json, 'schedule') ? undefined : FixedChargeScheduleTypeFromJSON(json['schedule']),
        'charge': !exists(json, 'charge') ? undefined : FixedChargeDetailTypeFromJSON(json['charge']),
        'url': !exists(json, 'url') ? undefined : json['url'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'instance': !exists(json, 'instance') ? undefined : json['instance'],
        'idContext': !exists(json, 'idContext') ? undefined : json['idContext'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'idExtension': !exists(json, 'idExtension') ? undefined : json['idExtension'],
    };
}

export function FixedChargeTypeToJSON(value?: FixedChargeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'schedule': FixedChargeScheduleTypeToJSON(value.schedule),
        'charge': FixedChargeDetailTypeToJSON(value.charge),
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
    };
}

