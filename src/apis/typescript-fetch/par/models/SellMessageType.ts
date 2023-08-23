/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { UsedInModuleType } from './UsedInModuleType';
import {
    UsedInModuleTypeFromJSON,
    UsedInModuleTypeFromJSONTyped,
    UsedInModuleTypeToJSON,
} from './UsedInModuleType';

/**
 * The SellMessagesType defines the standard attributes that will be retrieved with the sell message.
 * @export
 * @interface SellMessageType
 */
export interface SellMessageType {
    /**
     * This is the message text.
     * @type {string}
     * @memberof SellMessageType
     */
    message?: string;
    /**
     * This is the Central Reservation office code.
     * @type {string}
     * @memberof SellMessageType
     */
    croCode?: string;
    /**
     * This is the chain code.
     * @type {string}
     * @memberof SellMessageType
     */
    chainCode?: string;
    /**
     * This is the hotel code or resort.
     * @type {string}
     * @memberof SellMessageType
     */
    hotelId?: string;
    /**
     * This represents the room type code.
     * @type {string}
     * @memberof SellMessageType
     */
    roomType?: string;
    /**
     * This represents the rate plan code of the room type.
     * @type {string}
     * @memberof SellMessageType
     */
    ratePlanCode?: string;
    /**
     * This is the language code.
     * @type {string}
     * @memberof SellMessageType
     */
    languageCode?: string;
    /**
     * Flag which tells if this message is Sticky (Y/N).
     * @type {string}
     * @memberof SellMessageType
     */
    stickyFlagYn?: string;
    /**
     * 
     * @type {UsedInModuleType}
     * @memberof SellMessageType
     */
    usedInModule?: UsedInModuleType;
    /**
     * This is the Begin date for the configured SellMessage.
     * @type {Date}
     * @memberof SellMessageType
     */
    beginDate?: Date;
    /**
     * URL that identifies the location associated with the record identified by the UniqueID.
     * @type {string}
     * @memberof SellMessageType
     */
    url?: string;
    /**
     * A reference to the type of object defined by the UniqueID element.
     * @type {string}
     * @memberof SellMessageType
     */
    type?: string;
    /**
     * The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
     * @type {string}
     * @memberof SellMessageType
     */
    instance?: string;
    /**
     * Used to identify the source of the identifier (e.g., IATA, ABTA).
     * @type {string}
     * @memberof SellMessageType
     */
    idContext?: string;
    /**
     * A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
     * @type {string}
     * @memberof SellMessageType
     */
    id?: string;
    /**
     * Additional identifying value assigned by the creating system.
     * @type {number}
     * @memberof SellMessageType
     */
    idExtension?: number;
}

/**
 * Check if a given object implements the SellMessageType interface.
 */
export function instanceOfSellMessageType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function SellMessageTypeFromJSON(json: any): SellMessageType {
    return SellMessageTypeFromJSONTyped(json, false);
}

export function SellMessageTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): SellMessageType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'message': !exists(json, 'message') ? undefined : json['message'],
        'croCode': !exists(json, 'croCode') ? undefined : json['croCode'],
        'chainCode': !exists(json, 'chainCode') ? undefined : json['chainCode'],
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'roomType': !exists(json, 'roomType') ? undefined : json['roomType'],
        'ratePlanCode': !exists(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'languageCode': !exists(json, 'languageCode') ? undefined : json['languageCode'],
        'stickyFlagYn': !exists(json, 'stickyFlagYn') ? undefined : json['stickyFlagYn'],
        'usedInModule': !exists(json, 'usedInModule') ? undefined : UsedInModuleTypeFromJSON(json['usedInModule']),
        'beginDate': !exists(json, 'beginDate') ? undefined : (new Date(json['beginDate'])),
        'url': !exists(json, 'url') ? undefined : json['url'],
        'type': !exists(json, 'type') ? undefined : json['type'],
        'instance': !exists(json, 'instance') ? undefined : json['instance'],
        'idContext': !exists(json, 'idContext') ? undefined : json['idContext'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'idExtension': !exists(json, 'idExtension') ? undefined : json['idExtension'],
    };
}

export function SellMessageTypeToJSON(value?: SellMessageType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'message': value.message,
        'croCode': value.croCode,
        'chainCode': value.chainCode,
        'hotelId': value.hotelId,
        'roomType': value.roomType,
        'ratePlanCode': value.ratePlanCode,
        'languageCode': value.languageCode,
        'stickyFlagYn': value.stickyFlagYn,
        'usedInModule': UsedInModuleTypeToJSON(value.usedInModule),
        'beginDate': value.beginDate === undefined ? undefined : (value.beginDate.toISOString().substr(0,10)),
        'url': value.url,
        'type': value.type,
        'instance': value.instance,
        'idContext': value.idContext,
        'id': value.id,
        'idExtension': value.idExtension,
    };
}

