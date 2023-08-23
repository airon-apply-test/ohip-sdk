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
import type { BlockInventoryItemSourceType } from './BlockInventoryItemSourceType';
import {
    BlockInventoryItemSourceTypeFromJSON,
    BlockInventoryItemSourceTypeFromJSONTyped,
    BlockInventoryItemSourceTypeToJSON,
} from './BlockInventoryItemSourceType';
import type { DateRangeType } from './DateRangeType';
import {
    DateRangeTypeFromJSON,
    DateRangeTypeFromJSONTyped,
    DateRangeTypeToJSON,
} from './DateRangeType';
import type { ItemInfoType } from './ItemInfoType';
import {
    ItemInfoTypeFromJSON,
    ItemInfoTypeFromJSONTyped,
    ItemInfoTypeToJSON,
} from './ItemInfoType';

/**
 * An identifier used to uniquely reference an object in a system (e.g. an airline reservation reference, customer profile reference, booking confirmation number, or a reference to a previous availability quote).
 * @export
 * @interface BlockInventoryItemType
 */
export interface BlockInventoryItemType {
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof BlockInventoryItemType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof BlockInventoryItemType
     */
    idExtension?: number;
    /**
     * 
     * @type {ItemInfoType}
     * @memberof BlockInventoryItemType
     */
    item?: ItemInfoType;
    /**
     * Number of items booked.
     * @type {number}
     * @memberof BlockInventoryItemType
     */
    quantity?: number;
    /**
     * 
     * @type {DateRangeType}
     * @memberof BlockInventoryItemType
     */
    timeSpan?: DateRangeType;
    /**
     * 
     * @type {BlockInventoryItemSourceType}
     * @memberof BlockInventoryItemType
     */
    source?: BlockInventoryItemSourceType;
}

/**
 * Check if a given object implements the BlockInventoryItemType interface.
 */
export function instanceOfBlockInventoryItemType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockInventoryItemTypeFromJSON(json: any): BlockInventoryItemType {
    return BlockInventoryItemTypeFromJSONTyped(json, false);
}

export function BlockInventoryItemTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockInventoryItemType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'url': !exists(json, 'url') ? undefined : json['url'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'instance': !exists(json, 'instance') ? undefined : json['instance'],
        'idContext': !exists(json, 'idContext') ? undefined : json['idContext'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'idExtension': !exists(json, 'idExtension') ? undefined : json['idExtension'],
        'item': !exists(json, 'item') ? undefined : ItemInfoTypeFromJSON(json['item']),
        'quantity': !exists(json, 'quantity') ? undefined : json['quantity'],
        'timeSpan': !exists(json, 'timeSpan') ? undefined : DateRangeTypeFromJSON(json['timeSpan']),
        'source': !exists(json, 'source') ? undefined : BlockInventoryItemSourceTypeFromJSON(json['source']),
    };
}

export function BlockInventoryItemTypeToJSON(value?: BlockInventoryItemType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
        'item': ItemInfoTypeToJSON(value.item),
        'quantity': value.quantity,
        'timeSpan': DateRangeTypeToJSON(value.timeSpan),
        'source': BlockInventoryItemSourceTypeToJSON(value.source),
    };
}

